package main

import (
	"context"
	"fmt"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/guyarb/pixie-test/internal/proto"
	"google.golang.org/grpc"
)

var args struct {
	ServerAddress string `arg:"required,positional"`
}

func main() {
	arg.MustParse(&args)

	for {
		time.Sleep(time.Second * 20)
		conn, err := grpc.DialContext(context.Background(), args.ServerAddress, grpc.WithBlock(), grpc.WithInsecure())
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()
		grpcClient := proto.NewEchoClient(conn)
		request := proto.SampleRequest{Msg: "test at " + time.Now().String()}
		resp, err := grpcClient.Sample(context.Background(), &request)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("got %s\n", resp.GetMsg())
	}
}
