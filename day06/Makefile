.PHONY: build
build: 
	go build web/image/logo.go
	mv logo ./web/image/

.PHONY: run
run: clean build
	./web/image/logo
	mv amazing_logo.png ./web/image/
	-rm web/image/logo

.PHONY: clean
clean: 
	-rm web/image/logo web/image/amazing_logo.png

build-srv:
	go build ./cmd/app/main.go

run-srv: build-srv
	./app