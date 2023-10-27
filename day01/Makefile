build:
	go build -o readDB ./cmd/ex00/main.go
	go build -o compareDB ./cmd/ex01/main.go
	go build -o compareFS ./cmd/ex02/main.go

deps:
	go mod tidy

clean:
	@rm -rf readDB compareDB compareFS