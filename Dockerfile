# Stage 1: Build
FROM golang:1.27-alpine AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Runner
FROM alpine:latest

WORKDIR /app

# Copy binary dari stage builder
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]