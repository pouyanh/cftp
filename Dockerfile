FROM golang:1.7.1

RUN go get -v github.com/codeskyblue/fswatch
ADD . $GOPATH/src/github.com/pouyanh/chunked-file-transfer

WORKDIR $GOPATH/src/github.com/pouyanh/chunked-file-transfer
CMD $GOPATH/bin/fswatch
