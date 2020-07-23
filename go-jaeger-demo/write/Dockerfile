FROM golang
ADD . /go/src/write
WORKDIR /go/src/write
RUN go get write
RUN go install
ENTRYPOINT ["/go/bin/write"]

