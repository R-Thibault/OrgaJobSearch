# Testing Guide for OrgaJobSearch Project

This guide aims to help you understand the testing setup used in the OrgaJobSearch project and walk you through creating and running tests effectively. It covers the tools and packages required and the structure for adding unit tests to maintain high code quality.

## Overview of Testing in Go

Testing in Go is primarily done using the built-in `testing` package, which provides a simple yet powerful way to write unit tests. The `testify` package is also used extensively in this project for more expressive assertions and mocking functionality.

## Tools and Packages Used

- **`testing` Package**: The standard Go package for writing unit tests. It allows you to create functions named with the `Test` prefix to validate your code.
- **Testify (`github.com/stretchr/testify`)**: A third-party testing toolkit that provides helpful assertion methods and mocking capabilities for easier testing.
- **Mockery (`github.com/vektra/mockery`)**: A tool used to auto-generate mocks for interfaces, making it easier to isolate components during testing.

## Setting Up Testing in This Project

### Step 1: Creating Test Files

In Go, test files should be placed alongside the code they are testing. Test files should have the `_test.go` suffix, which ensures that the Go tool recognizes them as test files. For example, if you are testing `user_service.go`, the test file should be named `user_service_test.go`.

### Step 2: Structuring Unit Tests

Tests in Go are written as functions within the `_test.go` files. Each test function should start with the word `Test`. For instance, to write tests for the `RegisterUser` function of the `UserService`, you could create multiple test functions, such as:

- `TestRegisterUser_UserAlreadyExists`
- `TestRegisterUser_PasswordRegexCheckFail`
- `TestRegisterUser_PasswordRegexCheckPass`

Each function should have the following structure:

```go
func TestFunctionName(t *testing.T) {
    // Setup the environment and any mocks needed

    // Call the function to test

    // Validate the results using assertions
}
```

### Step 3: Writing Tests with Testify

We use `testify` for both assertions and mocking.

- **Assertions**: Testify assertions help to clearly express what is expected in the test. Examples include `assert.NoError(t, err)` or `assert.Equal(t, expected, actual)`.
- **Mocking**: Mocking allows you to isolate the component being tested by providing mock dependencies. This project uses `mockery` to generate mocks for the interfaces.

### Step 4: Generating Mocks with Mockery

To generate mocks for your interfaces, you use the `mockery` tool. Here is how to generate mocks for a specific interface:

```bash
mockery --name UserRepositoryInterface --dir ./repository --output ./repository/mocks
```

This command generates a mock for the `UserRepositoryInterface` located in the `./repository` directory and stores it in `./repository/mocks`.

### Step 5: Writing a Test Case Example

Here’s an example from `user_service_test.go`:

```go
func TestRegisterUser_UserAlreadyExists(t *testing.T) {
    mockRepo := new(mockRepo.UserRepositoryInterface)
    mockHashingService := new(mockUtil.HashingServiceInterface)
    userService := services.NewUserService(mockRepo, mockHashingService)

    creds := models.Credentials{
        Email:    "existing@example.com",
        Password: "superPassword1!",
    }

    // Setup mock expectations to simulate an existing user
    mockRepo.On("GetUserByEmail", creds.Email).Return(&models.User{Email: creds.Email}, nil)

    // Execute the function
    err := userService.RegisterUser(creds)

    // Assertions
    assert.Error(t, err)
    assert.Equal(t, "user already exists", err.Error())
    mockRepo.AssertExpectations(t)
}
```

### Step 6: Running the Tests

To run tests in Go, use the following command in the project root:

```bash
go test ./...
```

This will run all the tests in the project, including all subdirectories.

For a more detailed output, you can run:

```bash
go test -v ./...
```

This command provides verbose output, showing each test as it is executed.

### Step 7: Common Commands for Testing

- **Run All Tests**: `go test ./...`
- **Run Tests Verbosely**: `go test -v ./...`
- **Run a Specific Test File**: `go test -v ./services/user_service_test.go`

### Step 8: Generating Mocks for Testing

If you need to update or add new mocks, make sure you use `mockery` to generate mocks for your new interfaces:

```bash
mockery --name InterfaceName --dir ./path/to/interface --output ./path/to/mocks
```

Ensure the generated mock files are committed so that other developers do not need to regenerate them when they clone the project.

## Example Test Cases

### Testing User Service

- The `TestRegisterUser_UserAlreadyExists` checks if the registration function returns an error when attempting to register a user that already exists in the database.
- The `TestRegisterUser_PasswordRegexCheckFail` validates that passwords not matching the required regex pattern will cause the registration to fail.
- The `TestRegisterUser_PasswordRegexCheckPass` confirms that valid credentials are accepted and the user is registered successfully.

### Testing Guidelines

1. **Unit Tests**: Write unit tests for every function in the service and repository layers.
2. **Mocks**: Use `mockery` to create mocks of repository and service interfaces.
3. **Isolation**: Ensure each test only tests one functionality. Use mocks to isolate the function being tested.

### Tips for Writing Effective Tests

- **Arrange, Act, Assert**: Follow the `Arrange, Act, Assert` pattern to organize your tests. First, set up the necessary data (Arrange), then execute the code being tested (Act), and finally verify the results (Assert).
- **Use Mocks Effectively**: Mock all external dependencies so your tests focus only on the function's internal logic.
- **Meaningful Test Names**: Name your test functions to clearly indicate what is being tested and what is expected, e.g., `TestRegisterUser_UserAlreadyExists`.

## Running Tests Automatically with CI

Consider adding Continuous Integration (CI) to automate testing whenever new changes are pushed. You could use services like **GitHub Actions** or **GitLab CI** to run `go test ./...` on every push or pull request.

Refer to the [Project Code Convention Guide](project_code_convention.md) for additional best practices that also apply to writing test cases.
