# Builder
ARG BASE_IMAGE=golang:1.18.10-alpine3.17
FROM ${BASE_IMAGE} AS builder

COPY ./ /go/src/quiz-fetcher

WORKDIR /go/src/quiz-fetcher
RUN go build -buildvcs=auto -o fetcher ./cmd/fetcher

# fetcher

FROM alpine:3.17 as server

COPY --from=builder /go/src/quiz-fetcher/fetcher /bin/
COPY --from=builder /go/src/quiz-fetcher/config.yaml /etc/

CMD ["/bin/fetcher"]
