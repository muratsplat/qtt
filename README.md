[![Build Status](https://travis-ci.com/muratsplat/qtt.svg?branch=master)](https://travis-ci.com/muratsplat/qtt)
# Qtt is simple message broker on based webhook

**This project is not completed yet. Don't use it on your production until further notice..**

This is experiment about handling MQTT mesage from clients. Main goal is all message forwarded by HTTP2 to another point. It is easy to deliver massive message via HTTP2 layer on any load balancer(nginx, or kubernetes).

## Requiretments
- Golang >= 1.8
- Dep as dependency manager

## How to build and run it

You can build source code via calling Makefile.
```sh
$ make build
    rm -f qtt
    go get -u github.com/golang/dep/cmd/dep
    dep ensure -vendor-only -v
    (1/2) Wrote github.com/joho/godotenv@v1.2.0
    (2/2) Wrote github.com/eclipse/paho.mqtt.golang@v1.1.1
go build -v  -o qtt
```
`Makefile` builds application. Artifact name is `qtt`. An than you can run like this:
```sh
$ ./qtt
2018/08/05 15:09:08 :1883 address is listening...
```

Makefile is a program building tool which runs on Unix, Linux, and their flavors. You dont know this tool, you can build offical golang tool instead of `Makefile`.

## Building via `go build`
```sh
$ go get -u github.com/golang/dep/cmd/dep
$ dep ensure -vendor-only -v
$ go build -v  -o myArtifact
```

# Configuration
Qtt app firstly look at a file has name `.env`. If the file is not exist. The app will get congigurations on the evironment. 

You can find out an example of `.env` in the project. The name is `.env.example`. There are a lot of default configuration in there.


## Todo:
- [X] Auth and Publish feature
- [X] MQTT 3.1.1 support
- [ ] Full WebHook support for above feature
- [ ] Stabilty for production..
- [ ] SubScription support
- [ ] Simple configuration managment
- [ ] Kubernetes support

# Thanks to
- `github.com/eclipse/paho.mqtt.golang`