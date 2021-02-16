# Start from golang:1.12-alpine base image
FROM golang:1.12-alpine

RUN apk update && apk upgrade && apk add --update alpine-sdk && \
    apk add --no-cache bash git openssh make cmake 

WORKDIR /app

COPY . /app

# Build the Go app
RUN go build 

# Run the executable
CMD ["./gameObjectBackend"]