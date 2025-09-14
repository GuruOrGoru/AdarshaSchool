FROM golang:1.25.1-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .


FROM alpine:latest

WORKDIR /app

COPY --from=builder /adarsha-server .

COPY static/ ./static/
COPY views/ ./views/
COPY uploads/ ./uploads/

EXPOSE 8080

CMD ["./adarsha-server"]
