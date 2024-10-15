# HTTP Request Flow: Creating a User

This guide describes the flow of an HTTP request to **create a new user** in a Go application that uses **controller**, **service**, and **repository** layers. The request will go through several layers, each with a specific responsibility, from handling the incoming request to interacting with the database and returning a response.

## Overview of the Flow

1. **HTTP Request (from client)** →
2. **Router** (routes the request to the right controller) →
3. **Controller** (handles the HTTP request/response) →
4. **Service** (applies business logic) →
5. **Repository** (interacts with the database) →
6. **Database** (persists the data) →
7. **Response is returned through each layer back to the client**

## Summary of Each Layer's Role

- **Client**: Sends the HTTP request to create a user.
- **Router**: Directs the request to the correct controller based on the URL and HTTP method.
- **Controller**: Handles the request, validates the input, and calls the service.
- **Service**: Applies business rules, hashes the password, and calls the repository.
- **Repository**: Manages data persistence and interacts with the database.
- **Database**: Stores the data.
- **Response Path**: Each layer passes the response back up until it reaches the client.

## Visual Representation

Here is a textual visualization of the flow:

```plaintext
 +---------------+                     +----------------+
 |   Client      |   POST /api/users    |                |
 | (HTTP Request)| ------------------>  |  Router (Gin)  |
 +---------------+                     +----------------+
                                                |
                                                v
                                      +--------------------+
                                      |    UserController  |
                                      |  (CreateUser func) |
                                      +--------------------+
                                                |
                                                v
                                      +--------------------+
                                      |    UserService     |
                                      |  (RegisterUser)    |
                                      +--------------------+
                                                |
                                                v
                                      +--------------------+
                                      |   UserRepository   |
                                      |   (SaveUser func)  |
                                      +--------------------+
                                                |
                                                v
                                      +--------------------+
                                      |     Database       |
                                      | (GORM/SQL Backend) |
                                      +--------------------+
```

The flow moves **down** through each layer to the database and **back up** to provide the response to the client.
