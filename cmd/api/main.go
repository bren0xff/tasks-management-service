package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	_ "tasksManagement/docs"
	"tasksManagement/internal/delivery/http"
	"tasksManagement/internal/entity"
	"tasksManagement/internal/notifier"
	repository "tasksManagement/internal/repository/impl"
	"tasksManagement/internal/usecase"
	"tasksManagement/pkg/queue"
)

// @title Tasks Management API
// @version 1.0
// @description This is a REST API for managing tasks performed by technicians and managers.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Use the format: Bearer <token>
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using existing environment variables")
	}

	serverPort := os.Getenv("SERVER_PORT")
	mysqlDSN := os.Getenv("MYSQL_DSN")
	jwtSecret := os.Getenv("JWT_SECRET")
	rabbitmqURL := os.Getenv("RABBITMQ_URL")

	db, err := gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = db.AutoMigrate(&entity.Task{}, &entity.User{})
	if err != nil {
		log.Fatal("Error performing database migrations:", err)
	}

	q, err := queue.NewRabbitMQ(rabbitmqURL)
	if err != nil {
		log.Fatal("Error connecting to RabbitMQ:", err)
	}

	taskRepo := repository.NewTaskRepository(db)
	userRepo := repository.NewUserRepository(db)

	newNotifier := notifier.NewNotifier(q)

	taskUseCase := usecase.NewTaskUseCase(taskRepo, userRepo, newNotifier)
	userUseCase := usecase.NewUserUseCase(userRepo, jwtSecret)

	e := echo.New()

	http.NewTaskHandler(e, taskUseCase, jwtSecret)
	http.NewUserHandler(e, userUseCase)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	log.Printf("Server started on port %s", serverPort)
	if err := e.Start(":" + serverPort); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
