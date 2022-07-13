# go-jwt-authentication
---
Authenticating APIs with JWT tokens

---
######  Folder structure:
- src folder
    1. controllers
        : It mainly consists of functions that returns a handlerfunc for the routes defined in routes folder.
    2. database
        : It  creates a new Client from connection string provided in .env file and connects to the defined database and collection.
    3. helpers
        : This contains multiple functions broadly for token generation and validation.
    4. middleware 
        : The *GET request* routes are protected i.e. why :grin:  I did this project? This folder contains logic to fetch JWT token and validate the token from the request header.
    5. models
        : It contains the definition of user model which is used by the database as well as sent to the body of response.
    6. routes
        : It contains the logic to handle different requests to the backend server.
        
        - POST users/signup
        - POST users/login
        - GET /users
        - GET /users/:user_id
- main.go
    It's the entry point, starts the server, establish connection to db and after that server eagerly waits for the incoming requests.
- .env 
    This file contains the env variables such as PORT, JWT Secret and Connection URI
- go.mod
    It lists all the direct and indirect dependecies for this project.
---
###### What did I learn?
- Gin-Gonic It is a light weight web framework, mostly a wrapper around net/http package.
- context.Context One-liner explaination it communicates the cancellation notification.
Here we are using while interacting with the database as the request user can be cancelled before generating the response thus saving computational resources.
- JWT Token [Understanding through a video](https://www.youtube.com/watch?v=soGRyl9ztjIs)
Technical definition from [IBM](https://www.ibm.com/docs/en/order-management-sw/10.0?topic=features-jwt-authentication), this article also different types of JWT and its structure.
