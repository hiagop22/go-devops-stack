FROM golang:1.26.3-alpine3.23 AS builder

WORKDIR /app

COPY go.mod ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -o http-server-projeto-korp .

FROM alpine:3.23

WORKDIR /app

COPY --from=builder /app/http-server-projeto-korp .

EXPOSE 8080

CMD [ "./http-server-projeto-korp" ]