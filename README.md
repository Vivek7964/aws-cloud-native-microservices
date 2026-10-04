# Watchn – Cloud-Native Microservices Platform

A production-style microservices platform deployed on Kubernetes and Amazon EKS, demonstrating containerization, infrastructure as code, GitOps, autoscaling, monitoring, and centralized logging.

---

## 🏗️ Architecture

![Watchn Architecture](docs/images/watchn-architecture.png)

### Microservices Architecture

![Microservices Architecture](docs/images/Microservices-architecture.png)

---

## 🎥 Project Demo

Watch the complete project demonstration:

<temp>

---

## 📌 Overview

This project modernizes the **Watchn microservices application** into a lightweight cloud-native platform.

The application is containerized and deployed on **Amazon EKS** using **Docker, Kubernetes, Helm, Terraform, and Argo CD**. The original database-heavy architecture was simplified by replacing unnecessary persistent databases with JSON and in-memory storage while retaining the core microservice architecture.

The project also implements:

- Kubernetes-based microservices deployment
- AWS EKS infrastructure
- Terraform Infrastructure as Code
- Helm and Helmfile
- Argo CD GitOps
- Horizontal Pod Autoscaling
- AWS Pod Identity
- Prometheus and Grafana monitoring
- Fluent Bit, Elasticsearch and Kibana centralized logging
- ActiveMQ asynchronous messaging
- AWS Application Load Balancer

---

## 💾 Storage Architecture

The application was simplified to remove unnecessary database infrastructure.

```text
Catalog
   │
   └── products.json

Carts
   │
   └── In-Memory

Orders
   │
   └── In-Memory

Checkout
   │
   └── In-Memory

Assets
   │
   └── Stateless

ActiveMQ
   │
   └── Asynchronous Messaging
```

### Catalog

Product information is stored in:

```text
src/catalog/products.json
```

The Catalog service reads product information from the JSON file instead of MySQL.

### Carts

Carts use an in-memory data structure.

### Orders

Orders use an in-memory repository.

### Checkout

Checkout uses an in-memory repository instead of Redis.

### ActiveMQ

ActiveMQ is retained as the messaging layer for asynchronous communication between services.

---

## 🛠️ Technology Stack

### Application

- Java
- Go
- Node.js
- Nginx
- ActiveMQ

### Containerization

- Docker
- Docker Compose

### Cloud

- AWS
- Amazon EKS
- Amazon VPC
- Application Load Balancer
- S3
- IAM
- AWS Pod Identity

### Kubernetes

- Kubernetes
- Helm
- Helmfile
- Gateway API
- Horizontal Pod Autoscaler

### Infrastructure

- Terraform

### GitOps

- Argo CD
- Argo CD Image Updater
- GitHub

### Monitoring

- Prometheus
- Grafana
- Alertmanager

### Logging

- Fluent Bit
- Elasticsearch
- Kibana
- ECK

---

## 📁 Project Structure

```text
microservice-demo/
│
├── .github/
│
├── deploy/
│   ├── argocd/
│   ├── docker-compose/
│   ├── kubernetes/
│   └── monitoring/
│
├── docs/
│   └── images/
│       ├── argocd-dashboard.png
│       ├── grafana-dashboard.png
│       ├── kibana-logs.png
│       ├── kubernetes-pods.png
│       ├── Microservices-architecture.png
│       ├── watchn-application.png
│       └── watchn-architecture.png
│
├── images/
│   ├── activemq/
│   ├── java17/
│   └── nodejs/
│
├── infrastructure/
│   └── terraform/
│       └── aws/
│
├── scripts/
│
├── src/
│   ├── assets/
│   ├── catalog/
│   ├── carts/
│   ├── checkout/
│   ├── orders/
│   └── ui/
│
├── .gitignore
├── LICENSE
└── README.md
```

---

## 🐳 Run Locally with Docker Compose

### Prerequisites

- Docker
- Docker Compose
- Git

### Clone the Repository

```bash
git clone https://github.com/Vivek7964/microservice-demo.git
cd microservice-demo
```

### Start the Application

```bash
cd deploy/docker-compose
docker compose up -d
```

### Check the Containers

```bash
docker compose ps
```

### View Logs

```bash
docker compose logs -f
```

### Stop the Application

```bash
docker compose down
```

---

## ☸️ Kubernetes Deployment

The application can be deployed to Kubernetes using Helm/Helmfile.

### Check the Kubernetes Context

```bash
kubectl config current-context
```

### Create the Watchn Namespace

```bash
kubectl create namespace watchn
```

### Deploy Using Helmfile

```bash
helmfile -e dev sync
```

### Check the Pods

```bash
kubectl get pods -n watchn
```

Expected services:

```text
activemq
assets
carts
catalog
checkout
orders
ui
```

---

## 🌐 Application Access

The application is exposed through an AWS Application Load Balancer.

```text
User
  │
  ▼
Application Load Balancer
  │
  ▼
Kubernetes Gateway
  │
  ▼
UI Service
  │
  ├── Catalog
  ├── Carts
  ├── Checkout
  ├── Orders
  └── Assets
```

The backend microservices remain internal to the Kubernetes cluster.

---

## 🏗️ Infrastructure with Terraform

AWS infrastructure is provisioned using Terraform.

```text
Terraform
    │
    ▼
AWS VPC
    │
    ├── Public Subnets
    │   ├── Internet Gateway
    │   ├── NAT Gateway
    │   ├── Application Load Balancer
    │   └── Bastion Host
    │
    └── Private Subnets
        │
        └── Amazon EKS
            │
            └── Node Group
```

Terraform provisions the core infrastructure required to run the application on EKS.

---

## 🔄 GitOps with Argo CD

Argo CD is used to implement GitOps-based Kubernetes deployment.

```text
Git Repository
      │
      ▼
   Argo CD
      │
      ▼
     EKS
      │
      ▼
Watchn Application
```

Argo CD continuously reconciles the Kubernetes cluster with the desired state stored in Git.

---

## 📦 CI/CD Architecture

The project architecture supports a CI/CD workflow:

```text
Developer
    │
    ▼
  GitHub
    │
    ▼
GitHub Actions
    │
    ├── Checkout
    ├── Build
    ├── Test
    ├── Trivy Scan
    └── Push Image
            │
            ▼
        Amazon ECR
            │
            ▼
   Argo CD Image Updater
            │
            ▼
      Git Manifests
            │
            ▼
         Argo CD
            │
            ▼
           EKS
```

The CI/CD architecture separates image creation from Kubernetes deployment.

---

## 📈 Monitoring

Prometheus collects application and Kubernetes metrics.

```text
Watchn Pods
     │
     │ /metrics
     ▼
 Prometheus
     │
     ▼
  Grafana
```

Grafana provides dashboards for application and Kubernetes observability.

Monitoring includes:

- Pod CPU usage
- Pod memory usage
- Pod restarts
- Running pod count
- Application availability
- Kubernetes node metrics

---

## 🔥 Alerting

Alertmanager handles alerts generated from the monitoring stack.

```text
Prometheus
    │
    ▼
Alertmanager
    │
    ▼
  Slack
```

Slack can be used for operational notifications.

---

## 📜 Centralized Logging

Application logs are collected and centralized using the Elastic Stack.

```text
Watchn Pods
     │
     ▼
 Fluent Bit
     │
     ▼
Elasticsearch
     │
     ▼
  Kibana
```

The logging stack consists of:

- **Fluent Bit** — log collection
- **Elasticsearch** — log storage and indexing
- **Kibana** — log visualization
- **ECK** — Kubernetes operator for Elasticsearch and Kibana

---

## 📊 Observability

The complete observability architecture is:

```text
                 Watchn Pods
                /           \
               /             \
          Metrics             Logs
             │                 │
             ▼                 ▼
        Prometheus         Fluent Bit
             │                 │
             ▼                 ▼
          Grafana        Elasticsearch
                               │
                               ▼
                             Kibana

        Prometheus
             │
             ▼
        Alertmanager
             │
             ▼
           Slack
```

---

## ⚡ Horizontal Pod Autoscaling

HPA automatically adjusts application replicas based on resource utilization.

```text
              Metrics
                 │
                 ▼
                HPA
             ┌───┴───┐
             ▼       ▼
         Scale Up  Scale Down
             │       │
             └───┬───┘
                 ▼
          Application Pods
```

This allows Kubernetes workloads to scale according to demand.

---

## 🔐 AWS Pod Identity

AWS Pod Identity provides AWS permissions to Kubernetes workloads without storing AWS access keys inside containers.

```text
Kubernetes Pod
      │
      ▼
Pod Identity Agent
      │
      ▼
AWS IAM Permissions
```

This provides a secure mechanism for Kubernetes workloads that need access to AWS resources.

---

## 📨 ActiveMQ

ActiveMQ provides asynchronous messaging between application components.

```text
Catalog
   │
   ▼
ActiveMQ
   ▲
   │
Orders
   ▲
   │
Carts
   ▲
   │
Checkout
```

The message broker allows services to communicate asynchronously without requiring direct database dependencies.

---

## 🔍 Kubernetes Observability

Kubernetes resources can be inspected using:

```bash
kubectl get pods -n watchn
```

```bash
kubectl get svc -n watchn
```

```bash
kubectl get deployments -n watchn
```

```bash
kubectl get hpa -n watchn
```

---

## 📸 Screenshots

### Application

![Watchn Application](docs/images/watchn-application.png)

### Argo CD

![Argo CD](docs/images/argocd-dashboard.png)

### Grafana

![Grafana](docs/images/grafana-dashboard.png)

### Kibana

![Kibana](docs/images/kibana-logs.png)

### Kubernetes

![Kubernetes](docs/images/kubernetes-pods.png)

---

## 🎯 Key Learning Outcomes

This project demonstrates practical experience with:

- Microservices architecture
- Docker
- Docker Compose
- Kubernetes
- Amazon EKS
- AWS VPC
- Terraform
- Helm
- Helmfile
- Kubernetes Gateway API
- Application Load Balancer
- ActiveMQ
- GitOps
- Argo CD
- Argo CD Image Updater
- Horizontal Pod Autoscaler
- AWS Pod Identity
- Prometheus
- Grafana
- Alertmanager
- Fluent Bit
- Elasticsearch
- Kibana
- ECK
- Centralized logging
- Kubernetes monitoring

---

## 🔮 Future Enhancements

Possible future improvements include:

- Automated CI/CD pipelines
- OpenTelemetry distributed tracing
- Service mesh
- Progressive delivery
- Canary deployments
- Advanced autoscaling
- Automated security scanning
- Cost optimization

---

## 👨‍💻 Author

**Vivek**

Cloud & DevOps Engineering Project
