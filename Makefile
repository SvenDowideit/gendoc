
# Adds build information from git repo
#
# as suggested by tatsushid in
# https://github.com/spf13/hugo/issues/540

COMMIT_HASH=`git rev-parse --short HEAD 2>/dev/null`
BUILD_DATE=`date +%FT%T%z`
LDFLAGS=-ldflags "-X github.com/spf13/hugo/hugolib.CommitHash=${COMMIT_HASH} -X github.com/spf13/hugo/hugolib.BuildDate=${BUILD_DATE}"

build:
	go build -o gendoc main.go

shell: docker-build
	docker run --rm -it -v $(CURDIR):/go/src/github.com/SvenDowideit/gendoc gendoc bash

docker-build:
	rm -f gendoc.gz
	docker build -t gendoc .

docker: docker-build
	docker run --name gendoc-build gendoc gzip gendoc
	docker cp gendoc-build:/go/src/github.com/SvenDowideit/gendoc/gendoc.gz .
	docker rm gendoc-build
	rm -f gendoc
	gunzip gendoc.gz

run:
	./gendoc .

