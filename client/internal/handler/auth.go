package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/pramudyas69/Go-Grpc/client/dto"
	pb "github.com/pramudyas69/Go-Grpc/client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type authHandler struct {
}

func NewAuth(app *fiber.App, authMid fiber.Handler) {
	h := &authHandler{}

	app.Post("/register", h.Register)
	app.Post("/login", h.Login)
	app.Get("/validate", authMid, h.ValidateToken)
}

func (a authHandler) Register(ctx *fiber.Ctx) error {
	var req dto.RegisterReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"Error": "Failed to connect to the gRPC Server %s, err",
		})
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	userReq := &pb.RegisterRequest{
		Fullname: req.Fullname,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	_, err = client.Register(context.Background(), userReq)
	if err != nil {
		return ctx.SendStatus(500)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"Message": "Register Successfully",
	})
}

func (a authHandler) Login(ctx *fiber.Ctx) error {
	var req dto.LoginReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"Error": "Failed to connect to the gRPC Server %s, err",
		})
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	userReq := &pb.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	res, err := client.Login(context.Background(), userReq)
	if err != nil {
		return ctx.SendStatus(500)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"Data": res.Token,
	})
}

func (a authHandler) ValidateToken(ctx *fiber.Ctx) error {
	token := ctx.Context().UserValue("token")

	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"Error": "Failed to connect to the gRPC Server %s, err",
		})
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	userReq := &pb.ValidateTokenRequest{
		Token: token.(string),
	}

	md := metadata.Pairs("Authorization", token.(string))

	c := metadata.NewOutgoingContext(context.Background(), md)

	_, err = client.ValidateToken(c, userReq)
	if err != nil {
		return ctx.SendStatus(500)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"Data": "Valid!",
	})
}
