package middleware

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func AuthHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := strings.ReplaceAll(ctx.Get("Authorization"), "Bearer ", "")
		if token == "" {
			return ctx.SendStatus(401)
		}

		ctx.Context().SetUserValue("token", token)
		return ctx.Next()
	}
}
