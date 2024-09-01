# Stage 1: Builda go app
FROM golang:1.22-bullseye AS build

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o api-golang .

# Stage 2: Build final image
FROM debian:bullseye-slim

ARG HEADER=*
ARG ORIGIN=http://localhost:4000
ARG MONGO_CONNECTION_STRING=mongodb://admin:123456@Mongo:27017
ARG SQL_CONNECTION_STRING=root:123456@tcp(MySql:3306)
ARG KAFKA_BROKER=Kafka:29092

ENV HEADER=$HEADER
ENV ORIGIN=$ORIGIN
ENV MONGO_CONNECTION_STRING=$MONGO_CONNECTION_STRING
ENV SQL_CONNECTION_STRING=$SQL_CONNECTION_STRING
ENV KAFKA_BROKER=$KAFKA_BROKER

WORKDIR /app
COPY --from=build /app/api-golang .
COPY config.yaml .

EXPOSE 11000
CMD ["./api-golang"]