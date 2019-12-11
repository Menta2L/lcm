VERSION  := $(shell git describe --tags --exact-match 2> /dev/null || git rev-parse --short HEAD || echo "unknown")
REVISION := $(shell git rev-parse HEAD)

LINUX_LDFLAGS := -s -w -extldflags "-static"
DARWIN_LDFLAGS := -s -w
LINKFLAGS := \
	-X "github.com/menta2l/lcm/pkg/build.tag=$(VERSION)" \
	-X "github.com/menta2l/lcm/pkg/build.rev=$(REVISION)"
override LINUX_LDFLAGS += $(LINKFLAGS)

all: clean deps proto test build

deps:
	go mod download

build:
	GOOS=linux GOARCH=amd64 go build -a -tags netgo -installsuffix netgo -ldflags '$(LINUX_LDFLAGS)'

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go get -u -v github.com/golang/protobuf/protoc-gen-go
	for file in $$(git ls-files '*.proto'); do \
		protoc -I $$(dirname $$file) --go_out=plugins=grpc:pkg/api $$file; \
	done

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic -cpu 1,4 github.com/menta2l/lcm/...

clean:
	rm -rf ./dist
	rm -f coverage.txt
	go clean -i github.com/menta2l/lcm/...

.PHONY: \
	all \
	deps \
	build \
	proto \
	test \
	clean
