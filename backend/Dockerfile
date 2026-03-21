
FROM golang:1.26-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o istla main.go


FROM alpine:latest

RUN apk add --no-cache ca-certificates postgresql-client

WORKDIR /app

COPY --from=builder /app/istla .
COPY --from=builder /app/ddl.sql .
COPY --from=builder /app/dml.sql .
COPY --from=builder /app/init.sh .

COPY .env .

RUN mkdir -p resources imagen pdf

RUN chmod +x /app/init.sh /app/istla


EXPOSE 6000

CMD ["./istla"]