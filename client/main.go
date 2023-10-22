package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pramudyas69/Go-Grpc/client/internal/handler"
	"github.com/pramudyas69/Go-Grpc/client/internal/middleware"
)

func main() {
	authMid := middleware.AuthHandler()
	app := fiber.New()
	handler.NewAuth(app, authMid)

	app.Listen(":8080")
}
