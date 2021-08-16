FROM golang:1.16 as builder

WORKDIR /build

ENV GO111MODULE="on"
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# syntax=docker/dockerfile:1
FROM scratch
WORKDIR /app
COPY --from=builder /build/server .
CMD ["./server"]
