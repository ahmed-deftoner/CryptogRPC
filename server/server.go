package main

import (
	"context"
	"io"
	"log"
	"net"

	"github.com/adshao/go-binance/v2"
	pb "github.com/ahmed-deftoner/crypto-grpc/proto"
	"google.golang.org/grpc"
)

// this is the struct to be created, pb is imported upstairs
type BinanceServer struct {
	pb.BinanceServiceServer
}

func (s *BinanceServer) FetchAfterOneHour(stream pb.BinanceService_FetchAfterOneHourServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name : %v", req.Coin)
		client := binance.NewClient("", "")
		klines, err := client.NewKlinesService().Symbol(req.Coin).
			Interval("1h").Do(context.Background())
		if err != nil {
			return err
		}

		for _, v := range klines {

			resp := pb.Response{
				Price: v.High,
			}
			if err := stream.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
		}
	}
}

// define the port
const (
	port = ":8080"
)

func main() {
	//listen on the port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
	// create a new gRPC server
	grpcServer := grpc.NewServer()
	// register the greet service
	pb.RegisterBinanceServiceServer(grpcServer, &BinanceServer{})
	log.Printf("Server started at %v", lis.Addr())
	//list is the port, the grpc server needs to start there
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
