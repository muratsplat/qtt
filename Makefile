exec=qtt

.DEFAULT_GOAL := build

configure: clean
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -vendor-only -v
test: configure
	go vet ./...
	go test ./...
build: configure
	go build -v  -o ${exec}
docker-build: build
	docker build -t qtt:latest . 
	docker tag qtt:latest  muod/qtt:latest
	echo  ${DOCKER_PASSWORD} | docker login -u ${DOCKER_USERNAME} --password-stdin
	docker push muod/qtt:latest
clean:
	rm -f ${exec}

