FROM golang

ENV GOPATH /go
ENV USER root

WORKDIR /go/src/github.com/SvenDowideit/gendoc

RUN go get github.com/Sirupsen/logrus \
    && go get github.com/codegangsta/cli \
    && go get github.com/cloudfoundry-incubator/candiedyaml \
    && go get github.com/google/go-github/github \
    && go get golang.org/x/oauth2 \
    && go get github.com/miekg/mmark

ADD . /go/src/github.com/SvenDowideit/gendoc
RUN go get -d -v
RUN go build --race -o gendoc main.go
RUN go test -v ./...

