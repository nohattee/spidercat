ARG VERSION=1.23.4
FROM golang:${VERSION}-alpine

RUN set -ex; \
    apk update && apk add --no-cache \
        build-base \ 
        bash \
        make \
        git \
        binutils-gold \
        openssh

# Install hot-reloading tool
RUN go install github.com/air-verse/air@v1.61.1

RUN mkdir -p -m 0700 ~/.ssh

WORKDIR /workspace
COPY . .