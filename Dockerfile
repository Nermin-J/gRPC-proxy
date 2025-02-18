# Use a lightweight base image for building
FROM golang:1.20 as builder

# Set working directory
WORKDIR /app

# Copy Go files
COPY main.go go.mod go.sum ./

# Disable CGO for fully static binary
ENV CGO_ENABLED=0

# Download dependencies
RUN go mod download

# Build the Go app
RUN go build -o go-metrics-app

# Use a minimal image for production
FROM alpine:latest

# Set working directory
WORKDIR /app

# Install required certificates (if needed)
RUN apk add --no-cache ca-certificates

# Copy the compiled binary from the builder stage
COPY --from=builder /app/go-metrics-app .

# Expose the port Prometheus will scrape
EXPOSE 8080

# Run the application
CMD ["./go-metrics-app"]