FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o tembusptn-user-service .


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/tembusptn-user-service .

EXPOSE 3000

CMD ["./notes-app"]