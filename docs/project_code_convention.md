# Go Project Code Conventions

This document outlines the coding conventions for the Go project. The goal is to ensure that the codebase remains consistent, readable, and maintainable, so that onboarding new developers is straightforward. This guide covers project structure, naming conventions, best practices, and more.

## For Backend

### Folder Purposes

- **cmd/**: Contains the application entry point (`main.go`).
- **config/**: Configuration utilities and database connection setup.
- **controllers/**: Handles incoming HTTP requests and interacts with the service layer.
- **services/**: Contains the core business logic, which is reusable across controllers.
- **repository/**: Handles database interactions.
- **models/**: Defines data structures and mappings to the database.
- **routes/**: Defines API routes and maps them to controllers.
- **middleware/**: Contains middleware logic, such as authentication.
- **utils/**: Houses helper functions (e.g., hashing passwords).
- **frontend/**: Contains frontend code (e.g., React/Next.js).

### Naming Conventions

- **Packages**: Package names should be lowercase and singular (e.g., `repository`, `service`). Avoid underscores or camelCase.
- **Files**: Use snake_case for file names (e.g., `user_repository.go`, `auth_middleware.go`).
- **Functions**: Function names should be descriptive and use camelCase (e.g., `GetUserByEmail`, `RegisterUser`).
- **Variables**: Use camelCase for variable names, and keep them descriptive (`userRepo`, `hashedPassword`). Use short variable names (e.g., `i`, `c`) for loops or very limited scopes.
- **Constants**: Use uppercase with underscores (`MAX_ATTEMPTS`, `JWT_SECRET`).
- **Interfaces**: Should describe behavior and end in `-er` (e.g., `Hasher`, `UserRepositoryInterface`).
- **Structs**: Struct names should be capitalized and descriptive (e.g., `User`, `Credentials`).

### Code Formatting

- **Indentation**: Use tabs (Go default).
- **Line Length**: Limit lines to 100 characters when possible.
- **Imports**: Organize imports into three groups, separated by blank lines: standard library, third-party packages, and project-specific packages.

  ```go
  import (
      "fmt"
      "net/http"

      "github.com/gin-gonic/gin"

      "github.com/R-Thibault/OrgaJobSearch/backend/repository"
  )
  ```

- **Comments**: Use comments to explain why a particular decision was made, rather than what the code is doing. Public functions and structs should have doc comments (`// FunctionName does something...`).

## Best Practices for backend

### Controllers

- Controllers should only handle HTTP request/response logic.
- Pass requests to the service layer for actual business logic.

### Services

- Implement all business logic in the service layer.
- Services should be stateless; any necessary state should be fetched from the repository.

### Repository

- Repositories should only handle database access (e.g., CRUD operations).
- Always return clear errors from repository functions so services can handle them appropriately.

### Middleware

- Middleware should be reusable and stateless.
- Keep middleware functions simple and focused on one responsibility (e.g., checking authentication).

### Error Handling

- Always handle errors explicitly and return descriptive error messages.
- Use `errors.Wrap()` (from the `github.com/pkg/errors` package) when returning errors to provide more context.
- Avoid panics in your code; instead, return errors.

### Tests

- **Unit Tests**: Test individual functions in isolation (mocks are great for repositories).
- **Integration Tests**: Test the integration between layers (e.g., service and repository).
- Use descriptive names for test functions (`TestRegisterUser_Success`, `TestSignIn_InvalidPassword`).
- Group related tests in the same file for better organization.

## Security Best Practices

- **Passwords**: Always hash passwords before storing them. Use strong algorithms like Argon2.
- **Environment Variables**: Store sensitive configuration like database credentials in `.env` files, and never commit `.env` files to version control.
- **CORS**: Be cautious with CORS settings to avoid unnecessary exposure.
- **JWT Secret**: Keep the JWT secret secure and rotate periodically.

## Branching and Commit Messages

- **Branching**: Use a branching model like GitFlow for organized collaboration.
- **Commit Messages**: Write descriptive commit messages that follow a conventional format:
  - `feat: add user signup endpoint`
  - `fix: correct password hashing logic`
  - `refactor: split user service into multiple functions`

## Development Workflow

- **Live Reloading**: Use [Air](https://github.com/cosmtrek/air) for live reloading during backend development.
- **Code Reviews**: All changes should be reviewed before merging into the main branch. This ensures code quality and knowledge sharing.

## Tools

- **Mockery**: Use [Mockery](https://github.com/vektra/mockery) for generating mocks for interfaces in unit tests.
- **Go Lint**: Run `golangci-lint` to catch common issues and enforce code style guidelines.
- **VS Code Settings**: Use the `gopls` extension for Go language support, and configure it to format code on save.
