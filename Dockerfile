FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/rarimo/auth-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/auth-svc /go/src/github.com/rarimo/auth-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/auth-svc /usr/local/bin/auth-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["auth-svc"]
