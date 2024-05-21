# syntax=docker/dockerfile:1
FROM gcr.io/distroless/base

LABEL maintainer="hello@lukasbahr.de"

COPY http-server /usr/bin/http-server
ENTRYPOINT ["/usr/bin/http-server"]
