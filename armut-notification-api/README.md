![Go](	https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)

# Notification API

This project aims to provide recent ratings that has been added through rating-api. It only returns the ratings that are added since the previous API call. 

<hr/>

## Table of Contents:

- [Getting Started](#getting-started)
    - [Requirements](#requirements)
    - [Building and Running with Docker](#with-docker)
- [API Endpoints and Documentations](#api-endpoints-and-documentations)
    - [GET `/api/v1/notification`](#get-notifications)
- [Contact Information](#contact-information)
- [License](#license)
<hr/>


### Requirements:

- Go v1.20 or higher -> [Go Installation Page](https://go.dev/dl/)
- Docker 3.8 or higher -> [Docker Installation Page](https://docs.docker.com/engine/install/)
<hr/>

### Running the application

In the following steps, you can find information about running the application. After completing these steps, you can reach to the API from `http://0.0.0.0:8080/`.

#### With docker-compose
docker-compose image includes the PostgresQL database. It would be easier to run this project with docker-compose.

- First we need to build the docker image.

  ```
  docker-compose build
  ```

- To run the application, run the following command:

  ```
  docker-compose up -d
  ```

#### Running manually

- First we need to download the dependencies.

  ```
  go mod download
  ```

- Pulling the database

  ```
  docker pull postgres:14.1-alpine
  ```

- Running the database
  ```
  docker run --name postgres:14.1-alpine -e POSTGRES_PASSWORD=postgres -d postgres```
  ```
- Connect to the psql console and run the init.sql query in db console.

  You can find the init.sql in `scripts/db/init.sql`.

- Pulling swaggo/swag package and running

  ```
  # downloading swaggo
  go install github.com/swaggo/swag/cmd/swag@v1.8.12
  
  # initializing swaggo
  swag init
  ```

- Run the application

  ```
  go run .
  ```

<hr/>

### Running the tests

Before running the tests, you need to create the mocks. To create mock objects you need to run the following command:

```cd scripts/test && sudo sh mock-generate.sh```

You can delete the mocks with `mock-delete.sh` with the same commands.

To run the tests, you can execute the following command:

```go test -v```

<hr/>

## API Endpoints and Documentations

In the following part, you can find about api endpoints and their usages. You can also check the swagger documentation from http://0.0.0.0:8080/swagger

### GET `/api/v1/notification`

- Description: Return the recent rating from the last API call.


#### Request:

Query Params:
- `id`:
  - type: int
  - description: service provider id


#### Reponse:

```
{
    "Data": {
        "NotificationData": [
            4,
            2
        ]
    },
    "Message": "Success"
}
```

If the query param is not valid:

```
{
    "Data": null,
    "Message": "Query is not valid."
}
```

If there are no new ratings added:

```
{
    "Data": {
        "NotificationData": null
    },
    "Message": "Success"
}
```

| Status Code  | HTTP Meaning |                                                                   API Meaning |
| :------------ |:---------------:|------------------------------------------------------------------------------:|
| 200    | Success| Recent ratings added to the service provider id has been served successfully. |
| 400     | Bad Request       |                                       URL, JSON structure is not appropriate. |


<hr/>

#### Author: İlker Rişvan

#### Github: ilkerrisvan

#### Email: ilkerrisvan@outlook.com

#### Date: April, 2023

### License

<hr/>

[MIT](https://choosealicense.com/licenses/mit/)