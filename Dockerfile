FROM golang

ADD . /go/

ENV TARGET restful_ifconfig

RUN go get -d -v $TARGET
RUN go install $TARGET
RUN go test -v $TARGET

ENTRYPOINT /go/bin/$TARGET

EXPOSE 8080
