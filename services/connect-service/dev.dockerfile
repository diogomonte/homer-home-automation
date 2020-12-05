FROM golang:1.14

ADD /services/connect-service/ /app
ADD /go.mod /app
ADD /go.sum /app

WORKDIR /app

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /connect-service

EXPOSE 8081

CMD ["ls"]
CMD ["/connect-service"]