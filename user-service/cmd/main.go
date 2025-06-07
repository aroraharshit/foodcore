package main

import (
	"fmt"
	"log"
	"net"

	"github.com/aroraharshit/go-foodcore/user-service/config"
	"github.com/aroraharshit/go-foodcore/user-service/controllers"
	proto "github.com/aroraharshit/go-foodcore/user-service/proto"
	"github.com/aroraharshit/go-foodcore/user-service/repositories"
	"github.com/aroraharshit/go-foodcore/user-service/services"
	"google.golang.org/grpc"
)

func main() {
	db := config.ConnectDB("mongodb://localhost:27017", "foodcore")
	userCollection := db.Collection("users")
	repo := repositories.NewUserRepository(userCollection)
	service := services.NewUserService(repo)
	controller := controllers.NewUserController(service)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, controller)

	fmt.Println("User Service gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRpc server failed: %v", err)
	}

}
