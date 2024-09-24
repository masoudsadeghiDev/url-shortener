# Build stage
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN cd frontend && npm install && npm run build
RUN go build -o url-shortener cmd/server/main.go

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/url-shortener .
COPY --from=builder /app/frontend/build ./frontend/build

# Set the entrypoint
CMD ["./url-shortener"]
