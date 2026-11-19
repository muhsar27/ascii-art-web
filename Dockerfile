# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ascii-art-web .

# Final runtime stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ascii-art-web .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/banners ./banners

EXPOSE 8080

CMD ["./ascii-art-web"]

