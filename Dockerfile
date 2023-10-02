# Builder
FROM golang:1.18.3 AS builder

ARG GITHUB_PATH
ARG BRANCH

WORKDIR /go/src/
RUN git clone --branch $BRANCH $GITHUB_PATH
WORKDIR /go/src/quiz-fetcher
RUN make build

# fetcher

FROM golang:1.18.3 as server

COPY --from=builder /go/src/quiz-fetcher/fetcher /bin/
COPY --from=builder /go/src/quiz-fetcher/config.toml /etc/

CMD ["/bin/fetcher"]
