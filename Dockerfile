FROM golang:1.24-bookworm AS build
WORKDIR /pikolo

RUN apt-get update -y \
    && apt-get install -y git gcc g++

COPY go.mod .
RUN go mod download

COPY . .
RUN env CGO_ENABLED=0 go build -ldflags="-s -w"

FROM debian:bookworm-20250407-slim
RUN adduser --gecos "" --disabled-password --home /home/cfu --shell /bin/bash cfu
USER cfu

COPY --from=build --chown=cfu:cfu /pikolo/pikolo /usr/local/bin
COPY --chown=cfu:cfu VERSION /VERSION

ENTRYPOINT ["pikolo"]
CMD [ "--help" ]
