FROM golang:latest
ADD . /go/src/random-stuff-service
WORKDIR /go/src/random-stuff-service
RUN go build -o random-stuff-service main/main.go
EXPOSE 5000
CMD ["./random-stuff-service"]