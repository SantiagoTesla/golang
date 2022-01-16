FROM golang:1.17-alpine
WORKDIR /go/src
COPY go.mod /go/src
COPY *.go /go/src
RUN go build -o /go/bin/httpserver
EXPOSE 80
CMD [ "/go/bin/httpserver" ]