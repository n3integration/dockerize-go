FROM alpine:3.5
MAINTAINER <n3integration@gmail.com>

EXPOSE 8080

VOLUME /data

COPY dockerize-go /app/dockerize-go

CMD ["/app/dockerize-go"]
