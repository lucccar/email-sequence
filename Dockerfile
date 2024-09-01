# Use the official Golang image as the base image
FROM golang:1.23.0-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal image to run the application
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Set environment variables, if needed
# ENV ENV_VAR_NAME=value

# Expose the application port (e.g., 8080)
EXPOSE 8080

# Command to run the application
CMD ["./main"]
