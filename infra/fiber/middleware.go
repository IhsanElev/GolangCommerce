package infrafiber

import (
	"GolangEcommerceDDD/infra/response"
	"GolangEcommerceDDD/internal/config"
	"GolangEcommerceDDD/utility"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Trace() fiber.Handler {
	return func(c *fiber.Ctx) error {

		start := time.Now()

		// Log the request information
		log := logrus.WithFields(logrus.Fields{
			"method": c.Method(),
			"path":   c.Path(),
			"ip":     c.IP(),
		})

		// Log request headers
		log.Info("Request Headers:", c.Request().Header)

		// Proceed with the request
		err := c.Next()

		// Log the response information and timing
		log = log.WithFields(logrus.Fields{
			"status":     c.Response().StatusCode,
			"latency_ms": time.Since(start).Milliseconds(),
		})

		// Log response headers
		log.Info("Response Headers:", c.Response().Header)

		return err
	}
}
func CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		if authorization == "" {
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		bearer := strings.Split(authorization, "Bearer ")
		if len(bearer) != 2 {
			log.Println("token invalid")
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		token := bearer[1]

		publicId, role, err := utility.ValidateToken(token, config.Cfg.App.Encryption.JWTSecret)
		if err != nil {
			log.Println(err.Error())
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		c.Locals("ROLE", role)
		c.Locals("PUBLIC_ID", publicId)

		return c.Next()
	}
}

func CheckRoles(authorizedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := fmt.Sprintf("%v", c.Locals("ROLE"))
		isExist := false
		for _, authorizedRoles := range authorizedRoles {
			if role == authorizedRoles {
				isExist = true
				break
			}
		}
		if !isExist {
			return NewResponse(
				WithError(response.ErrorForbiddenAccess),
			).Send(c)
		}
		return c.Next()
	}
}
