# OrgaJobSearch: Project Specifications

## Table of Contents

1. [Project Overview](#project-overview)
2. [Objectives](#objectives)
3. [Target Users](#target-users)
4. [Functional Requirements](#functional-requirements)
5. [Non-Functional Requirements](#non-functional-requirements)
6. [Technology Stack](#technology-stack)
7. [System Architecture](#system-architecture)
8. [Detailed Features](#detailed-features)
9. [Deployment Requirements](#deployment-requirements)

## Project Overview

**OrgaJobSearch** is a platform designed to simplify job search and management processes, providing users with an intuitive experience for finding, tracking, and managing job applications. It offers different interfaces for public users and enterprise clients, including a back-office for managing company-specific needs.

## Objectives

1. **User-Friendly Job Search**: Provide users with a simple and effective interface for discovering and managing job opportunities.
2. **Enterprise Integration**: Allow companies to create customized job posts and manage their job listings through a dedicated back-office.
3. **Scalability and Maintainability**: Create a well-structured system that is easy to maintain, extend, and scale for new features or additional clients.

## Target Users

- **Job Seekers**: Individuals looking for job opportunities.
- **Enterprise Clients**: HR departments or hiring managers seeking candidates.
- **Admin Users**: Internal users managing the platform, including overseeing job posts, analytics, and security.

## Functional Requirements

### User Features

1. **Sign Up & Login**:

   - Users can register or login using email and password.
   - Implement password hashing and secure JWT-based authentication.
   - Use validation rules for secure and proper credential format (e.g., regex checks).

2. **Job Search**:

   - Users can search for available job opportunities using filters such as **location**, **job title**, **industry**, and **company**.

3. **Job Application Management**:

   - Users can apply for jobs and track application statuses.
   - Users can view a history of their submitted applications.

4. **Profile Management**:
   - Users can manage personal details such as **name**, **resume**, and **contact information**.

### Enterprise Client Features

1. **Job Post Management**:

   - Enterprise users can create, update, and delete job postings.
   - Ability to view analytics for job posts (e.g., applications per job).

2. **User Role Management**:
   - Assign different roles within an enterprise for **collaborative job posting**.

### Admin Features

1. **Platform Monitoring**:

   - Admin users can monitor the activities of regular users and enterprise clients.
   - Admins can block abusive users or job postings.

2. **Analytics Dashboard**:
   - Admins can view key metrics (e.g., most active users, trending job postings).

## Non-Functional Requirements

1. **Security**:

   - User credentials should be hashed and secured (using **Argon2**).
   - All sensitive operations should use HTTPS and secure cookies.

2. **Performance**:

   - The application should load pages within **2 seconds** for most interactions.
   - API response time should ideally be under **500 ms**.

3. **Scalability**:

   - The architecture should allow horizontal scaling to accommodate growth in user base.

4. **Maintainability**:

   - Code should follow a layered structure (controller, service, repository).
   - Each layer should have sufficient unit and integration tests to ensure stability.

5. **Accessibility**:
   - Application must follow basic **WCAG** accessibility guidelines to ensure all users, including those with disabilities, can use the platform.

## Technology Stack

- **Backend**: Golang (using **Gin Framework**)
- **Frontend**: Next.js (with **React** for component management)
- **Database**: PostgreSQL
- **ORM**: Gorm
- **Authentication**: JWT-based user sessions
- **Configuration**: Viper for configuration management
- **Deployment**: Docker for containerization, CI/CD with GitHub Actions
- **Development Tools**: Air (live-reloading), Lefthook (Git hooks for quality checks), Mockery (test mocking)

## System Architecture

- **Frontend (CSR/SSR)**: A hybrid approach using **Next.js** for Server-Side Rendering (SSR) to boost SEO and performance.
- **Backend RESTful API**: The backend is built in Golang using **Gin** to handle business logic and serve data to the frontend.
- **PostgreSQL Database**: All data for users, job applications, and roles are stored in PostgreSQL.
- **Docker**: Docker containers for all microservices to ensure consistency across environments.

## Detailed Features

1. **User Registration and Login**

   - Passwords are validated using **Regex** to enforce security policies.
   - Password hashing using **Argon2** to provide strong resistance to brute force attacks.

2. **Job Postings**

   - Enterprises can create detailed job postings with categories, locations, and benefits.

3. **Application Tracking**

   - Users can view the status of their applications (e.g., **Under Review**, **Accepted**, **Rejected**).

4. **Search and Filtering**

   - Users can search for jobs with keywords and apply various filters like location or job type.

5. **Session Management**
   - Cookies used for managing JWTs (secure and HTTP-only).
   - Implement session timeouts and revalidation for expired sessions.

## Deployment Requirements

1. **Environment Variables**: All secrets (DB credentials, JWT key) should be managed through environment variables.
2. **CI/CD Pipeline**: Build and deploy to production using **GitHub Actions**.
3. **Monitoring and Logging**: Use **Prometheus** and **Grafana** for monitoring server health, and **ELK stack** for centralized logging.

## Testing

- **Unit Tests**: Unit tests for each component (services, controllers) using **Testify** and mocks generated with **Mockery**.
- **Integration Tests**: Verify how different parts of the system interact.
- **Documentation for Testing**: To understand how to add new tests or run tests in this project, refer to [project_testing_backend.md](./docs/project_testing_backend.md).

## Summary

This specification provides a detailed overview of the requirements, both functional and non-functional, for the **OrgaJobSearch** platform. Each module and its responsibilities are clearly outlined to ensure that the development is well-structured, maintainable, and scalable for future features and expansions.
