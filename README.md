# social-network-notification-service
## Project Summary
This is project about social network that allows users to share content, images, and emotions, and have real-time communication capabilities, while ensuring high performance, security, and scalability using the microservices architecture.

#### Technologies:
- Back-end:
  - Language: Go.
  - Frameworks/Platforms: Gin-Gonic, gRPC, Swagger, JWT, Google-Wire, SQLX, Redis, Zap, WebSocket.
  - Database: MariaDB, MongoDB.
- Front-end:
  - Language: JavaScript.
  - Frameworks/Platforms: React, Tailwind CSS, FireBase.

## The project includes repositories
- [common-service](https://github.com/nhutHao02/social-network-common-service)
- [user-service](https://github.com/nhutHao02/social-network-user-service)
- [tweet-service](https://github.com/nhutHao02/social-network-tweet-service)
- [chat-service](https://github.com/nhutHao02/social-network-chat-service)
- [notification-service](https://github.com/nhutHao02/social-network-notification-service)
- [Front-end-service (in progress)](https://github.com/nhutHao02/)

## This service
This is the service that provides the APIs related to the Notification for the app and handle the real-time notification.

## Project structure
```
.
├── config
│   ├── config.go
│   └── local
│       └── config.yaml
├── database
│   └── database.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   ├── grpc
│   │   ├── http
│   │   │   ├── http_server.go
│   │   │   └── v1
│   │   │       ├── notification_handler.go
│   │   │       └── route.go
│   │   └── server.go
│   ├── application
│   │   ├── imp
│   │   │   └── notification_service_imp.go
│   │   └── notification_service.go
│   ├── domain
│   │   ├── entity
│   │   │   ├── collection_name.go
│   │   │   └── notification.go
│   │   ├── interface
│   │   │   └── notification
│   │   │       └── notification_repository.go
│   │   └── model
│   │       ├── notification.go
│   │       ├── user.go
│   │       └── websocket.go
│   ├── infrastructure
│   │   └── notification
│   │       ├── command_repository.go
│   │       └── query_repository.go
│   ├── wire_gen.go
│   └── wire.go
├── main.go
├── Makefile
├── pkg
│   ├── common
│   │   └── response.go
│   ├── constants
│   │   ├── action_constants.go
│   │   └── constants.go
│   ├── redis
│   │   └── redis.go
│   └── websocket
│       └── websocket.go
├── README.md
└── startup
    └── startup.go
```