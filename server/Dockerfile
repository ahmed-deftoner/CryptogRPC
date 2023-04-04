FROM golang:1.19.5 as builder

WORKDIR /CryptogRPC

RUN mkdir proto 

COPY proto ./proto
COPY server/server.go .

COPY go.mod .
COPY go.sum .

RUN go mod download

CMD ["go run server.go"]