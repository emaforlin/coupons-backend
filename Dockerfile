FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o couponsApp ./cmd/api

RUN chmod +x couponsApp

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/couponsApp .

CMD ["/app/couponsApp"]