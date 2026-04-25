# http_servers

![Project Logo](./assets/logo.png)

## Project Title & Description

**http_servers** is a foundational project developed as part of the [Boot.dev](https://boot.dev/) course, designed to teach and demonstrate the core concepts of building HTTP servers using the Go programming language. This repository serves as a practical exploration of handling HTTP requests, routing, database interactions, and implementing robust authentication mechanisms from scratch.

## Key Features & Benefits

This project encapsulates various essential components of an HTTP server, offering the following key features and benefits:

*   **RESTful API Implementation:** Demonstrates the creation of RESTful endpoints for resource management, specifically for "chirps" (e.g., fetching, creating, deleting, and updating).
*   **Robust Authentication System:** Includes a custom-built authentication module featuring:
    *   **JWT (JSON Web Token) Integration:** Secure token-based authentication for stateless API interactions.
    *   **Password Hashing:** Utilizes `argon2id` for strong, secure password storage.
    *   **Refresh Tokens:** Implements a mechanism for renewing access tokens without requiring re-authentication.
*   **Database Integration:** Designed to interact with a PostgreSQL database, managing persistence for user data and "chirps."
*   **Modular Architecture:** Organized into `internal/` packages for better separation of concerns and maintainability (e.g., `auth` module, `database` module).
*   **Environment Variable Management:** Utilizes `.env` files for configurable settings, enhancing flexibility and security.
*   **Learning Resource:** Serves as an excellent practical example for learning Go's `net/http` package, middleware, and backend development best practices.

## Prerequisites & Dependencies

Before you begin, ensure you have the following installed:

*   **Go:** Version `1.26.1` or higher. You can download it from [golang.org/dl](https://golang.org/dl/).
*   **PostgreSQL:** A running instance of PostgreSQL is required as the project uses `github.com/lib/pq` for database interaction.
*   **Git:** For cloning the repository.

The project relies on several Go modules, which will be automatically downloaded during the setup phase. Key dependencies include:

| Package                                  | Description                           |
| :--------------------------------------- | :------------------------------------ |
| `github.com/alexedwards/argon2id`        | Argon2 password hashing library       |
| `github.com/golang-jwt/jwt/v5`           | JSON Web Token implementation         |
| `github.com/google/uuid`                 | UUID generation                       |
| `github.com/joho/godotenv`               | Loads environment variables from `.env` |
| `github.com/lib/pq`                      | PostgreSQL driver for Go              |
| `golang.org/x/crypto`                    | Go supplementary cryptography library |
| `golang.org/x/sys`                       | Go system call interface              |

## Installation & Setup Instructions

Follow these steps to get the project up and running on your local machine:

1.  **Clone the Repository:**
    ```bash
    git clone https://github.com/Adfmu1/http_servers.git
    cd http_servers
    ```

2.  **Install Go Modules:**
    Go will automatically download the necessary dependencies.
    ```bash
    go mod tidy
    ```

3.  **Database Setup:**
    *   Ensure you have a PostgreSQL server running.
    *   Create a new database for the project (e.g., `chirpy_db`).
    *   The project expects a `DATABASE_URL` environment variable. You might need to create initial tables, which would typically be handled by migration scripts (not explicitly provided in the structure but implied by the `database` package).

4.  **Configuration (.env file):**
    Create a `.env` file in the root directory of the project. This file will hold your environment-specific configurations.
    ```
    # Example .env file
    DB_URL="postgres://user:password@localhost:5432/chirpy_db?sslmode=disable"
    PLATFORM="platform"
    SECRET="super-secret-key"
    POLKA_KEY="your-polka-api-key" # If POLKA_API is used
    ```
    *   Replace `user`, `password`, `localhost:5432`, `chirpy_db`, and secret keys with your actual database credentials and secure random strings. It's just a way to have your DB URL saved somewhere in case you forget.

5.  **Run the Server:**
    ```bash
    go run .
    ```
    The server should now be running on port 8080 (e.g., `http://localhost:8080`).

## Usage Examples & API Documentation

Once the server is running, you can interact with its API endpoints. The project's structure suggests the following core functionalities:

**Base URL:** `http://localhost:<PORT>` (e.g., `http://localhost:8080`)

### Authentication Endpoints

| Method | Endpoint             | Description                                  |
| :----- | :------------------- | :------------------------------------------- |
| `POST` | `/api/login`             | Authenticate a user and receive JWT tokens.  |
| `POST` | `/api/refresh`           | Refresh an expired access token using a refresh token. |
| `POST` | `/api/revoke`            | Invalidate a refresh token, logging out a user. |

### User Endpoints

| Method | Endpoint             | Description                                  |
| :----- | :------------------- | :-------------------------------------------- |
| `PUT`  | `/api/users`         | Update user data (e.g., email, password). Requires authentication. |

### Chirp Endpoints

"Chirps" are likely the primary resource of this application, similar to tweets or posts.

| Method   | Endpoint               | Description                                | Requires Auth |
| :------- | :--------------------- | :----------------------------------------- | :------------ |
| `GET`    | `/api/chirps`              | Retrieve all chirps.                       | No            |
| `POST`   | `/api/chirps`              | Create a new chirp.                        | Yes           |
| `GET`    | `/api/chirps/{id}`         | Retrieve a single chirp by its ID.         | No            |
| `DELETE` | `/api/chirps/{id}`         | Delete a chirp by its ID.                  | Yes           |

### Example: Getting Started

1.  **Start the server:**
    ```bash
    go run .
    ```

2.  **Access the welcome page:**
    Open your browser to `http://localhost:8080/` (or your configured port). You should see:
    ```html
    <html>
        <body>
            <h1>Welcome to Chirpy</h1>
        <body>
    </html>
    ```

3.  **Register a User (Hypothetical):**
    You'd `POST` to `/api/users` to create an account.
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"email": "test@example.com", "password": "securepassword"}' http://localhost:8080/users
    ```

4.  **Log In:**
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"email": "test@example.com", "password": "securepassword"}' http://localhost:8080/login
    ```
    This will return your JWT access and refresh tokens.

5.  **Create a Chirp (Requires Authorization Header):**
    ```bash
    curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer <YOUR_JWT_ACCESS_TOKEN>" -d '{"body": "Hello, Chirpy world!"}' http://localhost:8080/chirps
    ```

## Configuration Options

The primary way to configure `http_servers` is through environment variables, typically loaded from a `.env` file, you can also modify some of hard programmed values in main.go, ex. port etc.

*   **`DB_URL`**: The connection string for the PostgreSQL database (e.g., `postgres://user:password@host:port/dbname?sslmode=disable`).
*   **`SECRET`**: A secret key used for signing and verifying JWT access tokens and refresh tokens.
*   **`POLKA_KEY`**: If integrated with a "Polka API" (imaginary Boot.dev API), this would be the API key for authentication with that external service.

## 🗄️ Tech Stack

| Technology    | Purpose          |
|---------------|------------------|
| Go (net/http) | HTTP server      |
| PostgreSQL    | Database         |
| SQLC          | Type-safe SQL    |
| JWT (v5)      | Authentication   |
| Argon2id      | Password hashing |
| Goose         | Migrations       |

## ⚙️ Architecture
Client
↓
HTTP Handlers (net/http)
↓
Auth Layer (JWT + Refresh Tokens)
↓
Database Layer (SQLC + PostgreSQL)
