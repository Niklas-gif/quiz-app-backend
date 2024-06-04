# quiz-app-backend
 
### Install mongodb with Docker

```docker run --name quiz-app-db -p 27017:27017 -d mongo```

### Run Backend

```cd quiz-app-backend```\
```go run main.go``` 

### Create Admin for testing

Just visit "localhost:3030/admin" this will create a admin login for testing.\
Email    -> fake@mail.com\
Password -> Password1

### Run with docker

```cd docker```\
```docker-compose up --build```
you still have to create the Admin!