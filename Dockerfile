# syntax=docker/dockerfile:1

FROM golang:1.16-alpine AS build

RUN mkdir -p /app
WORKDIR /app

COPY . .

RUN go build -mod=vendor -ldflags '-s -w' -o /api-iso-to-json main.go

EXPOSE 8080

ENTRYPOINT [ "/api-iso-to-json" ]

