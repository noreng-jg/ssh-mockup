FROM golang:1.18-alpine3.15

RUN apk add ca-certificates git

WORKDIR /go/src/github.com/noreng-jg/sshserver

COPY ./server .

RUN go mod download

RUN go install github.com/markbates/refresh@v1.11.1

CMD ["refresh", "run"]
