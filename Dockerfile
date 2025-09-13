# Use the official Go image as build stage
FROM golang:1.19-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application for AMD64 architecture (ECS Fargate requirement)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

# Use minimal alpine image for final stage
FROM alpine:latest

# Install ca-certificates and wget for HTTPS requests and health checks
RUN apk --no-cache add ca-certificates wget

# Set working directory
WORKDIR /tmp/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Expose port
EXPOSE 5655

# Run the application
CMD ["./main"]