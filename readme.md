# PetStore AcidLab
---
## About this project
The purpose of this project is to see the real performance of Go as a Web server. In order to achieve
this, we've made this little project to approach as much as possible to a "real" project
---
## Entities
### Pet
```
{
    id          integer($int64)
    category    Category
    name        string
    status      string, Enum: [ available, pending, sold ]
}
```
### User
```
User {
    id          integer($int64)
    username    string
    firstName   string
    lastName    string
    email       string
    password    string
    phone       string
    userStatus  string, Enum: [ active, disabled ]
}
```
### Category
```
Category {
    id          integer($int64)
    name        string
}
```
### Order
```
Order {
    id          integer($int64)
    petId       integer($int64)
    userId      integer($int64)
    quantity    integer($int32)
    shipDate    string($date-time)
    status      string, Enum: [ placed, approved, delivered ]
    complete    boolean, default: false
}
```
---
## Methods
### Pet
- POST /pets
- GET /pets
- GET /pets/{id}
- PATCH /pets/{id}
- DELETE /pets/{id}
### User
- POST /users
- GET /users
- GET /users/{id}
- PATCH /users/{id}
- DELETE /users/{id}
- GET /me
- PATCH /me
- POST /auth (login)
- DELETE /auth (logout)
### Order
- POST /orders
- GET /orders
- GET /orders/{id}
- PATCH /orders
- DELETE /orders
---
## Technologies used
This project runs with a minimal configuration of following technologies:
- **Go server** in the backend
- **PostgreSQL** as data persistor database
- **JSON Web Token (JWT)** as a authentication method
