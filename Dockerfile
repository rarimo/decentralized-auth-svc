FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/rarimo/rarime-auth-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/rarime-auth-svc /go/src/github.com/rarimo/rarime-auth-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/rarime-auth-svc /usr/local/bin/rarime-auth-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["rarime-auth-svc"]
