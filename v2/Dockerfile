FROM golang:1.17

WORKDIR /

COPY . .

RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -o wordfindv2

FROM ubuntu:20.04

LABEL maintainer="stephane.beuret@gmail.com"

COPY --from=0 wordfindv2 /

ENTRYPOINT ["/wordfindv2"]