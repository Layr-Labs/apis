.PHONY: build clean test

GO = $(shell which go)
BIN = ./bin

PROTO_OPTS=--proto_path=protos --go_out=paths=source_relative:protos

all: build

.PHONY: deps/dev
deps/dev:
	${GO} get github.com/grpc-ecosystem/protoc-gen-grpc-gateway-ts
	cd protos && buf dep update

.PHONY: deps
deps/go:
	${GO} install github.com/vektra/mockery/v2@v2.42.3
	${GO} install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	${GO} install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	${GO} mod tidy


deps-linux: deps/go
	./scripts/installDeps.sh

deps: deps/go
	HOMEBREW_NO_AUTO_UPDATE=1 brew install bufbuild/buf/buf@1.33.0

.PHONY: proto
proto:
	buf generate protos
	rm -rf ts/src/generated/*
	cp -r ./gen/api-ts/* ./ts/src/generated

