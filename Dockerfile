FROM golang:1.23-alpine AS builder

RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

ENV CGO_ENABLED=1
ENV PORT=3000

RUN go build -o quotesly cmd/server/main.go

FROM alpine:latest

RUN apk add --no-cache sqlite

WORKDIR /app

COPY --from=builder /app/quotesly .

EXPOSE 3000

CMD ["./quotesly"]