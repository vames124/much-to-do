# Stage 1: Build the Go binary
FROM golang:1.25.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go mod and sum files first (for caching dependencies)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary (disabling CGO for a statically linked binary)
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create the minimal production image
FROM alpine:latest

WORKDIR /root/

# Install CA certificates for secure external connections (like MongoDB)
RUN apk --no-cache add ca-certificates

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 (which matches our ALB Target Group)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]