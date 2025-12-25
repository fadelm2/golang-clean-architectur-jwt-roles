package middleware

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		requestID := uuid.New().String()

		c.Locals("request_id", requestID)

		err := c.Next()

		status := c.Response().StatusCode()

		log := log.WithFields(logrus.Fields{
			"request_id": requestID,
			"method":     c.Method(),
			"path":       c.OriginalURL(),
			"status":     status,
			"latency_ms": time.Since(start).Milliseconds(),
			"ip":         c.IP(),
		})

		// ❌ jangan log body untuk endpoint auth
		if strings.Contains(c.Path(), "/login") ||
			strings.Contains(c.Path(), "/register") {
			log.Info("request finished (auth endpoint)")
			return err
		}

		log.Info("request finished")
		return err
	}
}
