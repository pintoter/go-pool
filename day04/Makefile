build:
		go build -o candy-server ./cmd/app/main.go
		go build -o candy-client ./cmd/client/main.go

run: build
		./candy-server

runtls: build
		./candy-server -tls

runclient:
		./candy-client -k AA -c 2 -m 30 -path https://candy.tld:3333/buy_candy

clean:
		rm candy-clint candy-server