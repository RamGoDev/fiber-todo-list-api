# Todo List API

## About
Restful API that develops with Go Fiber Framework for CRUD Todo List with Repository pattern that supports caching and Elasticsearch. This project also implements the pipeline to make it easier for search improvement.

## Requirements:
- Go version: 1.20.1

## Features:
- Restful CRUD Todo List
- Middleware
- Validator
- Seeder
- Pipline
- Database (MySQL/PostgreSQL)
- Caching (Redis/Memcache)
- [WIP] Elasticsearch 
- Pipline
- Dockerize
- Swagger

# Setup
## Without Docker
- Copy `/.env.example` to `/.env` and set with your own credentials
- RUN `go get -d -v` to download the packages
- RUN `make build` to build executable file OR
- RUN `make run` to build and run the app
- Open [http://localhost:3000](http://localhost:3000)
- Open [http://localhost:3000/api/v1/documentation](http://localhost:3000/api/v1/documentation) for swagger

## With Docker (Compose)
- Copy `/.env.example` to `/.env` and set with your own credentials
- Copy `/docker/.env.example` to `/docker/.env` and set with your own credentials
- On `/.env` set:
    - `APP_HOST=app-todo`
    - `DATABASE_HOST=mysql-todo` and `DATABASE_PORT=3306` for `DATABASE_DRIVER=mysql`
    - `DATABASE_HOST=postgres-todo` and `DATABASE_PORT=5432` for `DATABASE_DRIVER=postgres`
    - `REDIS_HOST=redis-todo` for `CACHE_DRIVER=redis`
    - `MEMCACHE_HOST=memcache-todo` for `CACHE_DRIVER=memcache`
    - `ELASTICSEARCH_HOST="http://es01-todo"`
- Run `docker compose up -d --build`
- Open [http://localhost:3000](http://localhost:3000)
- Open [http://localhost:3000/api/v1/documentation](http://localhost:3000/api/v1/documentation) for swagger

## Note
- Use `swag init` command to generate swagger for documentation. [more details](https://github.com/swaggo/swag)

## References
- [Go Fiber Web Framework](https://docs.gofiber.io)
- Inspired by [Goshaka Starter](https://github.com/auliawiguna/goshaka-starter)

## Coming soon
- Fixing Elasticsearch integration
- Implement ACL (Access Control List) and manage users
- Add unit test
- Support GraphQL