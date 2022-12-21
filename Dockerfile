FROM golang:1.19.4-alpine3.17

RUN apk add --no-cache ca-certificates

ENV GOLANG_VERSION=1.19.4

WORKDIR /usr/src/app

COPY . .

RUN go mod download && go mod verify && go mod tidy

RUN go build -o /usr/local/bin/app .

EXPOSE 8080

CMD ["app"]