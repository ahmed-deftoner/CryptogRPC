FROM golang:1.19.5 as builder
RUN CGO_ENABLED=0 go get github.com/ahmed-deftoner/crypto-grpc/client

FROM scratch
COPY --from=builder /go/bin/client /client
CMD ["/client"]