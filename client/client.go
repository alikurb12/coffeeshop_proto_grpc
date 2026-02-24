package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/alikurb12/proto_example/coffeeshop_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("eroor initializyng connection: %s", err)
	}

	defer conn.Close()

	c := pb.NewCoffeeShopClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	menuStream, err := c.GetMenu(ctx, &pb.MenuRequest{})
	if err != nil {
		log.Fatalf("Error while streaming menu: %s", err)
	}
	done := make(chan bool)
	var items []*pb.Item
	go func() {
		for {
			resp, err := menuStream.Recv()
			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				log.Fatalf("Cannot recieve %v", err)
			}

			items = resp.Items
			log.Printf("Recp recieved: %v", resp.Items)
		}
	}()

	<-done
	receipt, err := c.PlaceOrder(ctx, &pb.Order{Items: items})
	log.Printf("%v", receipt)

	status, err := c.GetOrderStatus(ctx, receipt)
	log.Printf("%v", status)
}
