ARG GO_VERSION=1.14.1
FROM golang:${GO_VERSION}-alpine

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -o ./app

EXPOSE 3000
CMD ["./app"]


