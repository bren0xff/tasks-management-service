package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	echoSwagger "github.com/swaggo/echo-swagger"
	"tasksManagement/internal/delivery/http"
	"tasksManagement/internal/entity"
	"tasksManagement/internal/notifier"
	"tasksManagement/internal/repository"
	"tasksManagement/internal/usecase"
	"tasksManagement/pkg/queue"
)

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
	notifier := notifier.NewNotifier(q)
	taskUseCase := usecase.NewTaskUseCase(taskRepo, notifier)

	e := echo.New()

	http.NewTaskHandler(e, taskUseCase, jwtSecret)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	log.Printf("Server started on port %s", serverPort)
	if err := e.Start(":" + serverPort); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
