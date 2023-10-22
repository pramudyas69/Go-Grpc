package main

import (
	"github.com/pramudyas69/Go-Grpc/server/internal/component"
	"github.com/pramudyas69/Go-Grpc/server/internal/config"
	rpc "github.com/pramudyas69/Go-Grpc/server/internal/grpc"
	"github.com/pramudyas69/Go-Grpc/server/internal/middleware"
	"github.com/pramudyas69/Go-Grpc/server/internal/repository"
	"github.com/pramudyas69/Go-Grpc/server/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cnf := config.Get()

	mongoClient, err := component.NewMongoClient(cnf)
	if err != nil {
		log.Fatalf("Error connecting to mongo: %v", err)
	}
	defer mongoClient.Disconnect()

	userRepo := repository.NewUser(mongoClient)

	userSvc := service.NewUser(userRepo, cnf)

	lis, err := net.Listen("tcp", ":"+cnf.Server.Port)
	if err != nil {
		log.Fatalf("failed to listen : %s", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.Authenticate(cnf, "Authorization")),
	)

	rpc.NewUser(grpcServer, userSvc)
	log.Println("Server running on port " + cnf.Server.Port)
	log.Fatal(grpcServer.Serve(lis))
}
