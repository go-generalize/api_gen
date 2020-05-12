PORT=8080
TEST_OPT=""

.PHONY: dev
dev: export LOCAL_ENV = 1
dev:
	go run cmd/*.go

gen:
	go generate ./...

.PHONY: init
init: bootstrap installdeps
	direnv allow
	cp ./scripts/pre-commit .git/hooks/

.PHONY: bootstrap
bootstrap:
	(cd $(mktemp -d); GO111MODULE=on \
			go get \
			github.com/golang/mock/gomock \
			github.com/golang/mock/mockgen \
			github.com/golang/protobuf/protoc-gen-go \
			github.com/favadi/protoc-go-inject-tag \
			github.com/moznion/go-errgen/cmd/errgen \
			github.com/OneOfOne/struct2ts/... \
	)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.24.0

.PHONY: installdeps
installdeps:
	go mod download
	mkdir -p ./bin
	go build -o ./bin/ ./tools/...

test:
	go test ./... -v