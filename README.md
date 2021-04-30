# About This Project
This project is example of REST API project build with Go language (especially gin-gonic framework).

# Implementation
This Project implement:
1. Gin gonic framework
2. Repository-Service Pattern.
3. unit testing using testify.
4. Basic auth for authentication
5. Validation for validate user input

# Project Structure
1. Config (this folder/package used as configuration such as server, database, etc)
2. Entity (this folder/package used to describe data type such as payload or database entity)
3. Service (this folder/package contain all business logic)
4. Repository (this folder/package is used to store data or communicate to database)

# Project dependencies
```go
    github.com/gin-gonic/gin 
	github.com/go-playground/validator/v10 
	github.com/stretchr/testify
	golang.org/x/crypto 
```

# Setup
1. Clone / Download this project
2. go run .

# Testing
```go
cd services/user
go test -v
```

# Usage
### Create User (Admin only)
 - URL
     - /user
 - Method
     - POST
 - Auth
     use basic auth:
     - username
     - password
 - Body / Parameter
     - ```rest
         {
            "fullname": "farhan ramadhana",
        	"email": "email@gmail.com",
        	"password": "pass1234567890", // minimal 10 characters
        	"role": "user" // role: user/admin
        }
     ```
 - Success Response
     - ```rest
         {
            "message": true,
            "status": "ok"
         }
     ```
 - Error Response s
     - ```rest
        {
            "message": "",
            "status": "error"
        }
     ```
### Get All User
### Get User By ID
### Get User By Email
### Get User By UserName
### Delete User By ID (Admin only)
### Update User By ID (Admin only)

