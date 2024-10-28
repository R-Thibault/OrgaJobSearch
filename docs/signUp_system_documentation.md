# Sign-Up System Documentation

## Overview

The **Sign-Up System** in the OrgaJobSearch application supports two primary user registration methods:

1. **Support Careers Sign-Up**: Allows admins to generate global invitation links with a JWT token containing an OTP.
2. **Candidate Sign-Up**: Allows support staff to generate individual invitation links for candidates with a JWT token containing a UUID.

This document provides a detailed explanation of these sign-up flows, including the API endpoints, JWT token structure, and validation processes.

---

## Sign-Up Flows

### 1. Support Careers Sign-Up Flow

- **Process**:
  1. An admin generates a **global invitation link** using the `/generate-url` endpoint. The link contains a JWT token with:
     - An **OTP** stored in the database with a type of `"GlobalInvitation"`.
     - An **expiration time** for security purposes.
  2. When a support career clicks the link, they are redirected to the sign-up page where the token is sent to the backend for verification.
  3. The backend verifies:
     - The OTP matches the one stored in the database.
     - The token has not expired.
  4. If verification is successful, the backend responds with the invitation type, and the frontend renders the appropriate sign-up form.
  5. The support career provides their email address and must verify it using another OTP sent via email to complete the registration.

### 2. Candidate Sign-Up Flow

- **Process**:
  1. Support staff generate an **individual invitation** by entering the candidate's email in a frontend form via the `/send-user-invitation` endpoint.
  2. The candidate receives an email with a sign-up link containing a JWT token with:
     - A **UUID** linked to the pre-registered candidate.
     - An **expiration time** for security.
  3. When the candidate clicks the link, the frontend sends the token to the backend for validation.
  4. The backend verifies:
     - The token has not expired.
     - The UUID exists in the database and matches a user who has not completed registration.
  5. If validation is successful, the frontend displays the candidate registration form, allowing the candidate to complete their sign-up.

---

## JWT Token Structure

### 1. Support Careers (Global Invitations)

The JWT token contains the following fields:

```json
{
  "otp": "otp123456",
  "type": "GlobalInvitation",
  "exp": 1698678600 // Expiration time (Unix timestamp)
}
```

- **`otp`**: The OTP generated for this invitation.
- **`type`**: Specifies the type of invitation (`"GlobalInvitation"`).
- **`exp`**: The expiration time of the token (in Unix timestamp format).

### 2. Candidates (Individual Invitations)

The JWT token contains the following fields:

```json
{
  "uuid": "550e8400-e29b-41d4-a716-446655440000",
  "type": "IndividualInvitation",
  "exp": 1698678600 // Expiration time (Unix timestamp)
}
```

- **`uuid`**: The UUID associated with the pre-registered candidate.
- **`type`**: Specifies the type of invitation (`"IndividualInvitation"`).
- **`exp`**: The expiration time of the token (in Unix timestamp format).

---

## API Endpoints

### 1. **Health Check**

- **Method**: `GET`
- **Path**: `/`
- **Description**: Verifies that the server is running and returns a success message.

### 2. **Authentication and Token Management**

- **`/login`**: Authenticates users and generates a JWT token upon successful login.
- **`/logout`**: Deconnection of users and delete the JWT token upon successful logout.
- **`/verify-token`**: Verifies the validity of tokens received from invitation links.

### 3. **Sign-Up and Invitation Management**

- **`/sign-up`**: Handles user registration based on the invitation type and token verification.
- **`/generate-otp`**: Generates an OTP for email verification during sign-up.
- **`/send-user-invitation`**: Allows support staff to invite candidates by sending an individual sign-up link.
- **`/generate-url`**: Admins generate a global invitation link for support careers.

### 4. **OTP Management**

- **`/send-otp`**: Sends an OTP to the user’s email for verification.
- **`/verify-otp`**: Validates the OTP entered by the user.

---

## Security Measures

- **JWT Token Expiration**: Each token has an expiration time to prevent reuse after a set duration.
- **Password Hashing**: Passwords are hashed using a secure hashing algorithm and stored securely.
- **OTP Verification**: OTPs are stored in the database with expiration settings to confirm user identity during registration.

---

## Database and Data Handling

- **OTP Storage**: OTPs are stored with details like user type and expiration time.
- **UUID Registration**: UUIDs are pre-generated for candidates and linked to their email addresses in the database until they complete registration.
- **User Information**: User data is securely stored and only accessed during necessary validation processes.

---

## Additional Considerations

- **Rate Limiting**: Implement rate limiting for endpoints like `/login`, `/send-otp`, and `/sign-up` to prevent abuse.
- **Logging and Monitoring**: Ensure that critical endpoints are logged and monitored for suspicious behavior or errors.

---

## Conclusion

This document outlines the architecture and flow of the Sign-Up System for OrgaJobSearch, explaining the use of JWT tokens, OTP verification, and how the system secures the registration process. For any further details or implementation queries, please refer to the codebase or contact the development team.
