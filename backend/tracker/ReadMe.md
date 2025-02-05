# Tracker API V1.0

This is an API to consume different API's and store the data in a database. <br>

The API is running on Port 8082.
The API is consuming the product API on Port 8081.

These Endpoints are available:<br> <br>

GET /tracker/products <br>
GET /tracker/products/:id <br>
GET /tracker/products/name/:name <br>

## Starting the API

To start the API, you can run the following command from the root directory of the project:

```shell script
go run api/cmd/main.go
```

## OpenApi Specification

For generating the OpenApi specification V3.0 we use the gin-swagger module.

https://github.com/swaggo/gin-swagger