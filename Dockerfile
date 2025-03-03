FROM golang:1.10 as go-builder

ARG LEANOTE_VERSION=2.6.1
COPY . /go/src/github.com/leanote/leanote

RUN echo build leanote ${LEANOTE_VERSION} from source \
    && cd /go/src/github.com/leanote/leanote \
    && go run app/cmd/main.go \
    && go build -o bin/leanote-linux-amd64 github.com/leanote/leanote/app/tmp

FROM node:9-alpine as node-builder

COPY --from=go-builder /go/src/github.com/leanote/leanote /go/src/github.com/leanote/leanote

RUN cd /go/src/github.com/leanote/leanote \
    && npm config set unsafe-perm true \
    && npm install \
    && npm install -g gulp \
    && npm install gulp-minify-css \
    && gulp

FROM debian:stable-slim

COPY --from=go-builder /go/src/github.com/leanote/leanote/bin/leanote-linux-amd64 /leanote/bin/
COPY --from=go-builder /go/src/github.com/leanote/leanote/bin/run-linux-amd64.sh /leanote/bin/run.sh
COPY --from=go-builder /go/src/github.com/leanote/leanote/bin/src/ /leanote/bin/src/

COPY --from=node-builder /go/src/github.com/leanote/leanote/app/views /leanote/app/views
COPY --from=node-builder /go/src/github.com/leanote/leanote/conf /leanote/conf
COPY --from=node-builder /go/src/github.com/leanote/leanote/messages /leanote/messages
COPY --from=node-builder /go/src/github.com/leanote/leanote/mongodb_backup /leanote/mongodb_backup
COPY --from=node-builder /go/src/github.com/leanote/leanote/public /leanote/public

RUN apt-get update \
    && apt-get install -y wget ca-certificates \
    && mkdir -p /leanote/data/public/upload \
    && mkdir -p /leanote/data/files \
    && mkdir -p /leanote/data/mongodb_backup \
    # && rm -r /leanote/public/upload \
    # && rm -r /leanote/mongodb_backup \
    && ln -s /leanote/data/public/upload /leanote/public/upload \
    && ln -s /leanote/data/files /leanote/files \
    && ln -s /leanote/data/mongodb_backup /leanote/mongodb_backup \
    && apt-get clean && rm -rf /var/lib/apt/lists/*


VOLUME /leanote/data/

EXPOSE 9000

CMD ["sh", "/leanote/bin/run.sh"]
