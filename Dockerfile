FROM golang:1.18beta1-alpine3.15

RUN apk add --update ffmpeg bash curl cmake make

WORKDIR /tmp/bento4

RUN apk add --update --upgrade curl python2 unzip bash gcc g++ scons git

RUN git clone https://github.com/axiomatic-systems/Bento4.git && cd Bento4 && ls -a && \
    mkdir cmakebuild && cd cmakebuild && \
    cmake -DCMAKE_BUILD_TYPE=Release .. && \
    make

WORKDIR /go/src

ENTRYPOINT [ "top" ]
