# --- Build Stage ---
FROM golang:1.26-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build from root (.) instead of ./cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server .

# --- Runtime Stage ---
FROM alpine:3.20

RUN apk add --no-cache ca-certificates && \
    adduser -D -g '' omni

USER omni
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .
# Note: Since your tree didn't show a 'migrations' folder, 
# I've commented this out. Uncomment it once you create that folder!
# COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./server"]