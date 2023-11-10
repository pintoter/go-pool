.PHONY: gen run-server run-client
gen:
	mkdir -p pkg
	protoc --go_out=pkg --go_opt=paths=source_relative --go-grpc_out=pkg/ --go-grpc_opt=paths=source_relative api/proto/transmitter.proto

clean:
	rm pkg/api/proto/*.go

build:
	go build cmd/client/client.go
	go build cmd/server/server.go

run-server:
	./server

run-client:
	./client
