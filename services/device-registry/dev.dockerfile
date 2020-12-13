FROM golang:1.14

ADD /services/device-registry/ /app
ADD /go.mod /app
ADD /go.sum /app

WORKDIR /app

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /device-registry

EXPOSE 8081

CMD ["ls"]
CMD ["/device-registry"]