package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	emailServer "grpc-sample/data"
)

const (
	address = "localhost:8000"
)

func main() {
	// Set up a connection to the gRPC-server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := emailServer.NewEmailServiceClient(conn)

	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c.SendMail(ctx, &emailServer.EmailData{
		Email:   "mrigank.anand52@gmail.com",
		Subject: "Welcome to our application",
		Body:    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
	})
	if err != nil {
		log.Fatalf("Could not send request : %v", err)
	}
	log.Printf("Response Text : %s", r.ResponseText)
	log.Printf("Response Status: %v", r.Status)

	elapsed := time.Since(start)
	log.Printf("Sending Email using gPRC took %s", elapsed)

	emailsData := func() []*emailServer.EmailData {
		var data []*emailServer.EmailData
		for i := 0; i < 2; i++ {
			data = append(data, &emailServer.EmailData{
				Email:   "mrigank.anand52@gmail.com",
				Subject: "Welcome to our application",
				Body:    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
			})
		}
		return data
	}()

	stream, err := c.SendBulkEmail(context.Background())
	if err != nil {
		log.Fatalf("Error on sending bulk email: %v", err)
	}
	for _, v := range emailsData {
		fmt.Println("Sending Request: \n", v)
		stream.Send(v)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error when closing the stream and receiving the response: %v", err)
	}
	fmt.Println(res)
}
