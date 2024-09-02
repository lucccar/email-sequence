FROM golang:1.23.0-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal image to run the application
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the wait-for-it script into the container
COPY wait-for-it.sh /app/wait-for-it.sh

# Make the wait-for-it script executable
RUN chmod +x /app/wait-for-it.sh

# Expose the application port (e.g., 8080)
EXPOSE 8080

# Command to run the application
CMD ["/app/wait-for-it.sh", "postgres:5432", "--", "./main"]
