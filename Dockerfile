FROM golang:1.15.7-alpine
COPY . $GOPATH/src/github.com/bm-krishna-source/tenant-client
WORKDIR $GOPATH/src/github.com/bm-krishna-source/tenant-client
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o ./cmd ./cmd/main.go
EXPOSE 3030
CMD [ "./cmd/main" ]