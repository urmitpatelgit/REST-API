# --- Stage 1: Build the binary ---
FROM golang:1.26-bookworm AS builder

# Install build essentials required for SQLite3 (CGO)
RUN apt-get update && apt-get install -y --no-install-recommends gcc g++ libc6-dev && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Enable CGO for SQLite3 compilation
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# --- Stage 2: Final minimal production image ---
# Using the matching Debian base to guarantee C library/CGO compatibility
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

# Copy the built binary from stage 1
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
