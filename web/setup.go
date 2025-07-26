package web

import (
	"trg/pkg/redis"
	"trg/web/handlers"

	"github.com/gofiber/fiber/v2"
)

func Run() error {

	api := fiber.New()
	c := redis.NewRedisProvider()
	api.Get("/", handlers.Trace(c))

	return api.Listen(":8000")
}
