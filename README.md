# Golang Todo List API

Requirements:
- Go version: 1.20.1

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
- On `/.env` set `APP_HOST=app-todo`, `DATABASE_HOST=mysql-todo` and `REDIS_HOST=redis-todo`
- Run `docker compose up -d --build`
- Open [http://localhost:3000](http://localhost:3000)
- Open [http://localhost:3000/api/v1/documentation](http://localhost:3000/api/v1/documentation) for swagger

## Note
- Use `swag init` command to generate swagger for documentation. [more details](https://github.com/swaggo/swag)

## References
- [Go Fiber Web Framework](https://docs.gofiber.io)
- Inspired by [Goshaka Starter](https://github.com/auliawiguna/goshaka-starter)
