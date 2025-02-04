# Product-API

## Docker Container

Before starting the API:
Run this command to start the Docker container

```shell script
docker run --name nutrition_products_db -e POSTGRES_USER=username -e POSTGRES_PASSWORD=password -e POSTGRES_DB=nutrition_products_db -p 5432:5432 postgres:17.2
```

## Running the application in dev mode

You can run your application in dev mode that enables live coding using:

```shell script
./mvnw quarkus:dev
```

## Testing Endpoints

API is available at localhost Port 8080