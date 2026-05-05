# --- Build stage ---
FROM golang:1.25.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main

# --- Run stage ---
FROM alpine:3.23.4

WORKDIR /app

COPY --from=builder /main /main

CMD ["/main"]
