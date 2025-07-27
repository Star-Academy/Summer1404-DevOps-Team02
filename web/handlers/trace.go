package handlers

import (
	"fmt"
	"time"
	"trg/internal/services"
	"trg/pkg/cache.go"

	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"encoding/json"
)

func add_log(data []*services.Result){
	file, err := os.OpenFile("/var/log/my_traceroute.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if (err != nil){
		log.Fatal(err)
	}
	defer file.Close()

	logEntry := struct {
		Timestamp time.Time          `json:"timestamp"`
		Results   []*services.Result `json:"results"`
	}{
		Timestamp: time.Now(),
		Results:   data,
	}

	jsonData, err := json.Marshal(logEntry)
	if err != nil {
		log.Printf("Failed to marshal results to JSON: %v", err)
		return
	}

	// Write to log file with newline
	if _, err := file.Write(append(jsonData, '\n')); err != nil {
		log.Printf("Failed to write to log file: %v", err)
	}
}

func Trace(cc cache.Provider) fiber.Handler {
	return func(c *fiber.Ctx) error {
		addr := c.Query("addr", "")

		result := services.Trace(addr)
		oc := cache.NewObjectCacher[[]*services.Result](cc, cache.SerializationTypeJSON)

		oc.Set(c.UserContext(), fmt.Sprintf("%s:%s", c.Request().String(), c.Context().Time()), time.Hour*1, result)
		
		add_log(result)

		return c.Status(200).JSON(result)
	}
}
