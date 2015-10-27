FROM golang:latest

RUN echo 'deb http://httpredir.debian.org/debian jessie non-free' > /etc/apt/sources.list.d/debian-non-free.list \
    && export DEBIAN_FRONTEND=noninteractive \
    && apt-get update \
    && apt-get -y upgrade \
    && apt-get -y dist-upgrade \
    && apt-get install nodejs npm -y \
    && npm install bower

RUN go get github.com/gin-gonic/gin

RUN mkdir -p /srv/tools/moldova-meta-search
WORKDIR /srv/tools/moldova-meta-search

ADD /server /server
ADD /config /config
ADD /client /client
ADD /ansible /ansible

CMD go run /server/main.go

EXPOSE 3000
