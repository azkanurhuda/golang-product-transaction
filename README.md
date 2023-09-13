# Product Transaction
## Libraries
- [gorilla/mux](https://github.com/gorilla/mux)

## Getting Started
### Run Server
1. Run containers
    ```shell
    docker compose up -d
    ```
2. Stop containers
   ```shell
   docker compose stop
   ```

## Entity Relationship Diagram
```shell
https://drive.google.com/file/d/1cbqI1dXDQRdrSX4Qyd4FT0KsgLjYAlX5/view?usp=sharing
```

## API
### POST /signup
```shell
curl --location --request POST 'localhost:8888/signup' \
--form 'email=azka@email.com' \
--form 'password=azka123' \
--form 'username=azkapass'
```

### POST /login
```shell
curl --location --request POST 'localhost:8888/login' \
--form 'email=azka@email.com' \
--form 'password=azkapass'

```

### POST /v1/me
```shell
curl --location --request GET 'localhost:8888/v1/me' \
--header "Authorization: Bearer $JWT_TOKEN"
```

### GET List Product By Product Category 
#### /v1/products/:product_category_id
```shell
curl --location 'http://localhost:8888/v1/products/7376ebc0-c532-4e40-8541-8c3a782f83aa' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZWUxOWFjZjItM2NmMi00ZWZmLWE2MzEtMDMxNjRhYWVjNzRkIiwiZXhwIjoxNzA1MDI3OTcyfQ.z369zTlmoAKEweFdy0g2BmnYF4nmLhiq4Ikt7BAZ5hk'
```

### POST Add Product To Cart
#### /v1/carts
```shell
curl --location 'http://localhost:8888/v1/carts' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZDYxOWI1NGItOTM2My00YjY2LWE4ZGQtMDJhNGZlMzMyN2M1IiwiZXhwIjoxNzA1MDI5Mzc2fQ.ZhgefxX8emLv2pByBCyQsd9O_Y4gAzC5CQDrmfrPnPc' \
--form 'product_id="4a09d87b-9b2d-4dfe-8d65-da1f52bcbdf9"'
```

### GET Cart
#### /v1/carts
```shell
curl --location 'http://localhost:8888/v1/carts' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZDYxOWI1NGItOTM2My00YjY2LWE4ZGQtMDJhNGZlMzMyN2M1IiwiZXhwIjoxNzA1MDI5Mzc2fQ.ZhgefxX8emLv2pByBCyQsd9O_Y4gAzC5CQDrmfrPnPc'
```

### DELETE Cart
#### /v1/carts
```shell
curl --location --request DELETE 'http://localhost:8888/v1/carts/dc431e30-a2ed-4f31-a3d4-26bcd39c78b5' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZWUxOWFjZjItM2NmMi00ZWZmLWE2MzEtMDMxNjRhYWVjNzRkIiwiZXhwIjoxNzA1MDI3OTcyfQ.z369zTlmoAKEweFdy0g2BmnYF4nmLhiq4Ikt7BAZ5hk'
```

### POST Checkout
#### /v1/carts/checkout
```shell
curl --location 'http://localhost:8888/v1/carts/checkout' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjMxMTMxNzMtNzU2Zi00YTFjLTkwZGYtNDNlNDUwNTg3Nzk4IiwiZXhwIjoxNzA1MDMwMjIyfQ.zMkCL8CYROsntQc7t2hi1kNan8E2iB5z9ZxILYHb2HI' \
--form 'cart_id="9bc45763-ce08-4b62-bb36-8bf0beaae383"'
```

### POST Payment Transaction
#### /v1/payment/transaction
order_id get from id when checkout in route /v1/carts/checkout
```shell
curl --location 'http://localhost:8888/v1/payment/transaction' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNzUzMzJhMjYtMjU2OS00Yjc2LWFiNDMtNzRmMTA5OTVjNjkzIiwiZXhwIjoxNzA1MDMwODM2fQ.U8eaEyHTWTX8u8hRkLRybI9wCagYLT80cHNojDArePk' \
--form 'order_id="65728b61-b2ab-469e-aaf6-83d636e0f2ec"' \
--form 'payment_type="COD"'
```