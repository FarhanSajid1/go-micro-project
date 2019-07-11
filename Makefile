build-consignment:
	protoc --micro_out=. --go_out=. proto/consignment.proto

build-vessel:
	protoc --micro_out=. --go_out=. vessel-service/proto/vessel/vessel.proto

run-vessel:
	MICRO_REGISTRY=consul go run vessel-service/vessel-server/main.go

run-server:
	MICRO_REGISTRY=consul go run consignment-server/main.go

run-client:
	MICRO_REGISTRY=consul go run consignment-client/main.go

docker-build:
	docker build -t farhansajid2/go-micro .

compile-vessel:
	CGO_ENABLED=0 GOOS=linux go build -o main