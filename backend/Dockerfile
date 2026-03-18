
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY . .

RUN go build -o istla main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/istla .
COPY --from=builder /app/ddl.sql .
COPY --from=builder /app/dml.sql .

RUN apk add --no-cache postgresql-client

COPY init.sh /app/init.sh

RUN chmod +x /app/init.sh

CMD ["./init.sh"]