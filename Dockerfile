# Build Stage
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
# CGO_ENABLED=0 is important for creating statically linked binaries
RUN CGO_ENABLED=0 GOOS=linux go build -o /adarsha-server ./cmd/main.go

# Run Stage
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=builder /adarsha-server .

# Copy static assets and views
COPY static/ ./static/
COPY views/ ./views/
COPY uploads/ ./uploads/

# Expose the port the application will run on
EXPOSE 8080

# Command to run the executable
CMD ["./adarsha-server"]
