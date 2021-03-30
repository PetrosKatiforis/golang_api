# Golang API With Authentication

This repository consists of a simple but feature rich REST API written in the go langauge. 

The project includes the following features:

- Json Web Tokens For Authentication
- A Basic Gorilla/Mux Router
- Authenticaiton Middleware
- Usage of Cookies
- A Postgres Database Using GORM
- Validation Using [Validator v10](https://github.com/go-playground/validator)
- Register / Login / Logout Routes
- Configuration With Environment Variables
- Routes For User Posts

## Getting Started

1) Setting up the environment variables
```python
PORT = "Define the local port that the server will use"

DATABASE_DIALECT = "Gorm's database dialect [e.g. 'postgres']"
DATABASE_CONNECTION_STRING = "Gorm's connection string to the database"

JWT_SECRET = "A secure key to sign your tokens"
```

2) Run the `` make dev `` command. The app will now run in *localhost:(env => PORT)* 

## API Routes

The project provides the following endpoints (Can be seen inside the utils/create_routes.go file):

- `` POST /users/new `` Register a new user
- `` POST /users/login `` Login from an account
- `` POST /users/logout `` Logout (Delete cookie)
- `` GET /psots/all `` Fetch the 15 latest posts
- `` GET /users/{id}/posts `` Fetch a user's 15 latest posts
- `` POST /posts/new `` Create a new post (Requires authentication cookie)
- `` DELETE /posts/{id}/delete `` Delete a specified post (Requires authentication cookie + The post's owner account)

