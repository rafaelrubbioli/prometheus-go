FROM golang:1.18 as builder

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o myapp

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY  --from=builder /app/myapp .

CMD ["./myapp"]
