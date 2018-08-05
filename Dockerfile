# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

FROM debian:latest

COPY qtt /qtt

RUN apt-get update -y && \
    apt-get install -y ca-certificates \
    && update-ca-certificates --verbose \
    && apt-get clean

WORKDIR /

ENV MQTT_PORT=1883
ENV MQTT_SSL_PORT=4883
ENV DEBUG=TRUE
ENV AUTH_PROVIDER=default
ENV MQTT_DEFAULT_USER=user
ENV MQTT_DEFAULT_PASS=secret

CMD ["./qtt"]
