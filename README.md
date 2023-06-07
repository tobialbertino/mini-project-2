# Mini Project 2
on this page:
- [dependencies](#dependencies)
- [Framework](#framework)
- [Architecture](#architecture-folder)
- [modules / internal description](#modules--internal-description)
- [Schema Database](#schema-database)
- [ERD Diagram](#erd-diagram)
- [Notes](#notes)
- [Things that can be developed](#things-that-can-be-developed)

# dependencies

using Go Gin Gonic  
```
go get -u github.com/gin-gonic/gin
```

## Framework

- Framework: Go Gin
- Configuration: GoDotEnv
- Database: MySQL

## Architecture

per-modules / internal:  
Delivery -> UseCase -> Repository

## modules / internal description:

- account: account admin, authentications
- customer: CRUD customer

## Schema Database:

The schema for the database is in path ./script/script.sql, and the previous repository [exercise_sql](https://github.com/tobialbertino/exercise_sql).
Using MySQL Database

## ERD Diagram

![ERD Diagram](/script/ERD.png)

## Notes:

- GoRoutine for child process multi Query at the endpoint:
    - PUT /account/admin-reg
    - GET /account
    - GET /customer

Some Tx implementations could error. Error tx with go routine, temporary solution using *sql.DB object queries rather than Tx, and normal again.
Error message show below,

 ![error message](/script/Error-tx-select-rows.png)

## Things that can be developed

- separation of module/internal account with authentications (login)
- In tokenizing JWT, implement Refresh Token
- add more error handler
- Better Error_Handler Like the default/custom error handler from fiber/echo framework
- Implement Unit Test start from layer UseCase

Enjoy the code