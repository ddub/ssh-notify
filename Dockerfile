# Docker builder for Golang
FROM golang as builder

COPY . .
RUN set -x && \ 
    cd install && \
    go get -d -v . && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
RUN set -x && \
    cd notify && \
    go get -d -v . && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# Docker run Golang app
FROM scratch
LABEL maintainer "Duncan Ward <docker@dub.to>"

ENV SSHNOTIFY_SOURCE=/notify SSHNOTIFY_LOOP=60
COPY --from=builder /go/install/install /
COPY --from=builder /go/notify/notify /
CMD ["/install"]
