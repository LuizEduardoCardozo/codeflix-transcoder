FROM alfg/bento4:latest

RUN apk add wget go

WORKDIR /go/src

ENTRYPOINT ["top"]
