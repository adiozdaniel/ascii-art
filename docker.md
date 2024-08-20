# Dockerizing the Ascii-Web

## Step 1: Build the Go application

- `FROM golang:1.20-alpine AS builder`

## Step 2: Set the Current Working Directory inside the container

- `WORKDIR /app`

## Step 3: Copy the module initializer go.mod file

- `COPY go.mod ./`

## Step 4: Download and cache all dependencies

- `RUN go mod download`

## Step 5: Copy the source code into the container

- `COPY . .`

## Step 6: Build the Ascii-Web Application

- `RUN go build -o main ./cmd/web/web.go`

## Step 7: Create an image for the Go application

- `FROM alpine:latest`

## Step 8: Set the Current Working Directory inside the container

- `WORKDIR /root/`

## Step 9: Copy the pre-built binary file from the builder stage

- `COPY --from=builder /app/main .`

## Step 10: Make port 8080 available to the world

- `EXPOSE 8080`

## Step 11: Run the binary program

- `CMD ["./main"]`

## Step 12: Create a .dockerignore file

- To exclude unnecessary build files during in the image

## Step 13: Build the Docker Image

- `docker build --no-cache -t ascii-web:v5.0.1 .` To ignore caching.

## Step 14: Run the Docker Container

- `sudo docker run -p 8080:8080 ascii-web:v5.0.1`
