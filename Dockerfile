# build stage
FROM golang:1.20-rc-buster

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o boris_backend

RUN apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates

RUN update-ca-certificates

EXPOSE 4000 4000

ENTRYPOINT ./boris_backend