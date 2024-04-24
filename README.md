# go-fiber-web-server

## Setup

### Initiate environment variables
copy `.env.sample` and rename it to `.env` and change the values based on your system.

### Install Prerequisite Dependencies

#### Install Golang
- Install Golang through this documentation [Golang installation](https://go.dev/doc/install). Version 1.17 or higher is required.

#### Install Fiber Web Framework
- Install Fiber v2 for stable release. Follow the installation process on [fiber installation](https://github.com/gofiber/fiber/tree/v2?tab=readme-ov-file#%EF%B8%8F-installation).

#### Install PostgreSQL
- Use PostgreSQL as databases in this project. See the installation process on [PosrgreSQL installation](https://www.postgresql.org/download/) based on your system. Setup your PostgreSQL with password as the password will be required.

#### Install GORM
- Install Golang ORM to communicate with databases to be developer friendly. To install it simply follow the process on [gorm installation](https://gorm.io/docs/#Install).

#### Install golang-migrate
- Use golang-migrate for database schema migration tools. To install golang-migrate please follow the [golang-migrate installation](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md#installation).

#### Setup Database
- Follow this tutorial [setup database PostgreSQL with golang-migrate](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md) to setup database and its migration process.

## Run the Service

Prerequisites: Golang, PostgreSQL, and golang-migrate installed locally
```
# source environment variables
source .env

# run database migration (proceed after database has created)
migrate -database ${DB_URL} -path ${DB_MIGRATION_PATH} up

# run the service
go run main.go
```
