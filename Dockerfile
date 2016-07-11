FROM golang

# Simplify making releases
RUN apt-get update \
	&& apt-get install -yq zip bzip2
RUN wget -O github-release.bz2 https://github.com/aktau/github-release/releases/download/v0.6.2/linux-amd64-github-release.tar.bz2 \
        && tar jxvf github-release.bz2 \
        && mv bin/linux/amd64/github-release /usr/local/bin/ \
        && rm github-release.bz2


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
RUN go test -v ./...

RUN go build -o gendoc main.go \
	&& GOOS=windows GOARCH=amd64 go build -o gendoc.exe main.go \
	&& GOOS=darwin GOARCH=amd64 go build -o gendoc.app main.go \
	&& zip gendoc.zip gendoc gendoc.exe gendoc.app
