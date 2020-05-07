FROM golang:latest
WORKDIR /go/src/groupbirthday
COPY . .
RUN mkdir bin
RUN	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/ groupbirthday

CMD ["./bin/groupbirthday"]