FROM alpine:3.8

RUN apk add --update ca-certificates

COPY dist/pikolo_linux_386/pikolo /usr/local/bin/
COPY VERSION /VERSION

LABEL io.codefresh.engine="true"

ENTRYPOINT ["pikolo"]
CMD [ "--help" ]