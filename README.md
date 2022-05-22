# Plant API
Plant API is a REST API built with Echo server. The code implementation uses Hexagonal Architecture. Below you will find a current list of the available methods on our plant, native, and user admin API.

## Requirements
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## How to Run Server
- Adjust your environment configuration in `docker-compose.yml`
- Run Docker Compose
```bash
docker-compose up -d
```

## How to Consume the API
First, you have to login using super account to have the full access in API.
```js
{
  "email": "super@company.com",
  "password": "super123"
}
```
Here is some of API that you can use:
- POST /v1/login
- POST /v1/plants
- GET /v1/plants
- etc.

See our [documentation](https://app.swaggerhub.com/apis-docs/dhichii/plant-api/v1) for more details.

## Tech Stack
- Language: **Go 1.18**
- Rest Server: **Echo v4**
- Database: **MySQL**
- ORM: **GORM**
- Deployment: **Docker**, **Amazon EC2**
- CI/CD: **Github Action**
- Authorization: **JWT**
- Unit Test: **Mockery (mocking)**, **Testify (testing toolkit)**
