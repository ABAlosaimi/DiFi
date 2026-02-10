# Build stage
FROM golang:1.25-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o difi ./cmd/main

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests (if needed)
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/difi .

# Create directories for input/output operations
RUN mkdir -p /input /output

# Set the entrypoint
ENTRYPOINT ["./difi"]

# Default command (shows usage if no arguments provided)
CMD []
