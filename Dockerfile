FROM golang:1.11.11-alpine3.8 as build

WORKDIR /pikolo

RUN apk add git gcc g++

COPY go.mod .
RUN go mod download

COPY . .
RUN env CGO_ENABLED=0 go build -ldflags="-s -w"

FROM alpine:3.8

RUN apk add --update ca-certificates

COPY --from=build /pikolo/pikolo /usr/local/bin
COPY VERSION /VERSION

LABEL io.codefresh.engine="true"

ENTRYPOINT ["pikolo"]
CMD [ "--help" ]