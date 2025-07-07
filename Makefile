include .env
export

run:
	go run cmd/app/main.go
.PHONY: run

proto-v1: ### generate source files from proto
	protoc --go_out=. \
		--go-grpc_out=. \
		api/proto/v1/*.proto
.PHONY: proto-v1