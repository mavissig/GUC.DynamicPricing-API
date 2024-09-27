FROM golang:1.22-bullseye as builder

WORKDIR /app

COPY . .

RUN sh /app/docker/lib.sh

CMD ["/app-run"]
