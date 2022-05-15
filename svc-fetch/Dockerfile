FROM golang:1.17-alpine

ENV GO111MODULE=on

ENV PKG_NAME=fetch-app
ENV PKG_PATH=$GOPATH/src/$PKG_NAME

WORKDIR $PKG_PATH/
COPY . $PKG_PATH/

RUN go mod vendor
WORKDIR $PKG_PATH

RUN go build app/main.go
CMD ["sh", "-c", "./main"]