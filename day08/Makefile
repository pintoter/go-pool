.PHONY: test

test: clean
	go test -v -coverprofile cover.out
	go tool cover -html cover.out -o cover.html
	open cover.html

.PHONY: deps

deps: 
	go get golang.org/x/tools/cmd/cover

clean:
	-rm cover.html cover.out
