FROM golang:latest as builder

COPY . /go/src/
WORKDIR /go/src
RUN ./scripts/build.sh

FROM alpine:latest
WORKDIR /root
COPY --from=builder /go/src/assignment .
CMD ./assignment



