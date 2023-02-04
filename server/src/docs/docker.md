```docker
# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Specify the command to run the application
CMD ["./main"]
```

```yaml
version: '3'
services:
  server:
    build:
      context: ./server
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      POSTGRES_HOST: db
      POSTGRES_PORT: 5432
      POSTGRES_DB: 
      POSTGRES_USER: 
      POSTGRES_PASSWORD: 
  db:
    image: postgres
    ports:
      - "5432:5432"
    environment:
      POSTGRES_USER: 
      POSTGRES_PASSWORD: 
  client:
    build:
      context: ./client
      dockerfile: Dockerfile
    ports:
      - "3000:3000"
```