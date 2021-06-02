.PHONY: dev
dev: export LOCAL_ENV = 1
dev:
	go run cmd/api_gen/*.go

.PHONY: init
init: bootstrap
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
	)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.24.0

.PHONY: test
test:
	go test ./... -v

.PHONY: api_gen
api_gen:
	go build -o ./bin/api_gen ./cmd/api_gen

.PHONY: gen_samples
gen_samples: api_gen
	bash ./samples/generate.sh

.PHONY: lint
lint:
	golangci-lint run --config .github/.golangci.yml
