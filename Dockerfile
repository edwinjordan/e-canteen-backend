# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o ecanteen-api main.go

# Run stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy binary from builder
COPY --from=builder /app/ecanteen-api .
COPY --from=builder /app/.env .

# Expose port (adjust if your app uses a different port)
EXPOSE 3000

# Run the application
CMD ["./ecanteen-api"]
