# OrgaJobSearch Documentation

Welcome to **OrgaJobSearch**! This project aims to provide a streamlined job search and management platform, integrating both a Go backend and a modern JavaScript frontend. Below you'll find an overview of how to set up the project, understand the code structure, and follow best practices.

## Table of Contents

- [Introduction](#introduction)
- [Installation Guide](#installation-guide)
- [Project Structure](#project-structure)
- [HTTP Request Flow](#http-request-flow)
- [Code Conventions](#code-conventions)
- [Getting Started](#getting-started)
- [License](#license)

## Introduction

**OrgaJobSearch** is a Work In Progress (WIP) aimed at simplifying job searching processes with enhanced backend and frontend capabilities.

The project uses:

- **Golang** for backend development
- **React** (or another modern JavaScript framework) for the frontend
- **PostgreSQL** for data persistence
- Additional tools such as **Air**, **Viper**, **Lefthook**, and **Gorm** to aid development and maintainability

## Installation Guide

For a detailed guide on how to get started with the installation of this project, including prerequisites and steps to set up the development environment, please refer to [starter.md](./docs/starter.md).

## Project Structure

The project follows a layered structure, designed to promote separation of concerns and maintainability. Each component, from controllers to models, serves a distinct purpose. To understand the directory structure in detail, visit [project_structure_overview.md](./docs/project_structure_overview.md).

## HTTP Request Flow

This document outlines how HTTP requests, such as creating a user, flow through different parts of the application—starting from routing, reaching the service layer, interacting with the repository, and ending in the database. This is crucial for understanding how different parts of the application interact. For more information, refer to [http_request_flow.md](./docs/http_request_flow.md).

## Code Conventions

To maintain consistency in the codebase, follow the provided code conventions. This is especially useful when new developers join the project. Conventions include rules for naming, formatting, and structuring code. For full details, see [project_code_convention.md](./docs/project_code_conventions.md).

## Getting Started

1. **Clone the Repository**
   ```sh
   git clone https://github.com/R-Thibault/OrgaJobSearch
   cd OrgaJobSearch
   ```
2. **Set Up Dependencies**
   Make sure you have [Golang 1.22 or higher](https://golang.org/dl/) installed, and set up a PostgreSQL database.

3. **Environment Variables**
   Copy `.env.sample` to `.env` and fill in the appropriate environment variables.

4. **Run the Application**
   - To run in development mode with live reloading:
     ```sh
     air
     ```
   - For a regular run:
     ```sh
     go run cmd/main.go
     ```

## License

This project is licensed under the MIT License.
