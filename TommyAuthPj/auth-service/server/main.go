package main

import (
	"log"

	"auth-service/config"
	"auth-service/internal/handler"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"auth-service/router"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("database initialization failed: %v", err)
	}
	// defer db.Close()

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	r := router.SetupRouter(authHandler)

	r.POST("/register", authHandler.Register)

	r.Run(":8080")
}
