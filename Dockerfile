# Stage 1: Build Go binary
FROM golang:1.25.1-alpine AS builder

# Install git and ca-certificates for Go modules (needed in alpine)
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o adarsha-server ./cmd

# Stage 2: Minimal runtime image
FROM alpine:latest

# Install CA certificates for HTTPS requests if needed
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/adarsha-server .

# Copy other assets
COPY static/ ./static/
COPY views/ ./views/
COPY uploads/ ./uploads/

# Expose the port your app listens on
EXPOSE 8080

# Command to run the binary
CMD ["./adarsha-server"]
