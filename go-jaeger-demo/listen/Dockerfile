FROM golang
ADD . /go/src/listen
WORKDIR /go/src/listen
RUN go get listen
RUN go install
ENTRYPOINT ["/go/bin/listen"]


