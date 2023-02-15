package main

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"pb/test"
	"pb/test/testconnect"
	"time"
)

type HelloWordServer struct{}

func (hello *HelloWordServer) HelloWord(ctx context.Context, c *connect.Request[test.HelloRequest], stream *connect.ServerStream[test.HelloResponse]) error {
	//TODO implement me
	log.Println("Subscribe ", stream.ResponseHeader())
	//var saying strings.Builder
	for {
		g := fmt.Sprintf("Hello, %s\n", time.Now().String())
		err := stream.Send(&test.HelloResponse{
			Say: g,
		})
		fmt.Printf("send %s", g)

		if err != nil {
			log.Println("err send client", err)
			return err
		}
		time.Sleep(time.Second * 5)

	}

}

func main() {
	sayServer := &HelloWordServer{}
	mux := http.NewServeMux()
	path, handler := testconnect.NewHelloServiceHandler(sayServer)
	mux.Handle(path, handler)

	err := http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		log.Println("error start serve", err)
		return
	}
}
