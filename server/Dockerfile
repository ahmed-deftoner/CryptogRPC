FROM golang:1.19.5 as builder
RUN CGO_ENABLED=0 go get github.com/grpc-ecosystem/grpc-health-probe
RUN CGO_ENABLED=0 go get github.com/ahmed-deftoner/crypto-grpc/server

FROM scratch
COPY --from=builder /go/bin/grpc-health-probe /grpc-health-probe
COPY --from=builder /go/bin/server /server
EXPOSE 8080
CMD ["/server"]