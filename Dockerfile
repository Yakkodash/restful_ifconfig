FROM golang

ADD . /go/

ENV TARGET server

RUN go get -d -v $TARGET
RUN go install $TARGET
RUN go test $TARGET

ENTRYPOINT /go/bin/$TARGET

EXPOSE 8080
