build:
	protoc -I. --go_out=plugins=micro:$$GOPATH/src/github.com/andreylm/learning/microservices/user-service proto/user/user.proto
	GOOS=linux GOARCH=amd64 go build -o user-service main.go datastore.go handler.go repository.go 
	# docker build -t user-service .