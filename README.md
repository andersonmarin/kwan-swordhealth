# Kwan Sword Health Test

This project is composed of two main services—api and notifier—which communicate via messaging (NATS) and persist data
using a MySQL (MariaDB) database. The system primarily focuses on the core functionalities: creating and listing tasks
performed by technicians.

## Architecture

```
├── cmd/            # Entry points for the services
│ ├── api/          # REST API service (using Echo and JWT)
│ └── notifier/     # Message consumer service (using NATS)
├── internal/       # Infrastructure layer including adapters
│ ├── broker/       # NATS integration and notification service
│ ├── http/         # HTTP handlers and middleware
│ └── mysql/        # Repository implementations for MySQL
├── pkg/            # Domain layer, containing business rules and use cases
│ ├── task/         # Task domain entity, repository, and use cases
│ ├── user/         # User domain definitions and repository interface
│ └── notification/ # Notification interface and mock implementation
├── sql/            # SQL scripts for database initialization and sample data load
├── docker-compose.yml
└── README.md
```

### API

API Service: Receives HTTP requests (secured with JWT via Echo) and publishes events through NATS.

#### Key Dependencies

- NATS: For publishing events.
- MySQL: For database connection.

#### Environment Variables

- DATABASE_URL
- NATS_URL
- JWT_KEY
- ADDRESS (defines the server listening port)

### Notifier

Subscribes to NATS events (specifically, notification.notifyTaskPerformed) and processes
notifications—currently logging the event details.

#### Key Dependency

- NATS: To subscribe and handle notifications.

#### Environment Variable

- NATS_URL

## Running Locally with Docker Compose

### Prerequisites

- Docker
- Docker Compose

### Instructions

#### 1. Clone the Repository, ensure you have the project source on your local machine.

#### 2. Start the Services, open your terminal and run:

```shell
docker-compose up --build
```

This command will:

- Build and start the api and notifier services.
- Start the mariadb (MySQL) container with the initial schema and seed data loaded via the sql/init.sql script.
- Launch the nats container for messaging.

#### 3. Accessing the API

The API service will be available at:

```
http://localhost:8080/
```

##### Available endpoints

###### Create Task

```shell
curl --request POST \
  --url http://localhost:8080/task \
  --header 'Authorization: Bearer {{JWT}}' \
  --header 'Content-Type: application/json' \
  --data '{
	"summary": "Some task",
	"performedAt": "2025-04-10T12:42:00.000Z"
}'
```

###### List Task

```shell
curl --request GET \
  --url http://localhost:8080/task \
  --header 'Authorization: Bearer {{JWT}}'
```

##### Test Users

- **tech1**
    - JWT: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.tjVEMiS5O2yNzclwLdaZ-FuzrhyqOT7UwM9Hfc0ZQ8Q`
    - Role: `technician`
- **tech2**
    - JWT: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Mn0.-ScBrpAXat0bA0Q-kJnL7xnst1-dd_SsIzseTUPT2wE`
    - Role: `technician`
- **boss3**
    - JWT: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6M30.l0XNQnn4xYlUafFxowInkYLvF3qvdwJ1iPcuf4Y_M90`
    - Role: `manager`

## Final Considerations

I chose not to implement additional features related to tasks (such as update and delete), nor the user-related functionalities.
This decision was made based on the limited time I had available to complete the challenge.

> I'm fully open to feedback regarding the solution and its implementation.