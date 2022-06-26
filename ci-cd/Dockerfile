FROM docker
ARG GOLANG_VERSION=1.18
MAINTAINER "deema_k@mail.ru"
LABEL version="1.0"
ENV DOCKER_DRIVER overlay
EXPOSE 2375 2376
RUN apk update && apk add go gcc bash musl-dev openssl-dev ca-certificates jq git buildah zip && update-ca-certificates
RUN apk add --no-cache --virtual .build-deps bash gcc musl-dev openssl go
RUN wget -O go.tgz https://dl.google.com/go/go$GOLANG_VERSION.src.tar.gz
RUN tar -C /usr/local -xzf go.tgz
RUN cd /usr/local/go/src && ./make.bash
ENV PATH="/usr/local/go/bin:$PATH"
ENV GOPATH=/opt/go/
ENV PATH=$PATH:$GOPATH/bin
RUN apk del .build-deps
RUN echo | go version
RUN echo | docker --version
