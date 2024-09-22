FROM golang:1.22
LABEL authors="admin"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /app-run /app/cmd/api/main.go

CMD ["/app-run"]