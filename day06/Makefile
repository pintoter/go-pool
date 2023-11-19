.PHONY: build
build: 
	go build web/image/logo.go
	mv logo ./web/image/
	go build ./cmd/app/main.go

.PHONY: run
run: clean build
	./web/image/logo
	mv amazing_logo.png ./web/image/
	-rm web/image/logo
	./main

.PHONY: clean
clean: 
	-rm web/image/logo web/image/amazing_logo.png
	-rm main

build-srv:
	go build ./cmd/app/main.go

run-srv: build-srv
	./app