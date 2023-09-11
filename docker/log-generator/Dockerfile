FROM alpine:latest

RUN mkdir -p /tmp/logs/

RUN touch /tmp/logs/log-level

COPY ./entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

WORKDIR /tmp/logs/
