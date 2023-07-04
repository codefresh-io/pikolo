FROM golang:1.19.9-bullseye as build

WORKDIR /pikolo

RUN apt-get update -y \
    && apt-get install -y git gcc g++

COPY go.mod .
RUN go mod download

COPY . .
RUN env CGO_ENABLED=0 go build -ldflags="-s -w"

FROM debian:11-slim

# Update package lists and upgrade existing packages
RUN apt-get update && apt-get upgrade -y

# Install packages required for the base image
RUN apt-get install -y ca-certificates busybox \
    && ln -s /bin/busybox /usr/bin/[[

# Clean up the package cache to reduce image size
RUN apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=build /pikolo/pikolo /usr/local/bin
COPY VERSION /VERSION

LABEL io.codefresh.engine="true"

RUN adduser --gecos "" --disabled-password --home /home/cfu --shell /bin/bash cfu

USER cfu

ENTRYPOINT ["pikolo"]
CMD [ "--help" ]
