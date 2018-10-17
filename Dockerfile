FROM golang:1.8.3-alpine

WORKDIR /go/src/github.com/agungdwiprasetyo/agungkiki-backend

ENV SRC_DIR=/go/src/github.com/agungdwiprasetyo/agungkiki-backend

ENV BUILD_PACKAGES="git curl"

ADD . $SRC_DIR/

RUN apk update && apk add --no-cache $BUILD_PACKAGES \
  && apk add rsyslog \
  && apk add supervisor \
  && apk add tzdata \
  && curl https://glide.sh/get | sh \
  && glide install \
  && apk del $BUILD_PACKAGES \
  && CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o agungkiki-backend .

COPY .env $SRC_DIR/.env
ADD scripts/supervisord.conf /etc/supervisord.conf

EXPOSE 8008

ENTRYPOINT ["sh", "-c", "supervisord -nc /etc/supervisord.conf"]