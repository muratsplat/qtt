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
clean:
	rm -f ${exec}

