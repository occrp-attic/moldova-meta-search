FROM golang:latest
MAINTAINER aleksandar@occrp.org

RUN echo 'deb http://httpredir.debian.org/debian jessie non-free' > /etc/apt/sources.list.d/debian-non-free.list \
    && export DEBIAN_FRONTEND=noninteractive \
    && apt-get update \
    && apt-get -y upgrade \
    && apt-get -y dist-upgrade \
    && apt-get clean \
    && apt-get install nodejs npm -y

RUN npm install bower

RUN go get github.com/gin-gonic/gin

RUN mkdir -p /srv/tools/moldova-meta-search
WORKDIR /srv/tools/moldova-meta-search

COPY /server /server
COPY /config /config
COPY /client /client
COPY /ansible /ansible

CMD ["go","run","/server/main.go"]

EXPOSE 3000
