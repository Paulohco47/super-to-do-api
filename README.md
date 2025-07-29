# To-Do List API in Go

A complete RESTful API for task management, developed in Go with PostgreSQL as the database and Docker configuration for easy deployment.

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Installation & Running](#installation--running)
- [API Documentation](#api-documentation)
- [Usage Examples](#usage-examples)
- [Database Structure](#database-structure)
- [Configuration](#configuration)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## 🎯 Overview

This To-Do List API was developed as an educational project to demonstrate the fundamental concepts of API development in Go, including:

- Use of Go's native `http.Server`
- JSON handling with `encoding/json`
- Basic route handling
- Integration with PostgreSQL database
- Containerization with Docker
- Full CRUD operations

## ✨ Features

- **Full CRUD**: Create, read, update, and delete tasks
- **RESTful API**: Well-defined endpoints following REST standards
- **Database**: Persistence with PostgreSQL
- **Containerization**: Docker and Docker Compose setup
- **Validation**: Input data validation
- **Error Handling**: Standardized error responses
- **Timestamps**: Automatic creation and update control
- **Health Check**: Database health check

## 🛠 Technologies Used

- **Go 1.18+**: Main programming language
- **PostgreSQL 13**: Relational database
- **Docker & Docker Compose**: Containerization and orchestration
- **github.com/lib/pq**: PostgreSQL driver for Go

## 📁 Estrutura do Projeto

```
super-to-do-api/
├── main.go              # Application entry point
├── models.go            # Data structures and models
├── handlers.go          # HTTP handlers and business logic
├── database.go          # Database configuration and connection
├── init.sql             # Database initialization script
├── Dockerfile           # API container configuration
├── docker-compose.yaml  # Services orchestration
├── .dockerignore        # Files ignored in Docker build
├── go.mod               # Go module dependencies
├── go.sum               # Dependencies checksums
├── docs/                # Swagger/OpenAPI documentation
│   ├── docs.go          # Documentation generation code
│   ├── swagger.json     # Swagger specification in JSON
│   └── swagger.yaml     # Swagger specification in YAML
└── README.md            # Project documentation
```

## 📋 Pré-requisitos

 - Docker 20.10+
 - Docker Compose 1.29+

**OR** for local execution:

 - Go 1.18+
 - PostgreSQL 17+

## 🚀 Instalação e Execução

### Using Docker Compose (Recommended)

1. Clone the repository:
```bash
git clone https://github.com/Paulohco47/super-to-do-api
cd super-to-do-api
```

2. Run with Docker Compose:
```bash
docker-compose up --build
```

3. The API will be available at `http://localhost:8080`

### Local Execution

1. Install dependencies:
```bash
go mod download
```

2. Set environment variables:
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=password
export DB_NAME=todoapp
```

3. Run the application:
```bash
go run .
```

## 📚 Documentação da API

### Base URL
```
http://localhost:8080
```

### Endpoints

#### 1. List all tasks
- **GET** `/tasks`
- **Description**: Returns all registered tasks
- **Success Response**: 200 OK

#### 2. Get a specific task
- **GET** `/tasks/{id}`
- **Description**: Returns a specific task by ID
- **Parameters**: `id` (integer) - Task ID
- **Success Response**: 200 OK
- **Error Response**: 404 Not Found

#### 3. Create a new task
- **POST** `/tasks`
- **Description**: Creates a new task
- **Body**: JSON with `title` and `description`
- **Success Response**: 201 Created

#### 4. Update a task
- **PUT** `/tasks/{id}`
- **Description**: Updates an existing task
- **Parameters**: `id` (integer) - Task ID
- **Body**: JSON with optional fields (`title`, `description`, `completed`)
- **Success Response**: 200 OK
- **Error Response**: 404 Not Found

#### 5. Delete a task
- **DELETE** `/tasks/{id}`
- **Description**: Removes a task
- **Parameters**: `id` (integer) - Task ID
- **Success Response**: 200 OK
- **Error Response**: 404 Not Found

### Data Structure

#### Task
```json
{
  "id": 1,
  "title": "Task title",
  "description": "Detailed task description",
  "completed": false,
  "created_at": "2023-07-25T10:30:00Z",
  "updated_at": "2023-07-25T10:30:00Z"
}
```

#### Default API Response
```json
{
  "success": true,
  "message": "Optional message",
  "data": {}
}
```

## 🔧 Exemplos de Uso

### 1. Create a new task
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Study Go",
    "description": "Learn the basics of the Go language"
  }'
```

### 2. List all tasks
```bash
curl -X GET http://localhost:8080/tasks
```

### 3. Get a specific task
```bash
curl -X GET http://localhost:8080/tasks/1
```

### 4. Update a task
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{
    "completed": true
  }'
```

### 5. Delete a task
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

## 🗄 Estrutura do Banco de Dados

### Table: tasks

| Field       | Type                        | Description                |
|-------------|-----------------------------|----------------------------|
| id          | SERIAL PRIMARY KEY          | Unique task identifier     |
| title       | VARCHAR(255) NOT NULL       | Task title                 |
| description | TEXT                        | Detailed task description  |
| completed   | BOOLEAN DEFAULT FALSE       | Completion status          |
| created_at  | TIMESTAMP DEFAULT CURRENT_TIMESTAMP | Creation date      |
| updated_at  | TIMESTAMP DEFAULT CURRENT_TIMESTAMP | Last update date   |

### Indexes
- `idx_tasks_completed`: Index on the `completed` field for status queries
- `idx_tasks_created_at`: Index on the `created_at` field for temporal ordering

## ⚙️ Configuração

### Environment Variables

| Variable    | Description                  | Default Value |
|-------------|------------------------------|---------------|
| DB_HOST     | PostgreSQL host              | localhost     |
| DB_PORT     | PostgreSQL port              | 5432          |
| DB_USER     | Database user                | postgres      |
| DB_PASSWORD | Database password            | password      |
| DB_NAME     | Database name                | todoapp       |

### Docker Compose

The `docker-compose.yaml` file configures:
- **postgres**: PostgreSQL container with persistent data
- **todo-api**: Go application container
- **Network**: Bridge network for container communication
- **Volumes**: PostgreSQL data persistence
- **Health Check**: Database health check before starting the API

## 🧪 Testes

To test the API, you can use:

1. **curl** (examples above)
2. **Postman** or **Insomnia**
3. **HTTPie**:
```bash
# Create task
http POST localhost:8080/tasks title="New task" description="Description"

# List tasks
http GET localhost:8080/tasks

# Update task
http PUT localhost:8080/tasks/1 completed:=true
```

## 🤝 Contribuição

1. Fork the project
2. Create a branch for your feature (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 Licença

This project is under the MIT license. See the `LICENSE` file for more details.

---

**Powered by caffeine and Go code ☕️⚙️**

