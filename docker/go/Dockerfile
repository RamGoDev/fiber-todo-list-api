# Step 1: Build Stage
FROM golang:1.20-alpine3.17 AS BuildStage

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go get -d -v

RUN go build -o /main main.go

ENTRYPOINT [ "/main" ]

# Step 2: Deploy Stage
FROM alpine:latest

WORKDIR /

COPY --from=BuildStage /main .
COPY --from=BuildStage /app/.env .

ENTRYPOINT [ "/main" ]