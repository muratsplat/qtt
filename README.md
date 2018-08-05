[![Build Status](https://travis-ci.com/muratsplat/qtt.svg?branch=master)](https://travis-ci.com/muratsplat/qtt)
# Qtt is simple message broker on based webhook


This project is not completed yet. Don't use it on your production until further notice..

This is experiment about handling MQTT mesage from clients. Main goal is all message forwarded by HTTP2 to another point. It is easy to deliver massive message via HTTP2 layer on any load balancer(nginx, or kubernetes).

## Requiretments
- Golang >= 1.8
- Dep as dependency manager


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