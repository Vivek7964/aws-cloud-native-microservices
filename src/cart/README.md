# Watchn - Cart Service

This API service provides shopping carts for users with the following characteristics:

- Spring Boot application framework
- Persistence using in-memory storage
- Cart data is ephemeral and is lost when the service restarts
- Ability to introduce turbulent conditions with embedded "Chaos Monkey"

## Running

### Java

The service can be run locally using the following command:

```bash
./mvnw spring-boot:run