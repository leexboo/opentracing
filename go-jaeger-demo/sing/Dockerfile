FROM golang
ADD . /go/src/sing
WORKDIR /go/src/sing
RUN go get sing
RUN go install
ENTRYPOINT ["/go/bin/sing"]

