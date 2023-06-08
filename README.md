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

# CI/CD ->  docker Container

# How to run

## Build Docker Image
```shell
docker build -t mini-project .
```

## Run Docker Image
```shell
docker run contoh
```

## Docker Compose

Menjalankan semua service:

```shell
docker-compose up -d
```

Shutdown service:

```shell
docker-compose down
```


## Migrasi Data
### Export Database Menggunakan Mysqldump

1. Export database dengan `mysqldump`

```shell
mysqldump -u root -p exercise_sql > ./script/dump.sql
```

2. Copy file sql ke dalam container, misalnya di path `/home/namadatabase.sql`

```shell
docker compose cp ./script/dump.sql db:/home/dump.sql
# or with script.sql
docker compose cp ./script/script.sql db:/home/script.sql
```

3. Masuk ke shell DB:

```shell
docker compose exec -it db sh
```

4. Masuk ke sesi mysql:

```shell
mysql -u root -p
```

5. Jalankan perintah:

```shell
use exercise_sql;
source /home/dump.sql;
```

# mockery

```shell
go install github.com/vektra/mockery/v2@v2.20.0
```
generate mocks
```shell
mockery --all --keeptree --case underscore --with-expecter
```