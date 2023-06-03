# Mini Project 2
on this page:
- [dependencies](#dependencies)
- [Framework](#framework)
- [Architecture](#architecture-folder)
- [modules / internal description](#modules--internal-description)
- [Schema Database](#schema-database)
- [ERD Diagram](#erd-diagram)

# dependencies

using go Gin Gonic  
```
go get -u github.com/gin-gonic/gin
```

## Framework

- Framework : Go Gin
- Configuration : GoDotEnv

## Architecture

per-modules / internal:  
Delivery -> UseCase -> Repository

## modules / internal description:

- account : account admin, authentications
- customer : CRUD customer

## Schema Database:

The schema for the database is in path ./script/script.sql, and the previous repository [exercise_sql](https://github.com/tobialbertino/exercise_sql).
Using MySQL Database

## ERD Diagram

![](/script/ERD.png)