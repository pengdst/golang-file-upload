
# Golang File Upload

Simple REST API with File Upload Support


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`PORT` #Optional, default 8080

`BASE_URL`

`DB_HOST`
`DB_PORT`
`DB_USER`
`DB_PASSWORD`
`DB_NAME`

`JWT_SECRET`


## Preparation

Install Soda CLI, skip if already Installed

```bash
go install github.com/gobuffalo/pop/v6/soda@latest
```
## Run Locally

Clone the project

```bash
git clone https://github.com/pengdst/golang-file-upload
```
or using SSH

```bash
git clone git@github.com:pengdst/golang-file-upload.git
```

Go to the project directory

```bash
cd golang-file-upload
```

Create environment

```bash
cp .env.example .env
```

Migrate Database

```bash
soda create -e development -c db/database.yml
soda migrate -e development -c db/database.yml
```

Or Migrate all Environment Database

```bash
soda create -a -c db/database.yml
soda migrate -a -c db/database.yml
```

Install dependencies

```bash
go mod download
go mod tidy
```

Run App

```bash
go run github.com/pengdst/golang-file-upload/cmd
```

Build Binary App

```bash
go build github.com/pengdst/golang-file-upload/cmd
```


## Tech Stack

**Database:** [Soda CLI](https://gobuffalo.io/documentation/database/soda/), [Gorm](https://gorm.io/docs/index.html#Quick-Start)

**Framework** [Gin](https://gin-gonic.com/docs/quickstart/)

**Environment** [GodotEnv](https://github.com/joho/godotenv), [caarlos0/env](https://github.com/caarlos0/env)


## Authors

- [@pengdst](https://www.github.com/pengdst)

