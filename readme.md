# Tasks Management API

This is a containerized REST API for managing tasks, developed with **Golang**, **MySQL**, and **RabbitMQ**. It allows technicians to create tasks, while managers can view, delete, and be notified of tasks.

## Features

- User roles:
    - Technicians can create and view their tasks.
    - Managers can view all tasks, delete them, and receive notifications when a task is created.
- JWT-based authentication for secure access.
- MySQL as the primary database for storing tasks and users.
- RabbitMQ for sending asynchronous task notifications.
- Fully containerized with Docker for easy deployment.
- Swagger documentation for API exploration.

## Prerequisites

Ensure you have the following tools installed:

- Docker
- Docker Compose

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/brenodl/tasks-management-service.git
   cd tasks-management-service
   
2. Create a .env file in the project root and configure it as follows:
    ```bash
    #Server Configuration
    SERVER_PORT=8080
    
    # MySQL Configuration
    MYSQL_DSN=managementUser:managementPassword@tcp(mysql:3306)/tasks?charset=utf8mb4&parseTime=True&loc=Local
    
    # RabbitMQ Configuration
    RABBITMQ_URL=amqp://managementUser:managementPassword@rabbitmq:5672/
    
    # JWT Configuration
    JWT_SECRET=your-secret-key
    
    #Replace managementUser, managementPassword, and your-secret-key with your secure values.

3. Build and start the services:
    ```bash
   docker-compose up --build

This will:

Build the API service using the Dockerfile.
Start MySQL and RabbitMQ services.
Launch the API on http://localhost:8080.

### Usage
Access the API documentation at http://localhost:8080/swagger/index.html.

