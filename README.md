# GoAuthFlow

**GoAuthFlow** is a Go-based authentication API service designed with clean architecture principles. It provides secure user registration and login endpoints with hashed password storage, using a modular codebase to promote maintainability and scalability.

## Features

- **User Registration** (`POST /v1/register`)
- **User Login** (`POST /v1/login`)
- **Password Hashing** for security
- **SQLite3 Backend** for persistence
- **Gin Web Framework** for fast HTTP handling
- **Flexible Architecture:** Data, Service, and Transport layers
- **Middleware Support:** Logging and error handling
- **Dockerized Deployment**

## Tech Stack

- **Go** 1.24.5+
- **Gin** (v1.11.0)
- **SQLite3** (v1.14.32)
- **Docker**

## Project Structure

API/
Data/ # Data models and repository
Service/ # Business logic and authentication
Transport/ # HTTP handlers, routes, middleware
DOC/ # License and Readme
main.go # Application entry point
Dockerfile # Container configuration
go.mod # Dependencies

## Installation

1. **Clone the repository:**
    ```
    git clone https://github.com/Voltamon/GoAuthFlow.git
    ```
2. **Navigate to the project directory:**
    ```
    cd GoAuthFlow
    ```
3. **Build the project:**
    ```
    go build -o goauthflow ./main.go
    ```
4. **(Optional) Run with Docker:**
    ```
    docker build -t goauthflow .
    docker run -p 8080:8080 goauthflow
    ```

## Usage

- **Register a new user:**
    ```
    POST /v1/register
    Content-Type: application/json

    {
      "email": "user@example.com",
      "password": "yourpassword"
    }
    ```
- **Login with credentials:**
    ```
    POST /v1/login
    Content-Type: application/json

    {
      "email": "user@example.com",
      "password": "yourpassword"
    }
    ```

## Contributing

- Fork the repository
- Create a new branch (`git checkout -b feat/YourFeature`)
- Commit your changes
- Open a pull request

## License

MIT License. See [LICENSE](LICENSE) for details.

## Maintainer

- Voltamon (GitHub: [Voltamon](https://github.com/Voltamon))
