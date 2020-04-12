all: test build docker helm
build: 
	go vet ../cmd/huicd
	go build -o huicd -v ../cmd/huicd
test: 
	go test ../cmd/huicd
run:
	go build -o huicd -v ../cmd/huicd
	DATA_DIR=/tmp ./huicd server
docker:
	docker build . -t loqutus/huicd:0.0.1
	docker push loqutus/huicd:0.0.1
helm:
	helm delete huicd
	helm install huicd ../deployments/huicd