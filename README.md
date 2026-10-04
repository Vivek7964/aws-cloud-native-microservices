# Watchn - Simplified Microservices Demo

Watchn is a microservices-based application designed to demonstrate distributed, decoupled application components using multiple programming languages.

This version of Watchn keeps the **original service architecture and service boundaries** while simplifying the persistence layer for lightweight local development, Docker Compose, and Kubernetes/Minikube practice.

The original project is no longer maintained and points to the AWS Retail Store Sample App as its successor.

---

## Architecture

The application consists of multiple independent services communicating through REST APIs and asynchronous events.

The original microservice boundaries have been preserved. The main change in this version is the simplification of persistent storage.

![Architecture](/docs/images/architecture.png)

### Services

| Component | Language | Storage | Description |
|-----------|----------|---------|-------------|
| UI | Java | None | Aggregates API calls to the other services and renders the web UI. |
| Catalog | Go | JSON file | Provides the product catalog using a local `products.json` file. |
| Carts | Java | In-memory | Provides shopping cart functionality using ephemeral in-memory storage. |
| Orders | Java | In-memory | Provides order functionality using ephemeral in-memory storage. |
| Checkout | Node.js | In-memory | Orchestrates the checkout process using ephemeral in-memory storage. |
| Assets | Nginx | None | Serves static assets such as product images. |
| ActiveMQ | Java ecosystem | Message broker | Handles asynchronous order events. |

---

# Simplified Storage Architecture

The original application used several different persistence technologies such as MySQL, DynamoDB, MongoDB, and Redis.

This version intentionally removes those runtime storage dependencies while keeping the existing service boundaries.

The resulting storage model is:

```text
                    ┌─────────────┐
                    │     UI      │
                    │    Java     │
                    └──────┬──────┘
                           │
          ┌────────────────┼─────────────────┐
          │                │                 │
          ▼                ▼                 ▼
    ┌───────────┐    ┌───────────┐    ┌───────────┐
    │  Catalog  │    │   Carts   │    │  Orders   │
    │    Go     │    │   Java    │    │   Java    │
    │           │    │           │    │           │
    │ JSON file │    │ In-memory │    │ In-memory │
    └───────────┘    └───────────┘    └─────┬─────┘
                                            │
                                            │ OrderCreatedEvent
                                            ▼
                                      ┌─────────────┐
                                      │  ActiveMQ   │
                                      │   Broker    │
                                      └─────────────┘

                           ┌─────────────┐
                           │  Checkout   │
                           │   Node.js   │
                           │             │
                           │  In-memory  │
                           └─────────────┘

                           ┌─────────────┐
                           │   Assets    │
                           │    Nginx    │
                           └─────────────┘