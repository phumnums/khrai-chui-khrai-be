# Khrai Chui Khrai — Backend

Backend API for **Khrai Chui Khrai**, a web application that helps people calculate how much each person should pay back when sharing purchases through the Thai government subsidy program.

Try the live application: [Khrai Chui Khrai](https://khrai-chui-khrai-fe.vercel.app/)

## Overview

The Thai government subsidy program covers 60% of eligible expenses, while the buyer pays the remaining 40%.

When multiple people buy items together, one person may pay the bill using the program. Since each person may purchase items of different values, splitting the payment equally may not be fair.

Khrai Chui Khrai calculates each person's share based on the value of the items they purchased.

## Tech Stack

- **Language:** Go
- **Framework:** Fiber v3
- **Middleware:** Fiber CORS
- **Deployment:** Render

## Project Structure

```text
backend/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── calculator/
│   │   └── calculator.go
│   ├── domain/
│   │   └── calculate.go
│   └── mapper/
│       └── calculate.go
├── go.mod
├── go.sum
└── README.md
```

- `cmd/api/` — Application entry point and HTTP route configuration.
- `internal/calculator/` — Business logic for calculating payment shares.
- `internal/domain/` — Request and response data models.
- `internal/mapper/` — Maps calculation results to response models.

## Getting Started

### Prerequisites

- Go installed on your machine.
- Git.

### Installation

Clone the repository and navigate to the backend directory.

```bash
git clone https://github.com/phumnums/khrai-chui-khrai-be.git
cd khrai-chui-khrai-be
```

Download dependencies.

```bash
go mod download
```

### Run Locally

Start the server:

```bash
go run ./cmd/api
```

The server runs on `http://localhost:3000` by default.

The application uses the `PORT` environment variable when provided.

## API Reference


### Calculate Payment Shares

Calculate how much each person should pay based on their purchased items.

**Endpoint**

```http
POST /api/v1/calculate
```

**Content-Type:** `application/json`

#### Request Body

```json
{
    "total_amount": 100,
    "government_support": 60,
    "paid_amount": 40,
    "order_items": [
        {
            "name": "Jubjib",
            "item_name": "Thai Tea",
            "purchased_amount": 35
        },
        {
            "name": "Jeje",
            "item_name": "Milk Green tea",
            "purchased_amount": 35
        },
        {
            "name": "Jame",
            "item_name": "Late",
            "purchased_amount": 30
        }
    ]
}
```

| Field | Type | Description |
|---|---|---|
| `total_amount` | integer | Total bill amount before the subsidy. |
| `government_support` | integer | Amount covered by the government. |
| `paid_amount` | integer | Amount paid by the person who scanned the payment. |
| `order_items` | array | List of purchased items. |
| `name` | string | Name of the person who purchased the item. |
| `item_name` | string | Name of the item. |
| `purchased_amount` | integer | Value of the purchased item. |

All monetary amounts must use the same unit. The current request model uses integer amounts.

#### Calculation Formula

Each item's share is calculated using the following formula:

```text
amount_to_pay = purchased_amount × paid_amount / total_amount
```

For example, if the total bill is 100 and the actual payment is 40:

- Alice purchases items worth 70 and pays 28.
- Bob purchases items worth 30 and pays 12.

The total amount to be paid back is 40.

#### Successful Response — `200 OK`

```json
{
    "data": [
        {
            "name": "Jubjib",
            "item_name": "Thai Tea",
            "purchased_total": 35,
            "amount_to_pay": "14.00"
        },
        {
            "name": "Jeje",
            "item_name": "Milk Green tea",
            "purchased_total": 35,
            "amount_to_pay": "14.00"
        },
        {
            "name": "Jame",
            "item_name": "Late",
            "purchased_total": 30,
            "amount_to_pay": "12.00"
        }
    ]
}
```

The response example illustrates the expected result. Field names and data types should match the current response model.

### Error Responses

**`400 Bad Request` — Invalid request body**

```json
{
  "error": "invalid request"
}
```

**`400 Bad Request` — Invalid bill amounts**

Returned when the bill total does not equal the government subsidy plus the amount paid by the buyer.

```text
total_amount = government_support + paid_amount
```

```json
{
  "error": "Invalid amount."
}
```

**`400 Bad Request` — Purchased total mismatch**

Returned when the sum of all purchased item amounts does not equal the total bill amount.

```text
sum(order_items[].purchased_amount) = total_amount
```

```json
{
  "error": "Total amount mismatch."
}
```

## Environment Variables

| Variable | Description | Default |
|---|---|---|
| `PORT` | Port used by the HTTP server. | `3000` |

## Deployment

The backend is deployed on Render.

**Calculate Endpoint**

```http
POST https://khrai-chui-khrai-be.onrender.com/api/v1/calculate
```

The backend allows requests from the configured frontend origin:

```text
https://khrai-chui-khrai-fe.vercel.app
```

## License

See the repository's LICENSE file if a license has been added.
