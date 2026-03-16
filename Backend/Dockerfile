# Stage 1: Build the Go binary
FROM golang:1.25-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
# CGO_ENABLED=0 makes it a statically linked binary (better for alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Final lightweight image
FROM alpine:latest

# Add ca-certificates for HTTPS requests (RSS fetching)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port (Render uses the PORT env var, but this is good practice)
EXPOSE 8080

# Command to run
CMD ["./main"]