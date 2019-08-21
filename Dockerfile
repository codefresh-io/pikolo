FROM alpine:3.8

RUN apk add --update ca-certificates

COPY dist/pikolo_linux_386/pikolo /usr/local/bin/
COPY VERSION /VERSION

ENTRYPOINT ["pikolo"]
CMD [ "--help" ]