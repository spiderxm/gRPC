package gprc_sample

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	emailServer "grpc-sample/data"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

const (
	address = "localhost:8000"
)

func Benchmark_gPRCCall(b *testing.B) {
	b.N = 10
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := emailServer.NewEmailServiceClient(conn)
	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		c.SendMail(ctx, &emailServer.EmailData{
			Email:   "195516@nith.ac.in",
			Subject: "Welcome to our application",
			Body:    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		})
		cancel()
	}
}

func Benchmark_RESTCall(b *testing.B) {
	b.N = 10
	for i := 0; i < b.N; i++ {
		url := "http://localhost:3000/api/sendEmail"
		method := "POST"

		payload := strings.NewReader(`{
    "email" :"195516@nith.ac.in",
    "subject": "Welcome to our application",
    "body":  "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."}`)

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))
	}
}
