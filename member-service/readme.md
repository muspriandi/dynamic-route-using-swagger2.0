# Member Service API

This project is a simple Node.js API, built with **Express** and documented using **Swagger 2.0**.

## Features

- **Express** framework for building the API.
- **Swagger 2.0** for API documentation.
- CRUD operations for `member` resources.

## Project Structure

```
project/
├── index.js                        # Entry point of the application
├── routes/                         # API route handlers
│   ├── member.js                   # CRUD endpoints for members
│   ├── ping.js                     # Health check endpoint
├── lib/                            # Utility and configuration files
│   ├── swagger.js                  # Swagger setup
│   ├── constants.js                # Centralized constants
├── swagger/                        # Generated Swagger files
│   ├── member-swagger-spec.json    # Swagger specification JSON
├── package.json                    # Node.js project metadata
└── README.md                       # Project documentation
```

## Endpoints

#### Base URL: `http://localhost:8001`

| Method | Endpoint          | Description              |
|--------|--------------------|--------------------------|
| GET    | `/ping`           | Health check             |
| POST   | `/member`         | Create a new member      |
| GET    | `/member`         | Get all members          |
| PUT    | `/member/{id}`    | Update a member by ID    |
| DELETE | `/member/{id}`    | Delete a member by ID    |

### Example Requests

#### Create a Member
**POST** `/member`  
Request Body:
```json
{
  "name": "John Doe",
  "age": 30
}
```

Response:
```json
{
  "message": "Member created",
  "data": {
    "name": "John Doe",
    "age": 30
  }
}
```

#### Get All Members
**GET** `/member`  
Response:
```json
[
  { "id": "1", "name": "John Doe", "age": 30 },
  { "id": "2", "name": "Jane Smith", "age": 25 }
]
```

#### Update a Member
**PUT** `/member/1`  
Request Body:
```json
{
  "name": "John Updated",
  "age": 35
}
```

Response:
```json
{
  "message": "Member 1 updated",
  "data": {
    "name": "John Updated",
    "age": 35
  }
}
```

#### Delete a Member
**DELETE** `/member/1`  
Response:
```json
{
  "message": "Member 1 deleted"
}
```

## Installation and Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/muspriandi/dynamic-route-using-swagger2.0.git
   cd member-service
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
   Visit `http://localhost:8001/swagger` in your browser.