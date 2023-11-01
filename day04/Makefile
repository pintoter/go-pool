build:
		go build -o candy-server ./cmd/app/main.go
		go build -o candy-client ./client/main.go

run:
		./candy-server

runtls:
		./candy-server -tls

runclient:
		./candy-client -k AA -c 2 -m 30

clean:
		rm candy-clint candy-server