# RECOMMENDED: YOU SHOULD USE FIXED VERSION FOR API_GEN
# API_GEN_VERSION := "v1.0.0"
# bootstrap:
# 	mkdir -p ./bin
# 	curl -L -o ./bin/client_generator.tar.gz https://github.com/go-generalize/api_gen/releases/download/$(API_GEN_VERSION)/client_generator_$(shell uname -s)_$(shell uname -m).tar.gz
# 	curl -L -o ./bin/server_generator.tar.gz https://github.com/go-generalize/api_gen/releases/download/$(API_GEN_VERSION)/server_generator_$(shell uname -s)_$(shell uname -m).tar.gz

.PHONY: goimports
goimports:
	cd /tmp && go get golang.org/x/tools/cmd/goimports

.PHONY: bootstrap
bootstrap: goimports
	mkdir -p ./bin
	curl -L -o ./bin/client_generator.tar.gz https://github.com/go-generalize/api_gen/releases/latest/download/client_generator_$(shell uname -s)_$(shell uname -m).tar.gz
	curl -L -o ./bin/server_generator.tar.gz https://github.com/go-generalize/api_gen/releases/latest/download/server_generator_$(shell uname -s)_$(shell uname -m).tar.gz
	cd ./bin && \
		tar xzf server_generator.tar.gz && \
		tar xzf client_generator.tar.gz && \
		rm *.tar.gz

.PHONY: server_generate
server_generate:
	cd ./backend && ../bin/server_generator ./interfaces

.PHONY: client_generate
client_generate:
	mkdir -p ./frontend/api
	cd ./frontend/api && ../../bin/client_generator ../../backend/interfaces

.PHONY: generate
generate: server_generate client_generate
