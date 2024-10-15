# Creating the Go Boilerplate

This guide will walk you through the steps to create the Go boilerplate using the required Go modules. Below is a list of the modules you'll need to install to set up the project with Gin, Gorm, Air, Lefthook, and Viper.

## Steps to Create the Boilerplate

1. **Initialize a New Go Module**

   Start by creating a new directory for your project and initializing a Go module:

   ```bash
   mkdir my-go-boilerplate
   cd my-go-boilerplate
   go mod init github.com/yourusername/my-go-boilerplate
   ```

2. **Install Necessary Go Modules**

   To set up the boilerplate, install the following Go modules:

   - **Gin (HTTP Framework)**: Gin is used for building fast, flexible web applications and APIs.

     ```bash
     go get -u github.com/gin-gonic/gin
     ```

   - **Air (Development Tool)**: Air is a live reload utility for Go, useful during development.

     ```bash
     go install github.com/air-verse/air@latest
     export PATH=$PATH:$(go env GOPATH)/bin
     ```

   - **Lefthook (Git Hooks Manager)**: Lefthook is used for managing Git hooks to ensure quality control.

     ```bash
     go install github.com/evilmartians/lefthook@latest
     export PATH=$PATH:$HOME/go/bin
     ```

   - **Viper (Configuration Management)**: Viper is a complete configuration solution for Go applications, useful for handling configuration files and environment variables.

     ```bash
     go get github.com/spf13/viper
     ```

   - **Gorm (ORM Library)**: Gorm is an ORM library for Go, used to interact with databases.

     ```bash
     go get -u gorm.io/gorm
     ```

   - **Gorm PostgreSQL Driver**: To connect Gorm to a PostgreSQL database, install the PostgreSQL driver.

     ```bash
     go get -u gorm.io/driver/postgres
     ```

   To set the connexion, install the following Go modules :

   - **JWT-Go JWT Token manager** : To create and manage JWT Token

   ```bash
    go get github.com/dgrijalva/jwt-go
   ```

3. **Verify Installation**

   After installing all dependencies, make sure they are properly listed in your `go.mod` file:

   ```bash
   go mod tidy
   ```

   This command will ensure that all dependencies are correctly added and any unnecessary packages are removed.

## Additional Setup

- Make sure to include `Air` in your PATH to use it directly from the command line.

  ```bash
  export PATH=$PATH:$(go env GOPATH)/bin
  ```

- Create an `.env` file to store environment variables and configuration settings, and consider using `Viper` to manage them effectively in your application.
