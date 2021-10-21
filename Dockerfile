# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR
Learn more about the "WORKDIR" Dockerfile command.
 /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

EXPOSE 11111

CMD [ "/docker-gs-ping" ]