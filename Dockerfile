# Use an official Go image as the base image
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o migrate .

# Use a minimal base image to run the application
FROM alpine:3.18

# Set up environment variables for PostgreSQL connection
ENV DB_HOST=localhost \
    DB_PORT=5432 \
    DB_USER=postgres \
    DB_PASSWORD=password \
    DB_NAME=postgres \
    MIGRATION_DIR=/migrations

# Set the working directory inside the container
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/migrate /app/migrate

# Copy the migration files
COPY migrations /migrations

# Make sure the binary is executable
RUN chmod +x /app/migrate

# Set the default command to run when the container starts
CMD ["/app/migrate", \
     "--dbHost=${DB_HOST}", \
     "--dbPort=${DB_PORT}", \
     "--dbUser=${DB_USER}", \
     "--dbPassword=${DB_PASSWORD}", \
     "--dbName=${DB_NAME}", \
     "--migrationDir=${MIGRATION_DIR}"]

