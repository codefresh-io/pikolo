FROM golang:1.17.8-bullseye as build

WORKDIR /pikolo

RUN apt-get update -y \
    && apt-get install -y git gcc g++

COPY go.mod .
RUN go mod download

COPY . .
RUN env CGO_ENABLED=0 go build -ldflags="-s -w"

FROM debian:11.6-slim

RUN apt-get update -y \
    && apt-get install -y ca-certificates busybox \
    && ln -s /bin/busybox /usr/bin/[[

COPY --from=build /pikolo/pikolo /usr/local/bin
COPY VERSION /VERSION

LABEL io.codefresh.engine="true"

RUN adduser --gecos "" --disabled-password --home /home/cfu --shell /bin/bash cfu

USER cfu

ENTRYPOINT ["pikolo"]
CMD [ "--help" ]
