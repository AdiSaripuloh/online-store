Online Store Challenge
====================

## Challenge

Because of these misreported inventory quantities, the Order Processing department was unable to fulfill a lot of
orders, and thus requested help from our Customer Service department to call our customers and notify them that we have
had to cancel their orders.

-------------------------------------------------------------------------

## Database Design

![Database Design](https://imgur.com/download/iDKw6tP/)

-------------------------------------------------------------------------

## API Design

**All responses in JSON and have the appropriate Content-Type header**

### GET /products

```
GET /products
Content-Type: "application/json"
```

##### Returns:

```
200 OK
Content-Type: "application/json"

[
    {
        "id": "f7630a17-2850-4c77-a7b6-99554814ccb8",
        "name": "Product 1",
        "price": 20000.00,
        "quantity": 100
    },
    ...
]
```

##### Errors

Error | Description
----- | ------------
500   | Internal server error

### POST /cart

```
POST /cart
Content-Type: "application/json"

[
    {
        "productId": "f7630a17-2850-4c77-a7b6-99554814ccb8",
        "quantity": 10,
    },
    ...
]
```

Attribute | Description
--------- | -----------
productId | Product ID
quantity  | Product quantity

##### Returns:

```
201 Created
Content-Type: "application/json"

{
    "message": "Product has been added to cart"
}
```

Status Code | Description
----------- | -----------
201         | Cart has been added

##### Errors

Error | Description
----- | ------------
422   | Validation error
500   | Internal server error

### GET /cart

```
GET /cart
Content-Type: "application/json"
```

##### Returns:

```
200 OK
Content-Type: "application/json"

{
    "id": "f7630a17-2850-4c77-a7b6-99554814ccb8",
    "products": [
        {
            "productId": "f7630a17-2850-4c77-a7b6-99554814ccb8",
            "quantity": 10,
        },
        ...
    ]
}
```

Status Code | Description
----------- | -----------
200         | Success get cart and products

##### Errors

Error | Description
----- | ------------
404   | Cart not found
500   | Internal server error

### PUT /cart

```
PUT /cart
Content-Type: "application/json"

[
    {
        "productId": "f7630a17-2850-4c77-a7b6-99554814ccb8",
        "quantity": 5,
    },
    ...
]
```

Attribute | Description
--------- | -----------
productId | Product ID
quantity  | Product quantity

##### Returns:

```
200 OK
Content-Type: "application/json"

{
    "message": "Cart has been updated"
}
```

Status Code | Description
----------- | -----------
200         | Cart has been updated

##### Errors

Error | Description
----- | ------------
404   | Cart or product not found
422   | Validation error
500   | Internal server error

-------------------------------------------------------------------------

## Installation

### Requirement
- Go version 1.15
- MySQL

### Steps
- Copy `.env.exanmple` to `.env`
- Create new database
- Setup `.env`
```dotenv
APP_MODE=debug
#APP_MODE=release
#APP_MODE=test
APP_PORT=8000

DB_DRIVER="mysql"
DB_HOST="127.0.0.1"
DB_PORT=3306
DB_USERNAME="root"
DB_PASSWORD="password"
DB_DATABASE="online_store"
```
- Run `go run cmd/http/httpMain.go`
- Open [http://localhost:8000](http://localhost:8000)

-------------------------------------------------------------------------

## Others

- Postman Collection [here](https://documenter.getpostman.com/view/3522941/Tz5jg1SG)
