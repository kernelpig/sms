all: build

export GOPATH := $(GOPATH)
export GOBIN := $(CURDIR)/bin

CURRENT_GIT_GROUP := wangqingang
CURRENT_GIT_REPO := sms_lib

build:
	protoc -I pb pb/sms.proto --go_out=plugins=grpc:pb

test:
	go test -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)

clean:
	@rm -rf bin

.PHONY: test build clean
