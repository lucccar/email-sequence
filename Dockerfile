# Stage 1: Build
FROM golang:1.23.0-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Install curl and tar, then set up the Go application
RUN apk add --no-cache curl tar

# Copy Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code and build
COPY . .
RUN go build -o main .

# Install migrate tool
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz -o migrate.tar.gz && \
    tar -xzf migrate.tar.gz && \
    mv migrate /usr/local/bin/migrate && \
    chmod +x /usr/local/bin/migrate && \
    rm migrate.tar.gz

# Stage 2: Runtime
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary and migrate tool from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate
COPY migrations /app/migrations
RUN chmod +x /app/migrations

# Copy the wait-for-it script into the container
COPY wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x /app/wait-for-it.sh

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["/app/wait-for-it.sh", "postgres:5432", "--", "/app/main"]
