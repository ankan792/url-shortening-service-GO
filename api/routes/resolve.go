package routes

import (
	"github.com/ankan792/url-shortening-service-GO/api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("id")

	db := database.CreateDbClient(0)
	defer db.Close()

	val, err := db.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "no shortened URL found"})
	} else if err != nil {
		panic(err)
	}

	return c.Redirect(val, 301)

}
