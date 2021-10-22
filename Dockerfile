# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /build

COPY . ./

RUN go build .

WORKDIR /dist

RUN cp /build/main .

EXPOSE 3000

CMD [ "/dist/main" ]