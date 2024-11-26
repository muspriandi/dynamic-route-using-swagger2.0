# API Gateway

This project is an **API Gateway** implemented in Go. It dynamically loads routes from Swagger 2.0 (`swagger-spec.json`) files for multiple services and forwards requests to the respective microservices with prefixed paths.

## Features

- Dynamically loads routes from `swagger-spec.json` files.
- Supports multiple services, each with its own route prefix.
- Proxies requests to the correct service, preserving headers and payloads.
- Easy-to-extend architecture for adding new services.

## Project Structure
project/
├── domain/                        
│   ├── config.go                  # Application-specific domain struct and configurations
├── constants/                     
│   ├── config.go                  # Definitions for constants used across the application
├── routes/                        
│   ├── routes.go                  # Routing definitions for API endpoints
├── handlers/                      
│   ├── handler.go                 # Handlers for processing incoming requests and responses
├── swagger/                       
│   ├── member-swagger-spec.json   # Swagger 2.0 specification for the Member service
│   ├── order-swagger-spec.json    # Swagger 2.0 specification for the Order service
├── main.go                        # Entry point of the application
├── go.mod                         # Go module file, manages dependencies
└── README.md                      # Documentation for the project


## How It Works

1. **Service Discovery**: 
   - The gateway reads `swagger-spec.json`(e.g. `member-swagger-spec.json`) files to extract routes for each microservice.
   - Each service is associated with a route prefix (e.g., `/member`, `/order`).

2. **Routing**:
   - The API Gateway matches incoming requests to the routes defined in the `swagger-spec.json` files.
   - Requests are forwarded to the appropriate service with headers and payloads preserved.

3. **Dynamic Service Addition**:
   - Add a new service by providing its `swagger-spec.json` file and configuration details.

## Installation

### Prerequisites

- Go 1.23 (my current golang version)
- Microservices providing Swagger 2/0 specifications (`swagger-spec.json` files).

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/muspriandi/dynamic-route-using-swagger2.0.git
   cd api-gateway
   ```

2. Place the `swagger-spec.json` files for your services in the project directory:
   - Example:
     - `member-swagger-spec.json` for the Member service.
     - `order-swagger-spec.json` for the Order service.

3. Build and run the application:
   ```bash
   go run main.go
   ```
   
## Usage

1. **Start the Microservices**:
   - Ensure your microservices (e.g., Member and Order services) are running.
   - Example:
     - Member Service: `http://localhost:8001`
     - Order Service: `http://localhost:8002`

2. **Access the Gateway**:
   - Member Service via Gateway: `http://localhost:8000/member-svc/<endpoint>`
   - Order Service via Gateway: `http://localhost:8000/order-svc/<endpoint>`

3. **Test the Gateway**:
   ```bash
   curl http://localhost:8000/member/ping
   curl http://localhost:8000/order/ping
   ```

## Adding a New Service

### Setup When Adding a New Service

1. Update the `constants/config.go` file and add new service config:
   ```go
	ORDER_SERVICE_CODE   = "order"
	ORDER_SERVICE_PREFIX = "/order-svc"
   ```
   
2. Update the `main.go` file and define swagger spec for new service:
   ```go
	orderSpecPath := filepath.Join(swaggerDir, fmt.Sprintf("%s-swagger-spec.json", constants.ORDER_SERVICE_CODE))
	orderPathPrefix := constants.ORDER_SERVICE_PREFIX
   ```
   
    ```go
	serviceConfigs[constants.ORDER_SERVICE_CODE] = routes.LoadSwaggerSpec(
		constants.ORDER_SERVICE_CODE,
		orderSpecPath,
		orderPathPrefix,
	)
   ```
3. Build and run the application:
   ```bash
   go run main.go
   ```