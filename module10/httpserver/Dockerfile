FROM golang:1.17-alpine AS build
WORKDIR /build
COPY go.mod .
COPY *.go .
RUN go build -o httpserver

FROM alpine:latest
WORKDIR /
COPY --from=build /build/httpserver /httpserver
EXPOSE 80
ENTRYPOINT [ "./httpserver" ]
