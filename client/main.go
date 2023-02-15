package main

import (
	"context"
	"errors"
	"github.com/bufbuild/connect-go"
	"log"
	"net/http"
	"pb/test"
	"pb/test/testconnect"
	"time"
)

func main() {
	ctx := context.Background()
	clientHttp := http.DefaultClient
	serviceClient := testconnect.NewHelloServiceClient(clientHttp, "http://localhost:8080")
	for {
		stream, err := serviceClient.HelloWord(ctx, connect.NewRequest(&test.HelloRequest{}))
		for stream.Receive() {
			msg := stream.Msg()
			log.Println("msg", msg)
		}
		//io.EOF
		err = stream.Err()
		if connectErr := new(connect.Error); errors.As(err, &connectErr) {
			switch connectErr.Code() {
			case connect.CodeUnavailable, connect.CodeInvalidArgument:
				time.Sleep(time.Second)
				log.Println("retry connect, prev err: ", connectErr.Code(), connectErr.Message())
				continue
			default:
				log.Println("disconnect", connectErr.Code(), connectErr.Message())
				return
			}
		}
	}
}
