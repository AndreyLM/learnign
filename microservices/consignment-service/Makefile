build:
	protoc -I. --go_out=plugins=micro:$$GOPATH/src/github.com/andreylm/learning/microservices/consignment-service proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build -o consignment-service main.go datastore.go handler.go repository.go 
	docker build -t consignment-service .

run:
	docker run -p 50051:50051 --network host -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-service
