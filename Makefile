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
			github.com/rakyll/statik \
	)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.24.0

.PHONY: installdeps
installdeps:
	go mod download
	mkdir -p ./bin
	go build -o ./bin/ ./tools/...

.PHONY: test
test:
	go test ./... -v

.PHONY: statik
statik:
	cd server_generator && \
	statik "-include=*.tmpl" -src=. && \
	gofmt -w ./statik/statik.go
	cd client_generator && \
	statik "-include=*.tmpl" -src=. && \
	gofmt -w ./statik/statik.go

.PHONY: server_generator
server_generator: statik
	go build -o ./bin/server_generator ./server_generator

.PHONY: client_generator
client_generator: statik
	go build -o ./bin/client_generator ./client_generator

build-release:
	$(shell bash ./scripts/build_release.sh)
