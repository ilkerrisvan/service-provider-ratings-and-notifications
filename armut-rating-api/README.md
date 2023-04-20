![Go](	https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)

# Rating API

This project aims to store rates and provide rating averages of service providers.

<hr/>

## Table of Contents:

- [Getting Started](#getting-started)
    - [Requirements](#requirements)
    - [Building and Running with Docker](#with-docker)
- [API Endpoints and Documentations](#api-endpoints-and-documentations)
    - [POST `/api/v1/rating`](#post-ratings)
    - [GET `/api/v1/rating/average`](#get-average)
- [Contact Information](#contact-information)
- [License](#license)
<hr/>


### Requirements:

- Go v1.20 or higher -> [Go Installation Page](https://go.dev/dl/)
- Docker 3.8 or higher -> [Docker Installation Page](https://docs.docker.com/engine/install/)
<hr/>

### Running the application

In the following steps, you can find information about running the application. After completing these steps, you can reach to the API from `http://0.0.0.0:8000/`.

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

In the following part, you can find about api endpoints and their usages. You can also check the swagger documentation from http://0.0.0.0:8000/swagger

### POST `/api/v1/rating`

- submitting a new service provider rating for a service provider id. 

#### Validations
- ServiceProviderRating must be between 0 and 5.
- ServiceProviderId must be greater than or equal to 0.

#### Request:

Body:

- `ServiceProviderId`:
  - description: id of the service provider 
  - type: int
  - required: true
  - validations: must be greater than or equal to 0.
- `ServiceProviderRating`:
  - description: rating of the service provider
  - type: int
  - required: true
  - validations: must be between 0 and 5.

<br/>

Example request body:

``` 
{
    "ServiceProviderId": 123,
    "ServiceProviderRating": 5
}
```

#### Reponse:

201 Created:
```
{
    "Data": {
        "ServiceProviderId": 123,
        "ServiceProviderRating": 5
    },
    "Message": "Success"
}
```

400 Bad Request:

```
{
    "Data": null,
    "Message": "Request is not valid. Check the body."
}
```
| Status Code  | HTTP Meaning |                                  API Meaning |
| :------------ |:---------------:|---------------------------------------------:|
| 201    | Created| The new rating was carried out successfully. |
| 400     | Bad Request       |      URL, JSON structure is not appropriate. |

<hr/>

### GET `/api/v1/rating/average`

- Description: Return the average rate of the id.


#### Request:

Query Params:
- `id`:
  - type: int
  - description: service provider id


#### Reponse:

```
{
    "Data": {
        "AverageRating": 3.2727273,
        "TotalRatingCount": 11
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

If there is no data added to the system:

```
{
    "Data": null,
    "Message": "This id has no rate information."
}
```

| Status Code  | HTTP Meaning |                                                             API Meaning |
| :------------ |:---------------:|------------------------------------------------------------------------:|
| 200    | Success| Average rating of the service provider id has been served successfully. |
| 400     | Bad Request       |                                 URL, JSON structure is not appropriate. |


<hr/>

#### Author: İlker Rişvan

#### Github: ilkerrisvan

#### Email: ilkerrisvan@outlook.com

#### Date: April, 2023

### License

<hr/>

[MIT](https://choosealicense.com/licenses/mit/)