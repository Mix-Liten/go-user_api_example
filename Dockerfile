# build stage
FROM golang:1.19.3-alpine3.16 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

ENTRYPOINT ./main