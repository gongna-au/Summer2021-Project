package commands

import (
	"context"
	"fmt"
	"github.com/Summer2021-Project/dotenv"
	pb "github.com/Summer2021-Project/grpc-skeleton/protos"
	"google.golang.org/grpc"
	"time"
)

type GrpcClientCommand struct {
}

func (t *GrpcClientCommand) Main() {
	addr := dotenv.Getenv("GIN_ADDR").String(":8080")
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = conn.Close()
	}()
	cli := pb.NewUserClient(conn)
	req := pb.AddRequest{
		Name: "Dubbo-go",
	}
	resp, err := cli.Add(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Add User: %d", resp.UserId))
}
