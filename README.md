# Vshare-ChatApp Backend 

This repository houses the backend of a modern chat application, designed with a microservices architecture for scalability and maintainability. The application is divided into separate microservices, each responsible for specific functionalities:

- **Auth Service:** Handles user authentication and authorization.
- **Chat Service:** Manages individual and group chat functionalities.
- **Group Service:** Facilitates group creation and management.
- **Video Call Service:** Supports real-time video call features.

## Key Features
- Dockerized microservices for easy deployment and scalability.
- Kubernetes configurations for container orchestration in a cloud environment.
- Database migrations for version-controlled schema changes.
- Scripting for deployment automation and seed data initialization.

## Project Structure
- `/auth-service`, `/chat-service`, `/group-service`, `/video-call-service`: Individual microservices.
- `/common`: Shared utilities and constants.
- `/migrations`: Database migration scripts.
- `/scripts`: Deployment and data seeding scripts.
- `/kubernetes`: Kubernetes deployment and service configurations.
- `docker-compose.yml`: Docker Compose file for local development.

## Getting Started
1. Clone the repository.
2. Set up and configure Docker and Kubernetes.
3. Run `docker-compose up` for local development.
4. Explore microservices individually and interact with the API.

## Contributing
Contributions are welcome! Feel free to open issues or pull requests.


### Database Diagram

Below is the database diagram illustrating the structure of our chat app backend:
[![Database Diagram]([https://www.google.com/url?sa=i&url=https%3A%2F%2Fillustoon.com%2F%3Fid%3D7826&psig=AOvVaw3rgRCk8eiQajBU4LaJjuqv&ust=1702903468915000&source=images&cd=vfe&opi=89978449&ved=0CBIQjRxqFwoTCKjK3_a_loMDFQAAAAAdAAAAABAY]))]([https://dbdiagram.io/d/Individual-Chatting-digram-6572001656d8064ca099a2c9])

