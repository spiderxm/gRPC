package main

import (
	"context"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	emailServer "grpc-sample/data"
	"grpc-sample/helpers"
	"io"
	"log"
	"net"
)

var port = ":8000"

type EmailServer struct {
	emailServer.UnsafeEmailServiceServer
}

func (s *EmailServer) SendMail(ctx context.Context, data *emailServer.EmailData) (*emailServer.Message, error) {
	// Send Email
	return helpers.SendEmail(data), nil
}

func (s *EmailServer) SendBulkEmail(stream emailServer.EmailService_SendBulkEmailServer) error {
	for {
		emailData, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&emailServer.Message{
				ResponseText: "Emails sent successfully",
				Status:       1,
			})
		}
		if err != nil {
			return stream.SendAndClose(&emailServer.Message{
				ResponseText: "There is some error please try again later",
				Status:       -1,
			})
		}
		helpers.SendEmail(emailData)
	}
}

func main() {
	e := godotenv.Load(".env")
	if e != nil {
		log.Fatalln("Could not start gRPC-server not able to load dependencies")
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	emailServer.RegisterEmailServiceServer(s, &EmailServer{})
	log.Printf("gRPC-server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
