package main

import (
	"bufio"
	"context"
	"fmt"
	proto "github.com/provider-go/pkg/grpc/pb"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	var serviceHost = "127.0.0.1:8001"

	conn, err := grpc.Dial(serviceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client := proto.NewHelloClient(conn)
	rsp, err := client.Say(context.TODO(), &proto.SayRequest{
		Name: "biwow" +
			"",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("接到返回:", rsp)

	fmt.Println("按回车键退出程序...")
	in := bufio.NewReader(os.Stdin)
	_, _, _ = in.ReadLine()
}
