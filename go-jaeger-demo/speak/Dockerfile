FROM golang
ADD . /go/src/speak
WORKDIR /go/src/speak
RUN go get speak
RUN go install
ENTRYPOINT ["/go/bin/speak"]

