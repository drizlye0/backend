package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/biznetbb/user-service/src/adapters/repositories"
	"github.com/biznetbb/user-service/src/adapters/router"
	"github.com/biznetbb/user-service/src/application/services"
	"github.com/biznetbb/user-service/src/infraestructure/db"
	"github.com/biznetbb/user-service/src/infraestructure/registry"
	"github.com/gin-gonic/gin"
)

func main() {
	EUREKAURL := "http://registry-service:8761/eureka" //os.Getenv("EUREKA_CLIENT_SERVICEURL_DEFAULTZONE")
	appName := "user-microservice"
	hostname := "user-microservice"
	ipAddr := getoutBoundIp().String()
	port := 8086

	db, err := db.DBConnection()
	if err != nil {
		log.Fatalf("error initializing database: %v", err)
	}

	fmt.Println("Database connection successful", db)

	err = os.MkdirAll("avatars", os.ModePerm)
	if err != nil {
		fmt.Printf("error creating avatars directory: %v\n", err)
	}

	client := registry.NewEurekaClient(EUREKAURL, appName, hostname, ipAddr, port)

	if err := client.Register(); err != nil {
		log.Fatalf("Failed to register service: %v", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	r := router.NewRouter(userService)
	router := r.SetupRoutes()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Run(":8086")
}

func getoutBoundIp() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}
