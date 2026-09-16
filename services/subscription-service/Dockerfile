FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o tembusptn-subscription-service .


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/tembusptn-subscription-service .
COPY --from=builder /app/migrations ./migrations

EXPOSE 3001

CMD ["./notes-app"]