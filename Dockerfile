FROM golang:alpine AS builder

WORKDIR /usr/src

ADD go.mod .
ADD go.sum .

COPY . .

RUN go build -o gameplatform ./cmd/web/

CMD ["/usr/src/gameplatform"]