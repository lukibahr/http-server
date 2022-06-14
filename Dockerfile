# syntax=docker/dockerfile:1
FROM scratch
COPY http-server /usr/bin/http-server
USER daemon
ENTRYPOINT ["/usr/bin/http-server"]
