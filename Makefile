.PHONY: build shell docker-build docker run release

RELEASE_DATE=$(shell date +%F)
COMMIT_HASH=$(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE=$(date +%FT%T%z)
LDFLAGS=-ldflags "-X github.com/SvenDowideit/gendoc.CommitHash=${COMMIT_HASH} -X github.com/SvenDowideit/gendoc.Version=${RELEASE_DATE}"

AWSTOKENSFILE ?= ../aws.env
-include $(AWSTOKENSFILE)
export GITHUB_USERNAME GITHUB_TOKEN

build:
	go build $(LDFLAGS) -o gendoc main.go

shell: docker-build
	docker run --rm -it -v $(CURDIR):/go/src/github.com/SvenDowideit/gendoc gendoc bash

docker-build:
	rm -f gendoc.gz
	docker build \
		--build-arg RELEASE_DATE=$(RELEASE_DATE) \
		--build-arg COMMIT_HASH=$(COMMIT_HASH) \
		-t gendoc .

docker: docker-build
	docker run --name gendoc-build gendoc
	docker cp gendoc-build:/go/src/github.com/SvenDowideit/gendoc/gendoc.zip .
	docker rm gendoc-build
	rm -f gendoc
	unzip -o gendoc.zip

run:
	./gendoc .



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

fmt:
	docker run --rm -it -v $(shell pwd):/data -w /data golang go fmt
	docker run --rm -it -v $(shell pwd):/data -w /data/commands golang go fmt
	docker run --rm -it -v $(shell pwd):/data -w /data/allprojects golang go fmt
	docker run --rm -it -v $(shell pwd):/data -w /data/render golang go fmt
