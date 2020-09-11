FROM golang:alpine as builder

RUN apk add --no-cache make git
WORKDIR /proxyaggre-src
COPY . /proxyaggre-src
RUN go mod download && \
    make docker && \
    mv ./bin/proxyaggre-docker /proxyaggre

FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata
WORKDIR /proxyaggre-src
COPY ./assets /proxyaggre-src/assets
COPY --from=builder /proxyaggre /proxyaggre-src/
ENTRYPOINT ["/proxyaggre-src/proxyaggre"]
