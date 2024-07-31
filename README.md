# LookinLabs Test Task

The goal of this task is to evaluate your skills in software engineering, problem solving and coding.

This repository uses MVC layered architecture. Middleware gets the request and sends to controller. Controller sends the request to service. Service sends the request to repository. Repository sends the request to database. Database sends the response to repository. Repository sends the response to service. Service sends the response to controller. Controller sends the response to middleware. Middleware sends the response to client.

**Note** There are layers, that are not needed for this task like service.

## Task 1 (Software Engineering)

**1. Write an golang http server with basic requests**

Write an golang web http server based on Gin framework that will have 2 endpoints:

- GET /api/v1/headers that will return all headers from the request
- POST /api/v1/help that will return a help message
- POST /api/v1/users/create which will create a user in database via JSON Request Body
- GET /api/v1/users/get which will get all users from database in JSON Response Body

The messages should be done in a json way, e.g:
```json
{
  "message": "help message"
}
```

**2. Write an database connection and queries for `*/users*` endpoints**

Configure database connection and write queries, which will be used in `*/users*` endpoints.

First query should create a user in database. The user should have the following fields:
- id
- name
- email
- password
- created_at
- updated_at
- deleted_at

Second query should get all users from database.

**3. Write tests for the endpoints**

Write golang tests for the endpoints.

**4. Add simple frontend with simple form**

Add simple frontend with simple form in the middle of the page. On top from form you should have two tabs (Create User) and (Get User). When you click on the tab, the form should change to the form, which is needed for the endpoint.

Create User tab will make an POST request to /api/v1/users/create endpoint with header "Content-Type" = "application/json". The form should have the following request body:
- id
- name
- email
- created_at
- updated_at
- deleted_at

Example:

```json
{
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@no-reply.com",
    "password": "password",
    "created_at": "2021-10-10T10:00:00Z",
    "updated_at": "2021-10-10T10:00:00Z",
    "deleted_at": "2021-10-10T10:00:00Z"
}
```

Get User tab will make an GET request to /api/v1/users/get endpoint with header "Content-Type" = "application/json". The form should return all users with user ID's. For example they can be return as a table.

**5. Backend validation (advanced)**

Add validation for the request body in the endpoints. The validation should be done in model layer.

## Task 2 (Cloud Operations - advanced)

**1. Write a Dockerfile for the application**

Write a dockerfile for the application using docker best practices.

**2. Write a docker-compose file for the application**

Write a docker-compose file for the application using docker best practices. Docker compose should read environmentals from .env file, it should run golang and postgresql containers.

**3. Make the application to run in AWS App Runner (advanced)**

Run the application inside of AWS app runner via terraform. You can use [following repository](https://github.com/KostLinux/aws-app-runner-tf-template) to deploy the application.

But if you use the repository, remove everything, which is unnecessary to run the application in AWS App Runner (like route53).

## Task 3 (Problem Solving)

Solve the following problems:
- golang modules are not working, find and fix it
- problemSolver causes memory leak, find the issue and fix it

Hint: [htop](https://hisham.hm/htop/) can be used to monitor the memory leak.
Hint: Data slice length is the key to solve the problem.

## Comments

Leave your comments here