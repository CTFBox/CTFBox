## for build
FROM golang:1.15-alpine as builder

WORKDIR /github.com/CTFBox/CTFBox
COPY go.mod go.sum ./
RUN go mod download

COPY main.go db.go ./
COPY ./repository ./repository
COPY ./router ./router
COPY ./model ./model

RUN go build -o /ctf_box -ldflags "-s -w"

## for run
FROM alpine:3.12
ENV TZ Asia/Tokyo

RUN apk --update --no-cache add tzdata \
  && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && apk del tzdata
RUN apk --update --no-cache add ca-certificates \
  && update-ca-certificates \
  && rm -rf /usr/share/ca-certificates /etc/ssl/certs

WORKDIR /app
COPY --from=builder /ctf_box ./

ENTRYPOINT ./ctf_box

## for development
FROM golang:1.15-alpine as development

WORKDIR /github.com/CTFBox/CTFBox
COPY go.mod go.sum ./
RUN go mod download

COPY main.go db.go ./
COPY ./repository ./repository
COPY ./router ./router
COPY ./model ./model

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

RUN wget https://github.com/go-task/task/releases/download/v3.0.0/task_linux_386.tar.gz \
    && tar -C /usr/local/bin -xvf task_linux_386.tar.gz task \
    && rm task_linux_386.tar.gz
COPY Taskfile.yml ./

RUN go build -o /app/ctf_box -ldflags "-s -w"

EXPOSE 8080
