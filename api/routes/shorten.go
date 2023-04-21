package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/ankan792/url-shortening-service-GO/api/database"
	"github.com/ankan792/url-shortening-service-GO/api/helpers"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Request struct {
	URL      string        `json:"url"`
	ShortURL string        `json:"short_url"`
	Expiry   time.Duration `json:"expiry"`
}

type Response struct {
	URL             string        `json:"url"`
	ShortURL        string        `json:"short_url"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_remaining"`
	XRateLimitReset time.Duration `json:"rate_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(Request)

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("failed to parse JSON")
	}

	//implement rate limiting
	db := database.CreateDbClient(0)
	defer db.Close()

	_, err = db.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil {
		db.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second)
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to connect to DB!"})
	}
	valIP, _ := db.Get(database.Ctx, c.IP()).Result()
	valIPInt, _ := strconv.Atoi(valIP)
	limit, _ := db.TTL(database.Ctx, c.IP()).Result()
	if valIPInt <= 0 {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error":       "Limit for API calls has exceeded!",
			"limit_reset": limit / time.Minute,
		})
	}

	//check whether the url is an actual url
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	//check for domain errors
	if !helpers.ContainsDomainError(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "That is already a " + os.Getenv("DOMAIN") + " link"})
	}

	//enforce http
	body.URL = helpers.EnforceHTTP(body.URL)

	//shorten URL

	var id string

	if body.ShortURL == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.ShortURL
	}

	val, _ := db.Get(database.Ctx, id).Result()
	if val != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Custom URL already in use"})
	}

	if body.Expiry == 0 {
		body.Expiry = 24
	}

	err = db.Set(database.Ctx, id, body.URL, body.Expiry*time.Hour).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to connect to DB!"})
	}

	resp := Response{
		URL:             body.URL,
		ShortURL:        "",
		Expiry:          body.Expiry,
		XRateRemaining:  10,
		XRateLimitReset: 30,
	}

	//decrements the API call value for the specific user by 1
	db.Decr(database.Ctx, c.IP())

	resp.XRateRemaining = valIPInt
	resp.XRateLimitReset = limit / time.Minute
	resp.ShortURL = "http://" + os.Getenv("DOMAIN") + "/" + id

	return c.Status(fiber.StatusOK).JSON(resp)
}
