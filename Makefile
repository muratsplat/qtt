exec=qtt

.DEFAULT_GOAL := build
configure: clean
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -vendor-only -v
	chmod +x fixGoDep.py
	./fixGoDep.py

test: configure
	go vet ./...
	go test ./...
build: configure
	go build -v  -o ${exec}
clean:
	rm -f ${exec}

