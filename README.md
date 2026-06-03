# Inventory Management System API

A RESTful backend API for managing product inventory, built with Go.
Supports product management, stock adjustments, and low stock monitoring.

## Tech Stack

- **Language:** Go
- **Framework:** Gin
- **Database:** PostgreSQL (Supabase)
- **ORM:** GORM

## Features

- Full CRUD for products
- Stock adjustment (stock in / stock out) with validation
- Low stock alert endpoint
- Pagination on product listing

## Getting Started

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

## Example Requests

### Create Product
```json
POST /api/v1/products
{
  "name": "Mechanical Keyboard",
  "sku": "KB-001",
  "category": "Electronics",
  "quantity": 50,
  "price": 299.90,
  "threshold": 10
}
```

### Adjust Stock
```json
POST /api/v1/products/1/adjust
{
  "quantity": 20,
  "type": "in"
}
```

```json
POST /api/v1/products/1/adjust
{
  "quantity": 5,
  "type": "out"
}
```

## Project Structure
inventory-system/
│
├── main.go           # Entry point
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

## Future Improvements

- JWT authentication
- Stock movement history/audit log
- Category management endpoints
- Export to CSV
- Docker support
- Frontend UI
