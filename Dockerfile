FROM golang:1.22-bullseye as builder

WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y musl-tools

RUN ln -s /usr/bin/musl-gcc /usr/local/bin/musl-gcc

RUN apt-get clean && rm -rf /var/lib/apt/lists/*

RUN go mod tidy

RUN CC=/usr/local/bin/musl-gcc go build --ldflags '-linkmode external -extldflags "-static"' -tags musl -o /app-run /app/cmd/api/main.go

EXPOSE 8080

CMD ["/app-run"]
