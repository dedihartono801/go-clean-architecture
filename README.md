
## Description

[Go clean go-clean-architecture]

Examples of types of communication;
- API
- CLI

Examples of data persistence;
- Mysql
- Mongo
- In memory

Example:
- rest api
- Dockerfile (for deployment prod/staging)
- DockerfileDev (for local with hot reload)
- Docker-compose (to run app)
- swagger for doc
- middleware auth jwt
- migration
- unit testing

## Run Service

```bash
$ docker-compose up -d
```

## Run Migration UP

```bash
$ make migration-up
```

## Run Migration Down

```bash
$ make migration-down
```

## Create Migration

```bash
$ make migration
#type your migration name example: create_create_table_users
```

## Generate JWT Secret

```bash
$ make generate-jwt-secret
#copy the secret key and then create new env called JWT_SECRET in .env file:
```

## Generate Swagger

```bash
$ make generate-swag
```

## Run Test Coverage

```bash
$ make test-cov
```

## API Documentation
API Documentation was created with swagger and is available at `http://localhost:5001/docs`
