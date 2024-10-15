# Project Structure Overview

The project directory structure will look like this after adding the frontend:

```
project-root/
    ├── cmd/
    │   └── main.go              # Entry point for the Go backend application. This is where the main function resides and the server starts running.
    ├── config/
    │   ├── db.go                # Database connection setup, including initializing the database and configuring connection properties.
    │   └── config.go            # General configuration settings for the application, such as loading environment variables.
    ├── controllers/
    │   └── auth_controller.go   # Controller that handles authentication-related logic, including login, registration, etc. Receives HTTP requests and calls appropriate services.
    ├── services/
    │   └── user_service.go      # Business logic for user management. Handles operations such as creating, updating, and fetching user data.
    ├── repository/
    │   └── user_repository.go   # Data access layer. Handles database queries related to the user, separating raw data logic from business logic.
    ├── models/
    │   └── user.go              # Definition of the User struct that represents the user entity in the system. Includes database field mappings and potentially validation tags.
    ├── routes/
    │   └── routes.go            # Definition of all the application routes/endpoints. Maps URLs to specific controllers and functions.
    ├── middleware/
    │   └── auth_middleware.go   # Middleware functions for the application, such as authenticating requests, checking tokens, or handling CORS.
    ├── frontend/                # The frontend project folder. Contains everything related to the client-side of the application.
    │   ├── public/              # Static assets like HTML, images, CSS files, etc., that are publicly accessible.
    │   ├── src/                 # Frontend source code. This is where JavaScript/TypeScript files, React components, or other frontend logic resides.
    │   └── package.json         # Frontend dependencies, scripts, and project configuration. Manages packages for the client-side using npm/yarn.
    ├── .air.toml                # Air configuration file for backend live-reloading during development. Useful for hot reloading Go code.
    ├── .env                     # Environment variables for the backend, such as database credentials, API keys, or other secrets.
    ├── go.mod                   # Go module definition file. Specifies the module path and the dependencies required by the Go backend.
    ├── go.sum                   # Checksum file to ensure the integrity of the Go modules. Keeps track of the exact versions of dependencies used.
    └── README.md                # Documentation for the project. Typically includes setup instructions, project overview, and usage information.
```
