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

helm-tiller:
	 curl -LO https://git.io/get_helm.sh
	 chmod 700 get_helm.sh
	 ./get_helm.sh
	 kubectl create serviceaccount --namespace kube-system tiller
	 kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
	 helm init --service-account tiller --upgrade

get-nginx:
	helm install stable/nginx-ingress --name my-nginx --set rbac.create=true