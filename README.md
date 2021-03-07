# Online Store Challenge

-------------------------------------------------------------------------

## Challenge

Because of these misreported inventory quantities, the Order Processing department was unable to fulfill a lot of
orders, and thus requested help from our Customer Service department to call our customers and notify them that we have
had to cancel their orders.

-------------------------------------------------------------------------

## Causes
- There is no transaction while insert to multiple tables or separated transactions to multiple transactions while inserting
  to multiple tables.
- When user doing checkout or payment application doesn't check product stock and after payment was successful product stock
is not updated.
- At the time of the flash sale there were too many users accessing **online store** and server was not strong enough
to bear too much traffic then made some transactions to the database unsuccessful.

## Solutions
- **Make** single database transaction when insert to multiple tables.
- **Do** stock validation when user checkout and payment.
- **Do** Update stock after payment was successful.
- **Do** rescale servers when at the time of the flash sale and consult to System Operations team.

-------------------------------------------------------------------------

## Database Design

![Database Design](https://imgur.com/download/iDKw6tP/)

-------------------------------------------------------------------------

## API Endpoint Design

**All responses in JSON and have the appropriate Content-Type header**

Postman Collection [here](https://documenter.getpostman.com/view/3522941/Tz5jg1SG)

### GET /user

```
GET /user
Content-Type: "application/json"
```

##### Returns:

```
200 OK
Content-Type: "application/json"

{
    "data": [
        {
            "id": "c39412d1-b547-42f6-9db5-ca27052841f9",
            "fullName": "Adi Saripuloh",
            "phone": "1234567890",
            "email": "adisaripuloh@gmail.com"
        },
        ...
    ],
    "status": "SUCCESS"
}
```

##### Errors

Error | Description
----- | ------------
500   | Internal server error

### GET /product

```
GET /product
Content-Type: "application/json"
```

##### Returns:

```
200 OK
Content-Type: "application/json"

{
    "data": [
        {
            "id": "00238ce0-e729-43f3-8c5b-2da93c9963cc",
            "name": "Product 1",
            "price": 10,
            "quantity": 4
        },
        ...
    ],
    "status": "SUCCESS"
}
```

##### Errors

Error | Description
----- | ------------
500   | Internal server error

### POST /cart

```
POST /cart
Content-Type: "application/json"
Authorization: "Bearer {{User ID}}"

{
    "items": [
        {
            "productID": "00238ce0-e729-43f3-8c5b-2da93c9963cc",
            "quantity": 2
        },
        {
            "productID": "0ccd34f6-97c9-4e25-911e-3e09536bdc29",
            "quantity": 2
        },
        ...
    ]
}
```

Attribute | Description
--------- | -----------
productID | Product ID
quantity  | Product quantity

##### Returns:

```
201 Created
Content-Type: "application/json"

{
    "data": {
        "id": "be4580d2-b148-49be-a707-ef84c7cd609b",
        "price": 40,
        "items": [
            {
                "id": "16429963-a925-49b2-a716-ea600db51097",
                "productID": "00238ce0-e729-43f3-8c5b-2da93c9963cc",
                "quantity": 2
            },
            {
                "id": "d5eedd90-d2a8-46d7-8f10-5d8aaecf7f64",
                "productID": "0ccd34f6-97c9-4e25-911e-3e09536bdc29",
                "quantity": 2
            },
            ...
        ]
    },
    "status": "SUCCESS"
}
```

##### Errors

Error | Description
----- | ------------
400   | Bad Request
422   | Validation error
500   | Internal server error

### GET /cart

```
GET /cart
Content-Type: "application/json"
Authorization: "Bearer {{User ID}}"
```

##### Returns:

```
200 OK
Content-Type: "application/json"

{
    "data": {
        "id": "be4580d2-b148-49be-a707-ef84c7cd609b",
        "price": 40,
        "items": [
            {
                "id": "16429963-a925-49b2-a716-ea600db51097",
                "productID": "00238ce0-e729-43f3-8c5b-2da93c9963cc",
                "quantity": 2
            },
            {
                "id": "d5eedd90-d2a8-46d7-8f10-5d8aaecf7f64",
                "productID": "0ccd34f6-97c9-4e25-911e-3e09536bdc29",
                "quantity": 2
            },
            ...
        ]
    },
    "status": "SUCCESS"
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

### POST /cart/checkout

```
POST /cart/checkout
Content-Type: "application/json"
Authorization: "Bearer {{User ID}}"

{
    "items": [
        {
            "productID": "00238ce0-e729-43f3-8c5b-2da93c9963cc",
            "quantity": 2
        },
        ...
    ]
}
```

Attribute | Description
--------- | -----------
productID | Product ID
quantity  | Product quantity

##### Returns:

```
200 OK
Content-Type: "application/json"

{
    "data": {
        "id": "28c8ca87-0c67-4a68-b439-fe02ad5ef854",
        "grandTotal": 30,
        "status": "UNPAID",
        "items": [
            {
                "id": "cba5155e-7014-406f-a1e8-0f4f992ac4ea",
                "productID": "00238ce0-e729-43f3-8c5b-2da93c9963cc",
                "quantity": 3
            },
            ...
        ]
    },
    "status": "SUCCESS"
}
```

Status Code | Description
----------- | -----------
201         | Cart has been added

##### Errors

Error | Description
----- | ------------
400   | Bad Request
422   | Validation error
500   | Internal server error

### GET /order

```
GET /order
Content-Type: "application/json"
Authorization: "Bearer {{User ID}}"
```

##### Returns:

```
200 OK
Content-Type: "application/json"

{
    "data": [
        {
            "id": "28c8ca87-0c67-4a68-b439-fe02ad5ef854",
            "grandTotal": 30,
            "status": "UNPAID",
            "items": [
                {
                    "id": "cba5155e-7014-406f-a1e8-0f4f992ac4ea",
                    "productID": "00238ce0-e729-43f3-8c5b-2da93c9963cc",
                    "quantity": 3
                }
            ]
        },
        ...
    ],
    "status": "SUCCESS"
}
```

Status Code | Description
----------- | -----------
200         | Success get orders

##### Errors

Error | Description
----- | ------------
400   | Bad request
500   | Internal server error

### GET /order/{{orderID}}

```
GET /order/{{orderID}}
Content-Type: "application/json"
Authorization: "Bearer {{User ID}}"
```

##### Returns:

```
200 OK
Content-Type: "application/json"

{
    "data": {
        "id": "28c8ca87-0c67-4a68-b439-fe02ad5ef854",
        "grandTotal": 30,
        "status": "UNPAID",
        "items": [
            {
                "id": "cba5155e-7014-406f-a1e8-0f4f992ac4ea",
                "productID": "00238ce0-e729-43f3-8c5b-2da93c9963cc",
                "quantity": 3
            }
        ]
    }
    "status": "SUCCESS"
}
```

Status Code | Description
----------- | -----------
200         | Success get order

##### Errors

Error | Description
----- | ------------
400   | Bad request
404   | Not found
500   | Internal server error

### POST /order/{{orderID}}/pay

```
POST /order/{{orderID}}/pay
Content-Type: "application/json"
Authorization: "Bearer {{User ID}}"

{
    "amount": 30
}
```

##### Returns:

```
200 OK
Content-Type: "application/json"

{
    "data": {
        "id": "28c8ca87-0c67-4a68-b439-fe02ad5ef854",
        "grandTotal": 30,
        "status": "PAID",
        "items": [
            {
                "id": "cba5155e-7014-406f-a1e8-0f4f992ac4ea",
                "productID": "00238ce0-e729-43f3-8c5b-2da93c9963cc",
                "quantity": 3
            }
        ]
    }
    "status": "SUCCESS"
}
```

Status Code | Description
----------- | -----------
200         | Success pay order

##### Errors

Error | Description
----- | ------------
400   | Bad request
422   | Unprocessable entity
404   | Not found
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
