FROM golang AS builder

WORKDIR /go/src/totp-proxy
COPY . .

RUN set -ex \
    && GO111MODULE=on go build -o /go/bin/totp-proxy


FROM debian:stretch-slim

COPY --from=builder /go/bin/totp-proxy /

EXPOSE 8080

ENTRYPOINT ["/totp-proxy"]
CMD ["help"]