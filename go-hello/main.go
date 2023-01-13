package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-hello/controller"
	"go-hello/middleware"
	"go-hello/model"
	"go-hello/storage"
	"log"
)

func main() {
	loadEnv()
	loadDatabase()
	router := gin.Default()
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)
	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.GET("/users", controller.ListUser)

	if err := router.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	storage.Connect()
	if err := storage.Database.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("Database auto migrate failed")
	}
}
