FROM golang:1.11.0 as builder

WORKDIR /go/src/github.com/andreylm/learning/microservices/consignment-service

COPY . .

RUN go get -d

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

WORKDIR /app


COPY --from=builder /go/src/github.com/andreylm/learning/microservices/consignment-service/consignment-service .

CMD ["./consignment-service"]