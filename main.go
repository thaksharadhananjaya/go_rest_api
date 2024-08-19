package main

import (
	"log"
	"restapi/config"
	"restapi/internal/infrastructure/database"
	"restapi/internal/interface/controller"
	"restapi/internal/interface/middleware"
	"restapi/internal/interface/repository"
	"restapi/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found or failed to load .env file")
    }

    cfg := config.LoadConfig()

    db := database.NewPostgresDB(cfg.DatabaseDSN)
    userRepository := repository.NewGormUserRepository(db)
    userUseCase := usecase.NewUserUseCase(userRepository)
    userController := controller.NewUserController(userUseCase)

	// Set up Gin router
    router := gin.Default()

	// Initialize error-handling middleware
    errorMiddleware := middleware.ErrorHandler()
	router.Use(errorMiddleware)
	protected := router.Group("/")
	protected.Use(middleware.JWTMiddleware())


    protected.POST("/users", userController.CreateUser)
    router.GET("/users/:id", userController.GetUserByID)
    router.GET("/users", userController.GetAllUsers)
    protected.PATCH("/users/:id", userController.UpdateUser)
    protected.DELETE("/users/:id", userController.DeleteUser)

    router.Run(cfg.ServerAddress)
}
