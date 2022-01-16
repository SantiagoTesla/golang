FROM golang:1.17-alpine
WORKDIR /httpserver
ADD *.go /httpserver
RUN go build -o httpserver
EXPOSE 80
ENTRYPOINT /httpserver