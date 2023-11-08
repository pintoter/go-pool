.PHONY: gen
gen:
	mkdir -p pkg
	protoc --go_out=pkg --go_opt=paths=source_relative --go-grpc_out=pkg/ --go-grpc_opt=paths=source_relative api/proto/math.proto

clean:
	rm pb/*.go

run:
	go run main.go