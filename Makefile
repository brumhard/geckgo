SHELL=/bin/bash -e -o pipefail
.DEFAULT_GOAL := all

PWD = $(shell pwd)
GOLANGCI_VERSION = 1.40.0

# TOOLS

# Go dependencies versioned through tools.go
GO_TOOLS = google.golang.org/protobuf/cmd/protoc-gen-go \
				google.golang.org/grpc/cmd/protoc-gen-go-grpc \
				github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
				github.com/bufbuild/buf/cmd/buf \
                github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking \
                github.com/bufbuild/buf/cmd/protoc-gen-buf-lint

define make-go-dependency
  # target template for go tools, can be referenced e.g. via /bin/<tool>
  bin/$(notdir $1):
	GOBIN=$(PWD)/bin go install $1
endef

# this creates a target for each go dependency to be referenced in other targets
$(foreach dep, $(GO_TOOLS), $(eval $(call make-go-dependency, $(dep))))

bin/golangci-lint: bin/golangci-lint-$(GOLANGCI_VERSION)
	@ln -sf golangci-lint-$(GOLANGCI_VERSION) $@

bin/golangci-lint-$(GOLANGCI_VERSION):
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b bin v$(GOLANGCI_VERSION)
	@mv bin/golangci-lint $@

# PROTOBUF DEPENDENCIES

protos:
	@mkdir -p protos/google/api
	curl -s https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto -o protos/google/api/annotations.proto
	curl -s https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto -o protos/google/api/http.proto

# MAIN TARGETS

.PHONY: all
all: bin/golangci-lint
all: protos bin/protoc-gen-go bin/protoc-gen-go-grpc bin/protoc-gen-grpc-gateway ## protobuf
all: bin/buf bin/protoc-gen-buf-breaking bin/protoc-gen-buf-lint ## buf
all: generate

.PHONY: clean
clean:
	@rm -rf protos bin

.PHONY: check
check: lint buf-lint

# SUB TARGETS

.PHONY: lint
lint: bin/golangci-lint
	bin/golangci-lint run

.PHONY: buf-lint
buf-lint: bin/buf bin/protoc-gen-buf-lint
	bin/buf lint

.PHONY: generate
generate: bin/buf
	PATH=$(PWD)/bin:$$PATH bin/buf generate --path api/geckgo/v1/geckgo.proto