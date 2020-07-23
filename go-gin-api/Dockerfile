FROM golang
ADD . /go/src/go-gin-api
WORKDIR /go/src/go-gin-api
RUN go get go-gin-api
RUN go install
ENTRYPOINT ["/go/bin/go-gin-api"]


