
# Golang Blog API

This is an advanced RESTful API built with Go (Golang) for a simple blog platform. It includes features like user authentication, CRUD operations for blog posts, and database support using PostgreSQL. The project is also containerized for easy deployment.

---

## Features

- **User Authentication**: JWT-based authentication.
- **CRUD Operations**: Create, read, update, and delete blog posts.
- **PostgreSQL Integration**: Data persistence using PostgreSQL.
- **Docker Support**: Containerized with a `Dockerfile` for deployment.

---

## Project Structure

```
golang-blog-api/
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── config/
│   └── config.go
├── controllers/
│   ├── auth_controller.go
│   ├── post_controller.go
├── middleware/
│   └── auth_middleware.go
├── models/
│   ├── post.go
│   ├── user.go
├── routes/
│   └── routes.go
├── utils/
│   └── token_utils.go
└── db/
    └── migration.sql
```

---

## Prerequisites

1. **Go** (v1.20 or higher) - [Download](https://golang.org/dl/)
2. **PostgreSQL** - [Install PostgreSQL](https://www.postgresql.org/download/)

---

## Setup Instructions

### 1. Clone the Project

Extract the downloaded zip file or clone it (if uploaded to a repository).

```bash
cd golang-blog-api
```

### 2. Install Dependencies

Run the following command to install required Go modules:

```bash
go mod tidy
```

### 3. Set Up the Database

1. Create a `.env` file in the root directory with the following variables:

   ```
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_NAME=your_database
   DB_HOST=localhost
   DB_PORT=5432
   ```

2. Run the migration file to set up the `posts` table:

   ```bash
   psql -U your_user -d your_database -f db/migration.sql
   ```

### 4. Run the Application

Start the application:

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

---

## Endpoints

| Method | Endpoint       | Description             |
|--------|----------------|-------------------------|
| POST   | `/login`       | User login              |
| GET    | `/posts`       | Get all blog posts      |

---

## Docker Deployment

Build and run the Docker container:

1. Build the Docker image:
   ```bash
   docker build -t golang-blog-api .
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 golang-blog-api
   ```

---

## License

This project is licensed under the MIT License.
