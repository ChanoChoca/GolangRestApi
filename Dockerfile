# ---------- build ----------
FROM golang:alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# ---------- runtime ----------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

COPY .env .env

EXPOSE 8080

CMD ["./app"]
