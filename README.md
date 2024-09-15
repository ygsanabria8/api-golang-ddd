
# Api Golang DDD

API Golang DDD is a RESTful API designed as a base template for the quick creation of robust and scalable APIs, following the principles of Domain-Driven Design (DDD). This API offers native integration with multiple SQL and NoSQL databases, as well as support for messaging systems like Kafka.

Among its key features are:

- Dependency injection for modular and flexible component management.
- Secure secret management for the configuration of sensitive variables.
- Implementation of the Mediator pattern, facilitating the adoption of the CQRS pattern to separate query and command responsibilities for data access.
- Container support with a Docker image, allowing for consistent deployments and portability across different environments.
- Helm package available to simplify the application's lifecycle management in Kubernetes, including the ability to perform deployments, updates, and rollbacks efficiently.

## Technologies

- [Gorm](https://gorm.io/index.html)
- [Viper](https://github.com/spf13/viper)
- [Fx](https://uber-go.github.io/fx/index.html)
- [Zap](https://pkg.go.dev/go.uber.org/zap#section-readme)
- [Protocol Buffer](https://protobuf.dev/getting-started/gotutorial/)
- [Mockery](https://vektra.github.io/mockery/latest/)
- [Gin Gonic](https://gin-gonic.com/docs/)

## Architecture
This project follows the Domain-Driven Design (DDD) pattern, dividing the project into multiple bounded contexts.

- **Domain Layer**: Contains core business logic, entities and domain services.
- **Application Layer**: Implements the use cases of the application using CQRS and Mediator.
- **Infrastructure Layer**: Contains MySQL and MongoDB repositories, and other external system integrations.
- **Api Layer**: Handles incoming HTTP requests using the Gin framework and the Kafka event handlers.

## Design Patterns

#### Mediator Pattern
The Mediator pattern is used in the Api Layer to handle communication between different components without them being aware of each other. This is particularly useful when orchestrating complex workflows in a decoupled way.

#### CQRS (Command Query Responsibility Segregation)
- **Command**: Handles requests that modify the state of the application (e.g., Create, Update).
- **Query**: Handles requests that retrieve data without modifying the state.
- CQRS is applied to separate read and write operations in the system, improving performance and scalability.

## Prerequisites

- [Golang 1.22+](https://go.dev/dl/)
- [Protoc](https://grpc.io/docs/protoc-installation/)
- [Air](https://github.com/air-verse/air)
- [Mockery](https://vektra.github.io/mockery/latest/installation/)
- [Docker](https://docs.docker.com/engine/install/)
- Make
    ```bash
      sudo apt-get install make
   ```
## Run Locally
To execute this project run:

##### Up Docker Compose
   ```bash
      docker compose up -d
   ```
##### Set up environment variables using Viper and Docker:

Create a .env file for local environment variables accordint to env.template.
Use Docker for spinning up Kafka, MySQL, and MongoDB containers.

##### Install dependencies
   ```bash
      make install
   ```
##### Start project
   ```bash
      make run
   ```

## Execute Unit Tests
To execute unit test run:

##### Create Mocks
      make build_mocks
##### Run Unit Tests
      make tests

## Contributing
Contributions are always welcome!
If you want to contribute to the project, feel free to submit a pull request or create an issue.

## License
[Apache License 2.0](https://github.com/ygsanabria8/api-golang-ddd/blob/main/LICENSE)


## Documentation
[Documentation](https://github.com/ygsanabria8/api-golang-ddd/blob/main/doc/api-golang-ddd.json)

