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

## Folder structure

- `frontend/` - Contains the Tauri App with Vue Frontend.
- `backend/` - Contains the Quarkus Backend.
