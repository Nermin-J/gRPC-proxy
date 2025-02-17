# Use a lightweight base image
FROM golang:1.20 as builder

# Set working directory
WORKDIR /app

# Copy Go files
COPY main.go go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Build the Go app
RUN go build -o go-metrics-app

# Use a minimal image for production
FROM alpine:latest

WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/go-metrics-app .

# Expose the port Prometheus will scrape
EXPOSE 8080

# Run the application
CMD ["./go-metrics-app"]