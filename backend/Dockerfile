FROM golang:alpine

RUN apk add --no-cache curl
RUN apk --update add git openssh && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

RUN curl https://glide.sh/get | sh

ENV PKG_NAME=github.com/Jim-Lin/bee-bee-alert/backend
ENV PKG_PATH=$GOPATH/src/$PKG_NAME
WORKDIR $PKG_PATH

COPY glide.yaml glide.lock $PKG_PATH/
RUN glide install

COPY . $PKG_PATH
RUN go build -o main .

CMD ["sh", "-c", "${PKG_PATH}/main"]
