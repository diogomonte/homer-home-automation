FROM golang:1.14

WORKDIR /app
ADD . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o device-registry .

CMD ["/app/device-registry"]
