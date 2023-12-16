# Vshare-Chatting-Backend
Implementing chatting as a separate microservice . Using Go Lang 🤓

# Contribution Guidelines

Thank you for considering contributing to our project! Before you get started, please read through these guidelines to ensure a smooth and collaborative development process.

1. **While implementing, follow the Observer Pattern**


### DB Diagram
https://dbdiagram.io/d/Individual-Chatting-digram-6572001656d8064ca099a2c9

### Folder Structure

Ensure that your code follows the recommended folder structure for the Observer pattern:

# cmd/

This directory contains the main application entry point.

- **main.go:** The primary entry point for the microservice.

# internal/

This directory holds internal packages that should not be imported by external packages.

- **chat/:** Package for the chat microservice.
  - **chatroom.go:** Implementation of the `ChatRoom` struct.
  - **user.go:** Definition of the `User` interface and `ConcreteUser` implementation.

# pkg/

External packages that can be used by other projects.

# api/

API definitions and documentation.

# docs/

Documentation and project-related files.

# scripts/

Useful scripts for automation.

# test/

Test files.

# vendor/

Vendor directory for dependencies.

# GitHub-Specific Files

- **.gitignore:** Specifies files and directories that should be ignored by Git.
- **README.md:** Project documentation.



# Other Directories

- **config/:** Configuration files.
- **scripts/:** Additional scripts for deployment, CI/CD, etc.
- **deploy/:** Deployment configurations.
