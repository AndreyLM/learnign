build:
	protoc -I. --go_out=plugins=micro:$$GOPATH/src/github.com/andreylm/learning/microservices/vessel-service proto/vessel/vessel.proto
	GOOS=linux GOARCH=amd64 go build -o vessel-service main.go datastore.go handler.go repository.go
	docker build -t vessel-service .

run:
	docker run -p 50052:50053 --network host -e MICRO_SERVER_ADDRESS=:50053 -e MICRO_REGISTRY=mdns vessel-service