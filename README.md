# Inventory Management System API

A RESTful backend API for managing product inventory, built with Go.
Supports product management, stock adjustments, and low stock monitoring.

## 🚀 Live Demo

**Base URL:** `https://inventory-system-production-fdfd.up.railway.app`

No setup needed. The API is live and connected to a real database.
You can test it directly using any HTTP client (Thunder Client, Postman)
or open GET endpoints directly in your browser.

**Quick test in browser:**

GET https://inventory-system-production-fdfd.up.railway.app/api/v1/products

GET https://inventory-system-production-fdfd.up.railway.app/api/v1/products/low-stock

## Tech Stack

- **Language:** Go
- **Framework:** Gin
- **Database:** PostgreSQL (Supabase)
- **ORM:** GORM
- **Deployment:** Railway

## Features

- Full CRUD for products
- Stock adjustment (stock in / stock out) with validation
- Low stock alert endpoint
- Pagination on product listing

## API Endpoints

### Products

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/products` | Create a new product |
| GET | `/api/v1/products` | List all products (paginated) |
| GET | `/api/v1/products/:id` | Get a single product |
| PUT | `/api/v1/products/:id` | Update a product |
| DELETE | `/api/v1/products/:id` | Delete a product |
| POST | `/api/v1/products/:id/adjust` | Adjust stock in or out |
| GET | `/api/v1/products/low-stock` | Get low stock products |

### Query Parameters

`GET /api/v1/products?page=1&limit=10`

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| page | int | 1 | Page number |
| limit | int | 10 | Results per page |

## Testing the API

Use Thunder Client (VS Code extension), Postman, or any HTTP client.
Set base URL to: https://inventory-system-production-fdfd.up.railway.app

### Example Requests

**Create a product**
```json
POST /api/v1/products
Content-Type: application/json

{
  "name": "Mechanical Keyboard",
  "sku": "KB-001",
  "category": "Electronics",
  "quantity": 50,
  "price": 299.90,
  "threshold": 10
}
```

**Adjust stock in**
```json
POST /api/v1/products/1/adjust
Content-Type: application/json

{
  "quantity": 20,
  "type": "in"
}
```

**Adjust stock out**
```json
POST /api/v1/products/1/adjust
Content-Type: application/json

{
  "quantity": 5,
  "type": "out"
}
```

**Update specific fields only**
```json
PUT /api/v1/products/1
Content-Type: application/json

{
  "quantity": 25
}
```

## Running Locally

### Prerequisites
- Go 1.21+
- PostgreSQL database (or Supabase account)

### Installation

1. Clone the repository
```bash
   git clone https://github.com/Nesa1000/inventory-system.git
   cd inventory-system
```

2. Create a `.env` file in the root directory
DATABASE_URL=your_postgresql_connection_string

3. Install dependencies
```bash
   go mod tidy
```

4. Run the server
```bash
   go run main.go
```

Server runs on `http://localhost:8080`

## Project Structure
```bash
inventory-system/
│
├── main.go           # Entry point
├── nixpacks.toml     # Railway deployment config
│
├── config/
│   └── db.go         # Database connection
│
├── models/
│   └── product.go    # Product model
│
├── handlers/
│   └── product.go    # Request handlers
│
├── routes/
│   └── routes.go     # Route definitions
│
├── .env              # Environment variables (not committed)
├── .gitignore
├── go.mod
└── README.md
```

## Future Improvements

- JWT authentication
- Stock movement history/audit log
- Category management endpoints
- Export inventory to CSV
- Docker support
- Frontend UI
