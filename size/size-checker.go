package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	emailServer "grpc-sample/data"
	"io"
	"io/ioutil"
	"log"
)

type CounterWr struct {
	io.Writer
	Count int
}

func (cw *CounterWr) Write(p []byte) (n int, err error) {
	n, err = cw.Writer.Write(p)
	cw.Count += n
	return
}

func main() {

	buf := &bytes.Buffer{}
	var out io.Writer = buf
	cw := &CounterWr{Writer: out}

	s := `"email" :"mrigank.anand52@gmail.com",
			"subject": "Welcome to our application",
			"body":  "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."}`
	if err := json.NewEncoder(cw).Encode(s); err != nil {
		panic(err)
	}

	fmt.Printf("Count: %d bytes, %d bits\n", cw.Count, cw.Count*8)

	data := &emailServer.EmailData{
		Email:   "mrigank.anand52@gmail.com",
		Subject: "Welcome to our application",
		Body:    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
	}

	output, err := proto.Marshal(data)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	fname := "encodedData"
	if err := ioutil.WriteFile(fname, output, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

}
