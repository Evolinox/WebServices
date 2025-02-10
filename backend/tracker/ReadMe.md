# Tracker API V1.0

This is an API to consume different API's and store the data in a database. <br>

The API is running on Port 8082.
The API is consuming the product API on Port 8081.

These Endpoints are available:<br> <br>

GET /tracker/products <br>
GET /tracker/products/:id <br>
GET /tracker/products/name/:name <br>

Data set for products endpoints
```json
{
  "ID": "15348",
  "Name": "Skyr",
  "Brand": "K-Classic",
  "PortionSizeInGrams": 50,
  "CaloriesPer100Grams": 400,
  "ProteinsInGrams": 20.5,
  "FatsInGrams": 15.2,
  "CarbsInGrams": 45.0
}
```

GET /tracker/settings <br>
PATCH /tracker/settings <br>

Data set for settings endpoints
```json
{
  "PlannedCalories": 2500,
  "ProteinsPercentage": 33,
  "FatsPercentage": 20,
  "CarbsPercentage": 47
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
docker run -d --name tracker_db_mariadb -p 3406:3306 custom-mariadb
```

## Starting the API

To start the API, you can run the following command from the root directory of the project:

```shell script
go run api/cmd/main.go
```

## OpenApi Specification

For generating the OpenApi specification V3.0 we use the gin-swagger module.

https://github.com/swaggo/gin-swagger