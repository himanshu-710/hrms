FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/hrms-server ./cmd/server

FROM alpine:3.22

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /bin/hrms-server /app/hrms-server
COPY --from=builder /app/migrations /app/migrations

ENV PORT=8080

EXPOSE 8080

CMD ["/app/hrms-server"]
