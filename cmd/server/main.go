package main

import (
	"GO-ECOMMERCE-BACKEND-API/configs"
	"GO-ECOMMERCE-BACKEND-API/internal/controller"
	"GO-ECOMMERCE-BACKEND-API/internal/model"
	"GO-ECOMMERCE-BACKEND-API/internal/repo"
	"GO-ECOMMERCE-BACKEND-API/internal/routers"
	"GO-ECOMMERCE-BACKEND-API/internal/service"
	"fmt"
)

func main() {
	fmt.Println("starting")
	// Initialize database connection
	configs.InitDB()

	// Migrate the database
	configs.DB.AutoMigrate(&model.User{})

	// Set up repository, service, and controller
	userRepo := repo.NewUserRepository(configs.DB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Set up router
	r := routers.SetupRouter(userController)

	// Start server
	r.Run(":3000")

}
