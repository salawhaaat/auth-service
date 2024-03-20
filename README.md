# Authentication Service

This repository contains an authentication service implemented in Go using Gin for HTTP routing and MongoDB for data storage. The service utilizes cookies to securely manage authentication tokens, protecting against client-side changes.

## Explanation

- **Secure Token Management**: The service uses HTTP-only and secure cookies to store access and refresh tokens, preventing client-side JavaScript from accessing them. This enhances security by mitigating certain types of attacks, such as cross-site scripting (XSS).

- **Token Verification**: When refreshing tokens, the service ensures that the refresh token provided by the client matches the one stored in the database for the corresponding user ID. This check is performed by comparing the bcrypt-hashed version of the refresh token from the database with the one provided by the client.

- **JWT Claims**: The user ID is embedded in the JWT claims of the access token. This allows the service to extract the user ID from the access token during token verification. The extracted user ID is then used to retrieve the corresponding refresh token from the database.

## Project Structure

```go
auth-service
├── cmd
│   └── main.go
├── go.mod
├── go.sum
└── internal
    ├── database
    │   ├── mongodb
    │   │   └── mongo.go
    │   └── repository.go
    ├── handlers
    │   ├── get_token.go
    │   └── refresh_token.go
    ├── models
    │   └── token.go
    └── services
        └── auth.go
```

- **cmd**: Contains the main entry point of the application.
- **internal**: Contains the core of the application, including database handling, request handlers, models, and services.
  - **database**: Contains database repository interfaces and implementations.
    - **mongodb**: Implements MongoDB repository.
  - **handlers**: Contains HTTP request handlers.
    - **get_token.go**: Handles requests for generating access and refresh tokens.
    - **refresh_token.go**: Handles requests for refreshing access and refresh tokens.
  - **models**: Defines data structures used within the application.
    - **token.go**: Defines the structure of authentication tokens.
  - **services**: Contains business logic.
    - **auth.go**: Provides functions for generating and refreshing authentication tokens.
