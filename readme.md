# WebServices

## Running Docker Container
If you want to run the application you need to have docker installed on your machine.
To start the docker container from the docker-compose.yaml run the following command:
```bash
make up
```

To stop the docker container run the following command:
```bash
make down
```

## Accessing the Frontend
Once the Docker Container is running (It might take up to 10 minutes) you can access the Frontend with your Browser. Just browse to the following URL:
```bash
localhost:1420
```

## Folder structure

- `frontend/` - Contains the Tauri App with Vue Frontend.
- `backend/` - Contains the Go Backend.
