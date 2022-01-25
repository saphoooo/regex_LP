FROM golang:1.17

WORKDIR /

COPY . .

RUN go get -d -v .
RUN GOOS=linux GOARCH=amd64 go build -o wordfindv1

FROM ubuntu:20.04

LABEL maintainer="stephane.beuret@gmail.com"

COPY --from=0 wordfindv1 /

ENTRYPOINT ["/wordfindv1"]