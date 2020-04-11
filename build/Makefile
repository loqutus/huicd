all: test build
build: 
	go ver ../cmd/huicd
	go build -o huicd -v ../cmd/huicd
test: 
	go test ../cmd/huicd
run:
	go build -o huicd -v ../cmd/huicd
	DATA_DIR=/tmp ./huicd server
docker:
	docker build . -t loqutus/huicd
	docker push loqutus/huicd