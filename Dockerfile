FROM golang:1.16.4-alpine3.13

ENV GO111MODULE=on


WORKDIR /app
COPY . .

RUN set -ex && \
    go mod download && \
    go mod verify

EXPOSE 50051

CMD [ "go","run","/app/server/user_server.go" ]