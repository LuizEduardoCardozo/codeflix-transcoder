FROM golang:1.18beta1-alpine3.15

WORKDIR /go/src

RUN apk add --update ffmpeg bash curl cmake make
RUN apk add --update --upgrade curl python2 unzip bash gcc g++ scons git

COPY ./scripts/install-bento4.sh ./scripts/install-bento4.sh
RUN sh ./scripts/install-bento4.sh 

ENTRYPOINT [ "top" ]
