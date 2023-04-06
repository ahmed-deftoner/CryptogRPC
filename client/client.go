package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/ahmed-deftoner/crypto-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func callBinanceOneHour(client pb.BinanceServiceClient, coin string) {
	log.Printf("Bidirectional Streaming started")
	stream, err := client.FetchAfterOneHour(context.Background())
	if err != nil {
		log.Fatalf("Could not send coin: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	req := &pb.Request{
		Coin: coin,
	}
	if err := stream.Send(req); err != nil {
		log.Fatalf("Error while sending %v", err)
	}
	time.Sleep(30 * time.Second)

	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional Streaming finished")
}

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBinanceServiceClient(conn)

	coin := "BTCUSDT"
	for {
		callBinanceOneHour(client, coin)
		time.Sleep(5 * time.Second)
	}
}
