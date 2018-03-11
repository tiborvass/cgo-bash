FROM golang:1.10
RUN apt-get update && apt-get -y install \
	libncurses5-dev bison patch
WORKDIR /go/src/github.com/tiborvass/cgo-bash
