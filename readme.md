
# Tasks Management API

This project is a RESTful API designed to manage tasks performed by technicians and managers. It provides functionalities for task creation, retrieval, user registration, and authentication. The API is built using **Go**, **Echo**, and **GORM**, with support for **JWT-based authentication** and **RabbitMQ** for notifications.

---

## Features

- **User Management**:
    - Register new users.
    - Authenticate users and generate JWT tokens.
- **Task Management**:
    - Create tasks.
    - Retrieve tasks based on user roles (`Manager` or `Technician`).
- **Role-Based Access Control (RBAC)**:
    - Managers can access all tasks.
    - Technicians can access only their own tasks.
- **Notifications**:
    - Notify managers when tasks are created by technicians.

---

## Technologies Used

- **Programming Language**: Go
- **Web Framework**: Echo
- **Database**: MySQL
- **ORM**: GORM
- **Authentication**: JWT
- **Queue**: RabbitMQ
- **Documentation**: Swagger

---

## Getting Started

### Prerequisites

Ensure you have the following tools installed:

- [Go](https://golang.org/dl/) 1.19+
- [Docker](https://www.docker.com/) (Optional for running services)
- [Docker Compose](https://docs.docker.com/compose/)
- [MySQL](https://www.mysql.com/) 8.0+
- [RabbitMQ](https://www.rabbitmq.com/) 3.11+

---

### Environment Variables

Create a `.env` file in the root directory with the following variables:

```bash
MYSQL_DSN="user:password@tcp(localhost:3306)/tasks_management?charset=utf8mb4&parseTime=True&loc=Local"
RABBITMQ_URL="amqp://guest:guest@localhost:5672/"
SERVER_PORT=8080
MYSQL_USER=managementUser
MYSQL_PASSWORD=managementPassword
RABBITMQ_DEFAULT_USER=managementUser
RABBITMQ_DEFAULT_PASS=managementPassword
JWT_SECRET=1243c0ad9f1f2bf5be4f6b310fe2b53394ad2339f5d80407e3d00cab61b2290d


```
---

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/brenodl/tasks-management-service.git
   cd tasks-management-service
   ```

2. Build and start the services:
   ```bash
   docker-compose up --build
   ```

   This will:
    - Build the API service using the `Dockerfile`.
    - Start MySQL and RabbitMQ services.
    - Launch the API on `http://localhost:8080`.

---

### Usage

Access the API documentation at `http://localhost:8080/swagger/index.html`.

#### Example Requests

- **Register User**:

`POST /users/register`
```json
  {
    "name": "Fulano de tal",
    "email": "fulano.tal@example.com",
    "password": "securepassword",
    "role": "technician"
  }
```

- **Login**:

`POST /users/login`
  ```json
  {
    "email": "fulano.tal@example.com",
    "password": "securepassword"
  }
  ```

- **Create Task**:

`POST /tasks`
  ```json
  {
    "summary": "Fix server issue",
    "date": "2024-11-12"
  }
  ```

---

## API Documentation

Swagger documentation is available at:

```
http://localhost:8080/swagger/index.html
```

---

## Testing

Run the tests using the following command:

```bash
go test -v ./...
```

---

## Project Structure

```
tasks-management-service/
│
├── internal/
│   ├── delivery/
│   │   └── http/         # HTTP Handlers
│   ├── entity/           # Data Models
│   ├── repository/       # Database Access Layer
│   │   └── impl/         # Repository Implementations
│   ├── usecase/          # Business Logic Layer
│   ├── notifier/         # Notification Service
│   └── pkg/
│       └── queue/        # RabbitMQ Integration
│
├── docs/                 # Swagger Documentation
├── docker-compose.yml    # Docker Compose Configuration
├── Dockerfile            # Dockerfile for API Service
├── main.go               # Entry Point
├── go.mod                # Dependencies
└── README.md             # Project Documentation
```

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Contributors

- **Breno Silveira** - [GitHub Profile](https://github.com/brenodl)

Feel free to contribute by opening issues or submitting pull requests!
