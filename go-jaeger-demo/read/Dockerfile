FROM golang
ADD . /go/src/read
WORKDIR /go/src/read
RUN go get read
RUN go install
ENTRYPOINT ["/go/bin/read"]

