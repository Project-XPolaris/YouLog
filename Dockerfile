ARG GOLANG_VERSION=1.17
FROM golang:${GOLANG_VERSION}-buster as builder
ARG GOPROXY=https://goproxy.cn
WORKDIR ${GOPATH}/src/github.com/projectxpolaris/youlog

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ${GOPATH}/bin/youlog ./youlog.go

FROM debian:buster-slim

COPY --from=builder /usr/local/lib /usr/local/lib
COPY --from=builder /etc/ssl/certs /etc/ssl/certs


COPY --from=builder /go/bin/youlog /usr/local/bin/youlog



ENTRYPOINT ["/usr/local/bin/youlog","run"]

