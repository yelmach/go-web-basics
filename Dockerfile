# Set the baseImage to use for this image Golang and Alpine
FROM golang:1.22.5-alpine

# Set the working directory for my program files
WORKDIR /app

COPY . .

# Execute this commands on this image as a new layer
RUN go build -o myapp main.go

RUN apk add bash

LABEL description="ASCII Art Web Application consists in creating and running a server, in which it will be possible to use a web GUI"

CMD [ "./myapp" ]