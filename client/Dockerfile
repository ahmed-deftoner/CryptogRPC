FROM golang:1.19.5 as builder

WORKDIR /CryptogRPC

RUN mkdir client
RUN mkdir proto 

COPY ../proto proto
COPY client.go ./client

COPY ../go.mod .
COPY ../go.sum .

RUN cd proto && ls -al
#RUN go build -o client ./client

#CMD ./client