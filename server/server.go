package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/adshao/go-binance/v2"
	pb "github.com/ahmed-deftoner/crypto-grpc/proto"
	"google.golang.org/grpc"
)

// define the port
const (
	port = ":8080"
)

// this is the struct to be created, pb is imported upstairs
type binanceServer struct {
	pb.BinanceServiceServer
}

func (b *binanceServer) FetchAfterOneHour(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	client := binance.NewClient("", "")
	klines, err := client.NewKlinesService().Symbol("BTCUSDT").
		Interval("1h").Do(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var arr []string
	for _, k := range klines {
		arr = append(arr, fmt.Sprintf("%v", k.High))
	}
	return &pb.Response{
		Price: arr,
	}, nil
}

func main() {
	//listen on the port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
	// create a new gRPC server
	grpcServer := grpc.NewServer()
	// register the greet service
	pb.RegisterBinanceServiceServer(grpcServer, &binanceServer{})
	log.Printf("Server started at %v", lis.Addr())
	//list is the port, the grpc server needs to start there
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
