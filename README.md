# AuthService
Ready-to-use Authorization service 

## Requirements
- docker
- docker-compose

## Run
**Using docker:**
- ```docker-compose up -d``` or ```make compose-up```

**Locally:**
- set up postgres and create database url, should be smth like this ```postgres://user:pass@localhost:5432/postgres?sslmode=disable```
- export database url into ENV ```export PG_URL="postgres://user:pass@localhost:5432/postgres?sslmode=disable" ```
- in config/config.yml change the default settings to your own
- set up migrations with ```migrate -path migrations -database 'postgres://user:pass@localhost:5432/postgres?sslmode=disable' up``` 
- run ```go run cmd/app/main.go```

## API

**POST http://localhost:8080/user/register**
Register a new user

Example input:
```
{
    "username": "user",
    "email": "user@gmail.com",
    "password" : "password",
    "number" : "89223334433"
}
```
**GET http://localhost:8080/login**

Login an existing user and givin jwt token back

Example input:
```
{
    "username": "user",
    "password" : "password"
}
```

Example output:
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InZsYWRtc25rMSIsImVtYWlsIjoidnl1bW9pc2VlbmtvdjFAZ21haWwuY29tIiwiZXhwIjoxNjc4MzIwNDcxfQ.qVkxUFHK7rUajhEHvR81I8q1B_II_bkC92lpO0ulzbI"
}
```

**GET http://localhost:8080/api/v1/greet**

Bearer token required in Authorization headers

Example output for valid token:
```
{
    "greeting": "You have access to API!"
}
```
