# Stage 1: Builda go app
FROM golang:1.22 AS build

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o app .

# Stage 2: Build final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /app .

EXPOSE 11000
CMD ["./app"]