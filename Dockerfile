# Build stage
FROM golang:1.21-bookworm AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
# CGO_ENABLED=1 is required for go-sqlite3
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/bin/publisher ./cmd/publisher

# Final stage
FROM debian:bookworm-slim

# Install runtime dependencies
# sqlite3 lib, ca-certificates for HTTPS, curl/jq for ngrok
RUN apt-get update && apt-get install -y \
    sqlite3 \
    ca-certificates \
    curl \
    jq \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/bin/publisher .

# Copy data directory (required for libraries.json and posted.json)
COPY data/ ./data/

# Copy assets (required for image generation)
COPY internal/image/assets/ ./internal/image/assets/

# Create a volume for persistent data if needed
# VOLUME /app/data

# Copy scripts
COPY scripts/ ./scripts/
RUN chmod +x scripts/entrypoint.sh

# Entrypoint to handle setup
ENTRYPOINT ["/app/scripts/entrypoint.sh"]

# Run the binary
CMD ["./publisher"]
