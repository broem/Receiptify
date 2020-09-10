FROM golang:1.14-alpine
ENV GO111MODULE=on

RUN apk add --no-cache \
    git \
    build-base \
    leptonica-dev \
    tesseract-ocr-dev \
    tesseract-ocr

ADD . /go/src/Receiptify
WORKDIR /go/src/Receiptify
COPY ./server /go/src/Receiptify/server
RUN go get ./...
RUN go build

CMD Receiptify

EXPOSE 8086