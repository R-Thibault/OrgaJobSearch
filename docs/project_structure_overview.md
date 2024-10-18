# Project Structure Overview

The project directory structure has been updated to separate the backend and frontend, with the frontend using Next.js and a `/src/app` directory structure:

```
project-root/
    ├── backend/                        # Backend project folder containing Go application code.
    │   ├── cmd/
    │   │   └── main.go                 # Entry point for the Go backend application. This is where the main function resides and the server starts running.
    │   ├── config/
    │   │   ├── db.go                   # Database connection setup, including initializing the database and configuring connection properties.
    │   │   └── config.go               # General configuration settings for the application, such as loading environment variables.
    │   ├── controllers/
    │   │   └── auth_controller.go      # Controller that handles authentication-related logic, including login, registration, etc. Receives HTTP requests and calls appropriate services.
    │   ├── services/
    │   │   └── user_service.go         # Business logic for user management. Handles operations such as creating, updating, and fetching user data.
    │   ├── repository/
    │   │   └── user_repository.go      # Data access layer. Handles database queries related to the user, separating raw data logic from business logic.
    │   ├── models/
    │   │   └── user.go                 # Definition of the User struct that represents the user entity in the system. Includes database field mappings and potentially validation tags.
    │   ├── routes/
    │   │   └── routes.go               # Definition of all the application routes/endpoints. Maps URLs to specific controllers and functions.
    │   ├── middleware/
    │   │   └── auth_middleware.go      # Middleware functions for the application, such as authenticating requests, checking tokens, or handling CORS.
    │   ├── .air.toml                   # Air configuration file for backend live-reloading during development. Useful for hot reloading Go code.
    │   ├── .env                        # Environment variables for the backend, such as database credentials, API keys, or other secrets.
    │   ├── go.mod                      # Go module definition file. Specifies the module path and the dependencies required by the Go backend.
    │   ├── go.sum                      # Checksum file to ensure the integrity of the Go modules. Keeps track of the exact versions of dependencies used.
    │   └── README.md                   # Documentation for the backend project. Typically includes setup instructions, project overview, and usage information.
    ├── frontend/                       # Frontend project folder using Next.js framework.
    │   ├── src/                        # Source directory for all frontend code.
    │   │   ├── app/                    # Next.js app directory, follows the new app router structure.
    │   │   │   ├── page.tsx            # Main entry point for the application.
    │   │   │   ├── layout.tsx          # Shared layout for the app, managing consistent structure across pages.
    │   │   │   ├── global.css          # Global styles for the application, imported in the root layout.
    │   │   │   └── sign-in/            # Folder for sign-in specific pages.
    │   │   │       └── page.tsx        # Page for the sign-in route, allowing routing to the /sign-in page.
    │   │   ├── components/             # Reusable React components.
    │   │   ├── styles/                 # CSS or SCSS files for styling the frontend.
    │   │   │   └── fonts/              # Fonts used in the application.
    │   │   └── lib/                    # Utility functions or custom hooks for the frontend.
    │   ├── public/                     # Static assets like images, favicon, etc., that are publicly accessible.
    │   ├── package.json                # Frontend dependencies, scripts, and project configuration. Manages packages for the client-side using npm/yarn.
    │   ├── next.config.js              # Next.js configuration file.
    │   └── tsconfig.json               # TypeScript configuration for the frontend.
    └── README.md                       # General documentation for the entire project, including both backend and frontend setup instructions.
```
