# Dynamic Route API Gateway with Microservices

This repository contains an API Gateway implemented in Go and two microservices (`member-service` and `order-service`) built with Node.js. The API Gateway dynamically loads routes from Swagger 2.0 specifications, forwarding requests to the appropriate microservices.

---

## Project Structure

```
project-repo/
├── api-gateway/                  # Go-based API Gateway
│   ├── domain/                   # Domain models and configurations
│   ├── constants/                # Application-wide constants
│   ├── routes/                   # Routing logic
│   ├── handlers/                 # Request handlers
│   ├── swagger/                  # Swagger files for microservices
│   ├── main.go                   # Application entry point
│   ├── go.mod                    # Go module dependencies
│   └── README.md                 # Documentation for the API Gateway
│
├── member-service/               # Node.js Member Service
│   ├── index.js                  # Application entry point
│   ├── routes/                   # CRUD endpoints for members
│   ├── lib/                      # Utilities and Swagger setup
│   ├── swagger/                  # Swagger specification
│   ├── package.json              # Node.js dependencies
│   └── README.md                 # Documentation for the Member Service
│
├── order-service/                # Node.js Order Service
│   ├── index.js                  # Application entry point
│   ├── routes/                   # CRUD endpoints for orders
│   ├── lib/                      # Utilities and Swagger setup
│   ├── swagger/                  # Swagger specification
│   ├── package.json              # Node.js dependencies
│   └── README.md                 # Documentation for the Order Service
│
└── README.md                     # Root project documentation
```

---

## Features

1. **API Gateway**:
   - Dynamically loads Swagger 2.0 routes.
   - Proxies requests to appropriate microservices.
   - Supports easy addition of new services.

2. **Member Service**:
   - CRUD operations for managing members.
   - Fully documented using Swagger 2.0.

3. **Order Service**:
   - Basic operations for managing orders.
   - Fully documented using Swagger 2.0.

---

## Prerequisites

- **Go**: Version 1.23 (my current golang version).
- **Node.js**: Version 16 (my current node version).
- **Swagger 2.0 Specifications**: JSON files for service documentation (will generated in directory `/swagger` when running the node service (e.g. member-service)).

---

## Installation and Setup

### 1. Clone the Repository
```bash
git clone https://github.com/muspriandi/dynamic-route-using-swagger2.0.git
cd dynamic-route-using-swagger2.0
```

### 2. Set Up the API Gateway
1. Navigate to the `api-gateway` directory:
   ```bash
   cd api-gateway
   ```
2. Build and run the API Gateway:
   ```bash
   go run main.go
   ```

### 3. Set Up Microservices

#### Member Service
1. Navigate to the `member-service` directory:
   ```bash
   cd member-service
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Start the service:
   ```bash
   npm run start
   ```

#### Order Service
1. Navigate to the `order-service` directory:
   ```bash
   cd order-service
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Start the service:
   ```bash
   npm run start
   ```

---

## Usage

1. **Access the Gateway**:
   - Member Service: `http://localhost:8000/member-svc/<endpoint>`
   - Order Service: `http://localhost:8000/order-svc/<endpoint>`

2. **Test Endpoints**:
   Example using `curl`:
   ```bash
   curl http://localhost:8000/member-svc/ping
   curl http://localhost:8000/order-svc/ping
   ```

3. **View Swagger Documentation**:
   - Member Service: `http://localhost:8001/swagger`
   - Order Service: `http://localhost:8002/swagger`

---

## Adding a New Microservice

1. Create a new service with its Swagger 2.0 specification.
2. Add the Swagger file to the `api-gateway/swagger/` directory.
3. Update the API Gateway to include the new service:
   - Modify `constants/config.go` and `main.go`.
   - Define the service's route prefix and Swagger file path.
4. Restart the API Gateway:
   ```bash
   go run main.go
   ```
