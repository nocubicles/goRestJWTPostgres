# Boilerplate for Golang REST API with JWT authentication and Postgres

## To run this project:

1. Clone this repo
2. Run `docker-compose up -d` to start postgres or edit `db/database.go` file with your Postgres credentials
3. Run `main.go` and enjoy.

### Features:

- JWT based authentication
- Postgres database
- DB migrations
- Auth middleware
- Very minimal logging middleware
- REST API
- Models

### How to create user and authenticate and access private route

1. Run the project

2. POST request `localhost:3000/v1/user/ with JSON body`{ "first_name": "Good", "last_name": "Joe", "email": "user.name@gmail.com", "password": "kalasupp1" }`

3. POST request `localhost:3000/v1/login` with JSON body `{ "email": "user.name@gmail.com", "password": "kalasupp1" }`

4. GET request `localhost:3000/v1/secret` and include header `Authorization` with the value of token received in the previous step. You should get access to the private route.
