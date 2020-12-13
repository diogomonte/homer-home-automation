FROM golang:1.14

ARG service_name
RUN echo $service_name

ENV SERVICE $service_name
RUN echo $SERVICE
RUN echo 123

ADD /services/$SERVICE/ /app
ADD /go.mod /app
ADD /go.sum /app

WORKDIR /app

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /$SERVICE

EXPOSE 8081

CMD ["ls"]
CMD ["/$SERVICE"]