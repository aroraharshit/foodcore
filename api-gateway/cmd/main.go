package main

import (
	"log"

	"github.com/aroraharshit/foodcore/api-gateway/controllers"
	proto "github.com/aroraharshit/foodcore/api-gateway/proto"
	"github.com/aroraharshit/foodcore/api-gateway/routes"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	defer conn.Close()

	if err != nil {
		log.Fatalf("Could not connect to user service %v", err)
	}

	defer conn.Close()

	userClient := proto.NewUserServiceClient(conn)
	userController := controllers.NewUserController(userClient)

	r := gin.Default()
	routes.UserRoutes(r, userController)

	log.Println("Api Gateway running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server existed with: %v", err)
	}

}
