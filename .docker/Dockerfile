# Stage 1: Builda go app
FROM golang:1.22-bullseye AS build

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o api-golang .

# Stage 2: Build final image
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=build /app/api-golang .
COPY config.yaml .

EXPOSE 11000
CMD ["./api-golang"]