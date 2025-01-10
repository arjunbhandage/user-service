# **User Service**

User Service is a RESTful API built with Go and PostgreSQL that handles user management functionality, including user registration, login, and listing users.

---

## **Features**
- User sign-up with encrypted password storage.
- User sign-in with email and password authentication.
- API to list all users.
- Secure communication with hashed passwords using `bcrypt`.

## Installation
Run the following command to download and install all required dependencies:
```
go mod tidy
```
Run the server locally:
```
go run main.go
```

## API Endpoints

1. User Sign Up

Endpoint: POST user/signup
Registers a new user.

Request:
```
{
  "email": "user@example.com",
  "password": "password123",
  "first_name": "John",
  "last_name": "Doe"
}
```
Response:
```
{
  "message": "User created successfully"
}
```

2. User Sign In

Endpoint: POST user/signin
Authenticates an existing user.

Request:
```
{
  "email": "user@example.com",
  "password": "password123"
}
```

3. List Users

Endpoint: GET /users
Returns a list of all registered users.

```
[
  {
    "id": 1,
    "email": "user@example.com",
    "first_name": "John",
    "last_name": "Doe"
  },
  {
    "id": 2,
    "email": "another@example.com",
    "first_name": "Jane",
    "last_name": "Smith"
  }
]
```

