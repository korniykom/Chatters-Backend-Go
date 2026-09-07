# Chatterz Backend

Backend for a chat application built with Go and designed as a set of independent services.

## Tech Stack

- Go
- Chi
- PostgreSQL
- pgx
- Goose
- Docker
- Kubernetes

## Architecture

```mermaid
flowchart LR
    Client --> UserService
    Client --> ChatService

    UserService --> UserDB[(PostgreSQL)]
    ChatService --> ChatDB[(PostgreSQL)]

    ChatService --> NotificationService
    NotificationService --> Push[Push Notifications]

    subgraph Chatterz
        UserService[User Service]
        ChatService[Chat Service]
        NotificationService[Notification Service]
    end