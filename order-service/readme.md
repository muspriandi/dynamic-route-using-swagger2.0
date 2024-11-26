# Order Service API

This project is a simple Node.js API, built with **Express** and documented using **Swagger 2.0**.

## Features

- **Express** framework for building the API.
- **Swagger 2.0** for API documentation.
- CRUD operations for `order` resources.

## Project Structure

```
project/
├── index.js                        # Entry point of the application
├── routes/                         # API route handlers
│   ├── ping.js                     # Health check endpoint
├── lib/                            # Utility and configuration files
│   ├── swagger.js                  # Swagger setup
│   ├── constants.js                # Centralized constants
├── swagger/                        # Generated Swagger files
│   ├── order-swagger-spec.json     # Swagger specification JSON
├── package.json                    # Node.js project metadata
└── README.md                       # Project documentation
```

## Endpoints

#### Base URL: `http://localhost:8002`

| Method | Endpoint          | Description              |
|--------|--------------------|--------------------------|
| GET    | `/ping`           | Health check             |


## Installation and Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/muspriandi/dynamic-route-using-swagger2.0.git
   cd order-service
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the server:
   ```bash
   npm run start        # node index.js
   npm run start:dev    # nodemon index.js
   ```

4. Open the Swagger UI for documentation:
   Visit `http://localhost:8002/swagger` in your browser.