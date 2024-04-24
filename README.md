# go-fiber-web-server

## Setup

### Initiate environment variables
copy `.env.sample` and rename it to `.env` and change the values based on your system.

### Install Prerequisite Dependencies (Run with Docker)

#### Install Docker
- Use Docker to install any dependencies such as Go and PostgreSQL which is used for current development. To install Docker, please follow the official [Get Docker](https://docs.docker.com/get-docker/) guideline.

### Install Prerequisite Dependencies (Run on Local Machine)

#### Install Golang
- Install Golang through this documentation [Golang installation](https://go.dev/doc/install). Version 1.17 or higher is required.

#### Install PostgreSQL
- Use PostgreSQL as databases in this project. See the installation process on [PosrgreSQL installation](https://www.postgresql.org/download/) based on your system. Setup your PostgreSQL with password as the password will be required.

#### Install golang-migrate
- Use golang-migrate for database schema migration tools. To install golang-migrate please follow the [golang-migrate installation](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md#installation).

#### Setup Database
- Follow this tutorial [setup database PostgreSQL with golang-migrate](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md) to setup database and its migration process.

## Run the Service

### With Docker
Prerequisites: Docker installed locally
```
docker-compose up
```

### On Local Machine
Prerequisites: Golang, PostgreSQL, and golang-migrate installed locally
```
# source environment variables
source .env

# run database migration (proceed after database has created)
migrate -database ${DB_URL} -path ${DB_MIGRATION_PATH} up

# install dependencies
go mod download

# run the service
go run main.go
```
