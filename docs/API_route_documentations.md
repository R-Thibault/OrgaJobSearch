# API Routes Documentation

This document provides a summary of the API endpoints available for frontend developers working with our Go backend. Each endpoint includes a description, HTTP method, and example payload when relevant. Protected routes require authentication.

---

## Public Routes

These routes do not require authentication.

### 1. Login

- **Endpoint**: `/login`
- **Method**: `POST`
- **Description**: Authenticates a user based on provided credentials.
- **Payload**:
  ```json
  {
    "email": "user@example.com",
    "password": "your_password"
  }
  ```
- **Model**: `Credentials`
  - **Fields**:
    - `Email` (string): The user's email address.
    - `Password` (string): The user's password.
- **Response**: Returns a JWT token if authentication is successful.

### 2. Logout

- **Endpoint**: `/logout`
- **Method**: `POST`
- **Description**: Logs the user out by invalidating the session token.
- **Payload**: None

### 3. Sign-Up

- **Endpoint**: `/sign-up`
- **Method**: `POST`
- **Description**: Registers a new user account.
- **Payload**:
  ```json
  {
    "email": "new_user@example.com",
    "lastName": "Doe",
    "firstName": "John",
    "password": "your_password",
    "confirmPassword": "your_password"
  }
  ```
- **Model**: `Credentials`
  - **Fields**:
    - `Email` (string): The user's email address.
    - `LastName` (string): The user's last name.
    - `FirstName` (string): The user's first name.
    - `Password` (string): The user's password.
    - `ConfirmPassword` (string): The password confirmation for validation.
- **Response**: Returns a success message if registration is successful or an error if validation fails (e.g., passwords do not match).

### 4. Generate OTP for Sign-Up

- **Endpoint**: `/generate-otp`
- **Method**: `POST`
- **Description**: Generates an OTP (One-Time Password) for account verification during sign-up.
- **Payload**:
  ```json
  {
    "email": "user@example.com"
  }
  ```

### 5. Send OTP

- **Endpoint**: `/send-otp`
- **Method**: `POST`
- **Description**: Sends an OTP to the user's email. **This route is not used on the frontend; it is called internally by the backend.**

### 6. Verify OTP

- **Endpoint**: `/verify-otp`
- **Method**: `POST`
- **Description**: Verifies the OTP entered by the user for account activation.
- **Payload**:
  ```json
  {
    "email": "user@example.com",
    "otp": "123456"
  }
  ```

---

## Protected Routes

These routes require authentication. The frontend must provide a valid token in the request header.

### 1. Get My Profile

- **Endpoint**: `/me`
- **Method**: `GET`
- **Description**: Retrieves the profile information of the authenticated user.
- **Headers**: `Authorization: Bearer <token>`

### 2. Update User Profile

- **Endpoint**: `/update-user`
- **Method**: `POST`
- **Description**: Updates the profile information for the authenticated user.
- **Headers**: `Authorization: Bearer <token>`
- **Payload**:
  ```json
  {
    "userFirstName": "John",
    "userLastName": "Doe",
    "email": "john.doe@example.com"
  }
  ```
- **Model**: `UserProfileUpdate`
  - **Fields**:
    - `UserFirstName` (string): The user's first name.
    - `UserLastName` (string): The user's last name.
    - `Email` (string): The user's email address.
- **Response**: Returns a success message if the profile is updated successfully or an error message if there is an issue with the update.

### 3. Create Application

- **Endpoint**: `/create-application`
- **Method**: `POST`
- **Description**: Submits a new job application linked to the authenticated user.
- **Headers**: `Authorization: Bearer <token>`
- **Payload**:
  ```json
  {
    "userID": 1,
    "url": "https://company.com/job-listing",
    "title": "Software Engineer",
    "company": "Tech Corp",
    "location": "Remote",
    "description": "A full-time software engineering position.",
    "salary": "100000",
    "jobType": "Full-Time",
    "applied": true,
    "response": false,
    "followUp": false
  }
  ```
- **Model**: `Application`
  - **Fields**:
    - `UserID` (uint): The ID of the user submitting the application (automatically set based on the authenticated user).
    - `Url` (string): The URL to the job listing or application page. Required.
    - `Title` (string): The title of the job. Required, max 255 characters.
    - `Company` (string): The name of the company offering the job, max 255 characters.
    - `Location` (string): The job location, max 255 characters.
    - `Description` (string): A description of the job.
    - `Salary` (string): The offered salary for the job, max 255 characters.
    - `JobType` (string): The type of job (e.g., Full-Time, Part-Time), max 255 characters.
    - `Applied` (bool): Indicates whether the user has applied to the job. Defaults to `true`.
    - `Response` (bool): Indicates whether a response has been received from the company. Defaults to `false`.
    - `FollowUp` (bool): Indicates if a follow-up has been done. Defaults to `false`.
- **Response**: Returns a success message with the application ID if the creation is successful or an error message if there’s an issue with the submission.

### 4. Get Applications by User ID

- **Endpoint**: `/get-applications-by-user`
- **Method**: `POST`
- **Description**: Fetches all applications associated with the authenticated user.
- **Headers**: `Authorization: Bearer <token>`
- **Payload**: None

### 5. Update Application

- **Endpoint**: `/update-application`
- **Method**: `POST`
- **Description**: Updates an existing application for the authenticated user. Only the fields that require updating should be provided, along with the application ID.
- **Headers**: `Authorization: Bearer <token>`
- **Payload**:

  ```json
  {
    "applicationID": 123,
    "title": "Senior Software Engineer",
    "company": "Tech Innovations",
    "location": "New York",
    "description": "An updated job description with new responsibilities.",
    "salary": "120000",
    "jobType": "Contract",
    "applied": true,
    "response": true,
    "followUp": false
  }
  ```

- **Model**: `Application` (Partial)
  - **Fields**:
    - `ApplicationID` (uint): The ID of the application to be updated.
    - **Optional Fields**: Only include fields from `Application` that need to be updated.
- **Response**: Returns a success message if the application is updated successfully or an error message if there is an issue with the update.

### 6. Soft Delete Application

- **Endpoint**: `/delete-application`
- **Method**: `POST`
- **Description**: Soft deletes an application using Gorm's functionality.
- **Headers**: `Authorization: Bearer <token>`
- **Payload**:
  ```json
  {
    "applicationID": 123
  }
  ```
- **Response**: Returns a success message if the application is deleted successfully or an error message if there is an issue with the deletion.

### 7. Update Application Status

- **Endpoint**: `/update-application-status`
- **Method**: `POST`
- **Description**: Updates only the status of an application, typically displayed on the dashboard.
- **Headers**: `Authorization: Bearer <token>`
- **Payload**:
  ```json
  {
    "applicationID": 123,
    "status": "Interview Scheduled"
  }
  ```
- **Response**: Returns a success message if the application status is updated successfully or an error message if there is an issue with the update.

### 8. Get Application by ID

- **Endpoint**: `/get-application-by-id`
- **Method**: `POST`
- **Description**: Retrieves the details of a specific application based on its ID. Useful for displaying a single application on a detailed page.
- **Headers**: `Authorization: Bearer <token>`
- **Payload**:

  ```json
  {
    "applicationID": 123
  }
  ```

- **Response**:
  ```json
  {
    "applicationID": 123,
    "userID": 1,
    "url": "https://company.com/job-listing",
    "title": "Software Engineer",
    "company": "Tech Corp",
    "location": "Remote",
    "description": "A full-time software engineering position.",
    "salary": "100000",
    "jobType": "Full-Time",
    "applied": true,
    "response": false,
    "followUp": false,
    "createdAt": "2024-01-01T12:00:00Z",
    "updatedAt": "2024-01-02T12:00:00Z"
  }
  ```

---

### Additional Information

- **Authorization**: All protected routes require an `Authorization` header with a valid JWT token:
  ```
  Authorization: Bearer <token>
  ```
- **Error Handling**: Ensure to check for possible validation errors and handle responses appropriately on the frontend.

---

> **Note from the Creator**: Hello Lucie! I hope this document makes your work with the frontend a bit easier.
> If you have any questions about these routes or need further clarification, don't hesitate to reach out.
> Wishing you all the best with the project! 😊

> **Note from the Creator**: Hello Lucie! I hope this document makes your work with the frontend a bit easier.
> If you have any questions about these routes or need further clarification, don't hesitate to reach out.
> You can also check [this link](https://chatgpt.com/share/672b35fa-e3cc-800d-94a0-7f5c4dcb1a6b) for the original prompt.
> Wishing you all the best with the project! 😊
>
> – ChatGPT, the creator of this document
