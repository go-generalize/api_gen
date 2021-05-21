# RECOMMENDED: YOU SHOULD USE FIXED VERSION FOR API_GEN
# API_GEN_VERSION := "v2.0.0"
# bootstrap:
#	mkdir -p ./bin
#	curl -L -o ./bin/client_generator.tar.gz https://github.com/go-generalize/api_gen/releases/download/$(API_GEN_VERSION)/api_gen_$(shell uname -s)_$(shell uname -m).tar.gz
#	cd ./bin && \
#		tar xzf api_gen.tar.gz && \
#		rm *.tar.gz

.PHONY: goimports
goimports:
	cd /tmp && go get golang.org/x/tools/cmd/goimports

.PHONY: bootstrap
bootstrap: goimports
	mkdir -p ./bin
	curl -L -o ./bin/api_gen.tar.gz https://github.com/go-generalize/api_gen/releases/latest/download/api_gen_$(shell uname -s)_$(shell uname -m).tar.gz
	cd ./bin && \
		tar xzf api_gen.tar.gz && \
		rm *.tar.gz

.PHONY: server_generate
server_generate:
	cd ./backend && ../bin/api_gen server ./interfaces

.PHONY: client_generate
client_generate:
	mkdir -p ./frontend/api
	cd ./frontend/api && ../../bin/api_gen client ts ../../backend/interfaces

.PHONY: generate
generate: server_generate client_generate
