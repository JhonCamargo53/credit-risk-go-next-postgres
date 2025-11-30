FROM golang:1.25.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/server .

ENV PORT=4000
EXPOSE 4000

CMD ["./server"]
