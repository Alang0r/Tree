# Build stage
FROM golang:1.18.2-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go build -o informer informer/cmd/informer.go
# RUN apk add curl
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/informer .
# COPY --from=builder /app ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY lib/migrations/ ./migration
COPY informer/cmd/config.yaml .

EXPOSE 8080
CMD ["/app/informer"]
ENTRYPOINT ["/app/start.sh"]