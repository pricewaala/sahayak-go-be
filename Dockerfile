# Start from the official golang image
FROM golang:1.21.5 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Start a new stage from scratch
FROM alpine:latest

# Install MongoDB client
RUN apk add --no-cache mongodb-tools

# Set the current working directory inside the container
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/app .

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./app"]
