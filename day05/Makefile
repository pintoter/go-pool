.PHONY: test
test: clean
	cd ./task && go test -coverprofile=coverage.out
	cd ./task && go tool cover -html=coverage.out

.PHONY: clean
clean: 
	-rm ./task/coverage.out