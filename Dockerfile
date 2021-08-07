FROM golang:1.16 as builder
WORKDIR /go/src/github.com/openmock/openmock
ADD . .
RUN make build

FROM quay.io/goswagger/swagger as swagger-builder

FROM alpine:3.6
WORKDIR /go/src/github.com/openmock/openmock
RUN apk add --no-cache ca-certificates libc6-compat
COPY --from=builder /go/src/github.com/openmock/openmock/om ./om
COPY --from=builder /go/src/github.com/openmock/openmock/omctl ./omctl
COPY --from=builder /go/src/github.com/openmock/openmock/docs/api_docs/bundle.yaml ./swagger-docs.yaml
COPY --from=swagger-builder /usr/bin/swagger ./swagger
ENV OPENMOCK_HTTP_HOST=0.0.0.0
ENV OPENMOCK_TEMPLATES_DIR=/data/templates
CMD ./om
