.PHONY: generate
generate:
	server_generator ./interfaces
	go generate ./...

.PHONY: dev
dev: generate	
	go run main.go