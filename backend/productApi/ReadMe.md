# Product API V1.0

This is an API for managing products. It allows you to create, read, update and delete products.
The API is written in Go and uses the Gin framework.

The API provides data in JSON format. <br>
The following endpoints are available:<br> <br>
GET /products <br>
GET /products/:id <br>
GET /products/name/:name <br>
POST /products <br>
DELETE /products/:id <br>

The data retrieved from the database is in the following format:  <br>
for example:
```json
{
  "ID": 15348,
  "Name": "Skyr",
  "Brand": "K-Classic",
  "CaloriesPer100Grams": 400,
  "ProteinsInGrams": 20.5,
  "FatsInGrams": 15.2,
  "CarbsInGrams": 45.0
}
```

## Starting the Docker container

To run this Product API you need to have a MariaDB database running on Port 3306. 
You can use the following Docker commands to start a MariaDB container.

### Build the Docker image
```shell script
docker build -t custom-mariadb .
```

### Start the Docker container
```shell script
docker run -d --name product_db_mariadb -p 3405:3306 custom-mariadb
```

## Starting the API

To start the API, you can run the following command from the root directory of the project:

```shell script
go run api/cmd/main.go
```

## OpenApi Specification

For generating the OpenApi specification V3.0 we use the gin-swagger module.

https://github.com/swaggo/gin-swagger