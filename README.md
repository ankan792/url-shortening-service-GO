This project is a simple URL shortening service written in Go and using Redis as the primary database. It uses **gofiber** as the web framework and **go-redis** as the Go redis client.

It provides an API for creating and retrieving shortened URLs, with a quota provided to prevent the user from using the service indefinitely.

The project consists of two Docker containers: one for the Redis database and one for the API service. The containers are configured using a docker-compose file, which makes it easy to run the project locally or deploy it to a server.

To run the project, you need to have Docker and docker-compose installed on your machine. 

Then, you can clone this repository using:
```git clone https://github.com/ankan792/url-shortening-service-GO```
and run the following commands in the project directory:

To build and start the docker containers, run ```docker-compose up -d```.
This will build and start the containers in the background. The API service will be available at http://localhost:3000/api/v1 by default. 
The API has two endpoints: `/api/v1` and `/:id`. 
The `/api/v1` endpoint takes a URL as a query parameter and returns a shortened URL along with expiry, use rate limit and duration in which the use rate limit will reset in JSON format. 
The `/:id` endpoint will redirect to the actual url, `:id` here is a dynamic and unique string.

To test the API, you can use Postman or run the following commands in the terminal:

```curl -d '{"url":"<THE URL YOU WANT TO BE SHORTENED>"}' -H "Content-Type: application/json" -X POST localhost:3000/api/v1```
This will make a Post request to the `/api/v1` endpoint, returning:
```
{
    "url": "THE URL YOU WANT TO BE SHORTENED",
    "short_url": "localhost:3000/SOME_RANDOM_ID",
    "expiry": 24,
    "rate_remaining": 10,
    "rate_reset": 30
}
```
To make your own custom ID, specify short_url in the json body request along with the actual url as:
```
{
    "url": "THE URL YOU WANT TO BE SHORTENED",
    "short_url": "YOUR_CUSTOM_ID"
}
```
You can make a GET request using curl ```localhost:3000/SOME_RANDOM_ID``` or you can also access the shortened URL directly in your browser and it will redirect you to the original URL.

To the stop the project run ```docker-compose down```. This will stop and remove the containers.

The project contains a `.env` file where you can change the domain name according to where you deploy the service and other variables like the database credentials, the port where the service will run and the quota limit.

