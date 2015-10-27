FROM golang:latest
MAINTAINER Aleksandar Todorović <aleksandar@occrp.org>

RUN export DEBIAN_FRONTEND=noninteractive \
    && apt-get update \
    && apt-get -y upgrade \
    && apt-get -y dist-upgrade \
    && apt-get install nodejs npm -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

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
