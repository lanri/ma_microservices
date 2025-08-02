# syntax=docker/dockerfile:1

# Base image
FROM golang:1.24.5

# Set working directory
WORKDIR /app

# Copy go mod and sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source
COPY . .

# Build the Go app
RUN go build -o diktum-service

# Expose port 8083
EXPOSE 8083

# Command to run the executable
CMD ["./diktum-service"]
