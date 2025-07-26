package handlers

import (
	"fmt"
	"time"
	"trg/internal/services"
	"trg/pkg/cache.go"

	"github.com/gofiber/fiber/v2"
)

func Trace(cc cache.Provider) fiber.Handler {
	return func(c *fiber.Ctx) error {
		addr := c.Query("addr", "")

		result := services.Trace(addr)
		oc := cache.NewObjectCacher[[]*services.Result](cc, cache.SerializationTypeJSON)

		oc.Set(c.UserContext(), fmt.Sprintf("%s:%s", c.Request().String(), c.Context().Time()), time.Hour*1, result)

		return c.Status(200).JSON(result)
	}
}
