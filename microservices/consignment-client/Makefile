build:
	GOOS=linux GOARCH=amd64 go build -o consignment-client main.go
	docker build -t consignment-cli .

run:
	docker run --network host -e MICRO_REGISTRY=mdns consignment-cli