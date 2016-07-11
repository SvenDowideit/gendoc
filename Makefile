
# Adds build information from git repo
#
# as suggested by tatsushid in
# https://github.com/spf13/hugo/issues/540

COMMIT_HASH=`git rev-parse --short HEAD 2>/dev/null`
BUILD_DATE=`date +%FT%T%z`
LDFLAGS=-ldflags "-X github.com/spf13/hugo/hugolib.CommitHash=${COMMIT_HASH} -X github.com/spf13/hugo/hugolib.BuildDate=${BUILD_DATE}"

AWSTOKENSFILE ?= ../aws.env
-include $(AWSTOKENSFILE)
export GITHUB_USERNAME GITHUB_TOKEN

build:
	go build -o gendoc main.go

shell: docker-build
	docker run --rm -it -v $(CURDIR):/go/src/github.com/SvenDowideit/gendoc gendoc bash

docker-build:
	rm -f gendoc.gz
	docker build -t gendoc .

docker: docker-build
	docker run --name gendoc-build gendoc
	docker cp gendoc-build:/go/src/github.com/SvenDowideit/gendoc/gendoc.zip .
	docker rm gendoc-build
	rm -f gendoc
	unzip -o gendoc.zip

run:
	./gendoc .


RELEASE_DATE=`date +%F`

release: docker
	# TODO: check that we have upstream master, bail if not
	docker run --rm -it -e GITHUB_TOKEN gendoc \
		github-release release --user SvenDowideit --repo gendoc --tag $(RELEASE_DATE)
	docker run --rm -it -e GITHUB_TOKEN gendoc \
		github-release upload --user SvenDowideit --repo gendoc --tag $(RELEASE_DATE) \
			--name gendoc \
			--file gendoc
	docker run --rm -it -e GITHUB_TOKEN gendoc \
		github-release upload --user SvenDowideit --repo gendoc --tag $(RELEASE_DATE) \
			--name gendoc-osx \
			--file gendoc.app
	docker run --rm -it -e GITHUB_TOKEN gendoc \
		github-release upload --user SvenDowideit --repo gendoc --tag $(RELEASE_DATE) \
			--name gendoc.exe \
			--file gendoc.exe
