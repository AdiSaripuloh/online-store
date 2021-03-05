Flash Sale Challenge
====================

## Challenge

Because of these misreported inventory quantities, the Order Processing department was unable to fulfill a lot of
orders, and thus requested help from our Customer Service department to call our customers and notify them that we have
had to cancel their orders.

-------------------------------------------------------------------------

## Database Design

![Database Design](https://imgur.com/download/JpfmHlJ/)

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
201         | Chart has been added

##### Errors

Error | Description
----- | ------------
422   | Validation error
500   | Internal server error