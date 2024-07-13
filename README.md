# GoMongo API

This is a RESTful API built using Go, Gin, and MongoDB. The API provides basic CRUD operations for users.

## Prerequisites

- Go 1.17 or higher
- MongoDB 4.4 or higher
- Gin 1.7.4 or higher

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/Ris-Codes/goMongo.git
    ```
2. Navigate to the project directory:
    ```sh
    cd goMongo
    ```
3. Run the API:
    ```sh
    go run main.go
    ```

## API Endpoints

### Users

- **POST `/v1/user/create`** : Create a new user
- **GET `/v1/user/:name`** : Retrieve a user by username
- **GET `/v1/user/getall`** : Retrieve a list of all users
- **PATCH `/v1/user/update`** : Update a user
- **DELETE `/v1/user/delete/:name`** : Delete a user by name

## Running the API

The API runs on port 9090 by default. You can access the API endpoints using a tool like curl or a REST client like Postman.

## MongoDB Connection

The API connects to a local MongoDB instance on port 27017. You can modify the connection string in the `init()` function to connect to a different MongoDB instance.

## Testing
```sh
go test -v ./...
```