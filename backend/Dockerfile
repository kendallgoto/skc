FROM golang:1.18-alpine as builder

RUN apk add --no-cache make git gcc g++ pkgconfig openjdk11

# download correct python version
RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.10/main' >> /etc/apk/repositories
RUN apk update && apk --no-cache add python3-dev=3.7.10-r0

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN make

FROM python:3.7.16-alpine

WORKDIR /app

COPY --from=builder /app/bin/skcserver /app/skcserver

ENTRYPOINT ["/app/skcserver"]
