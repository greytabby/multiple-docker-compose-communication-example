FROM golang:1.15-alpine
WORKDIR /app

COPY . .

RUN go build -o server .

EXPOSE 8080

CMD ["/app/server"]