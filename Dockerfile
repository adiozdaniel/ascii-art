FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/web/*.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY ./views /root/views
COPY ./.env /root/

EXPOSE 8080

CMD ["./main"]

LABEL version="v5.0.1"
LABEL frontend_engineer="josephineopondo5@gmail.com"
LABEL backend_engineer="andyovvo@gmail.com"
LABEL documentation="adiozdaniel@gmail.com"
