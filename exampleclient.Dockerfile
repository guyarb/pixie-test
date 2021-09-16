FROM golang:1.16-alpine as builder

ADD . /build/
WORKDIR /build/example

RUN apk add gcc libc-dev
RUN GOOS=linux go build -a -mod=mod -ldflags '-extldflags "-static" -w' -o /build/client ./cmd/exampleclient/service.go

FROM alpine

COPY --from=builder /build/client /app/
WORKDIR /app
CMD ./client