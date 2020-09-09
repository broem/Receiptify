FROM golang:1.14-alpine
ENV GO111MODULE=auto

RUN apk add --no-cache \
    git \
    build-base \
    leptonica-dev \
    tesseract-ocr

RUN go get

ENV PORT=8086
CMD $GOPATH/bin/Receiptify
