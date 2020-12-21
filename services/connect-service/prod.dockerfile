FROM golang:1.14-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
ADD /services/connect-service /app
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN env GOOS=linux GOARCH=arm GOARM=5 go build -o /services/connect-service .

# This container exposes port 8081 to the outside world
EXPOSE 8081
# Run the binary program produced by `go install`
CMD ["/services/connect-service"]