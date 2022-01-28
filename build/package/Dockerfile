FROM golang:1.17 as builder

WORKDIR /go/src/app
COPY . .

RUN go build -o /build/bin/api ./cmd/api

FROM centos:7

COPY --from=builder /build/bin/api /build/bin/api

ENTRYPOINT ["/build/bin/api"]