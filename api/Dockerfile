# ---------- STAGE 1: Build ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# ---------- STAGE 2: Run ----------
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
COPY .env .env

EXPOSE 8080
CMD ["./main"]
