package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ahmed-deftoner/crypto-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func callServer(client pb.BinanceServiceClient, payload *pb.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	res, err := client.FetchAfterOneHour(ctx, &pb.Request{Coin: payload.Coin, Interval: payload.Interval})
	if err != nil {
		log.Fatalf("Could not get data: %v", err)
	}
	for i, v := range res.Price {
		log.Printf("%d %s", i, v)
	}
}

func main() {
	conn, err := grpc.Dial("server-service"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBinanceServiceClient(conn)

	payload := &pb.Request{
		Coin:     "BTCUSDT",
		Interval: "1h",
	}
	callServer(client, payload)
}
