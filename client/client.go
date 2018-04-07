package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"

	pb "github.com/ynishi/grpc-tasks"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	id := flag.Arg(0)
	id64, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("arg err", err)
	}
	idi := int32(id64)

	conn, err := grpc.Dial(":1315", grpc.WithInsecure())
	if err != nil {
		log.Fatal("conn err", err)
	}
	defer conn.Close()

	client := pb.NewTaskServiceClient(conn)
	message := &pb.Task{
		Name:    "abc",
		Id:      idi,
		Title:   "title1",
		Tags:    nil,
		DueTime: nil,
		Added:   nil,
		Updated: nil,
	}

	res, err := client.AddTask(context.TODO(), message)
	fmt.Printf("result:%#v %#v\n", res, res.Done)
	fmt.Printf("err:%#v \n", err)

	req := &pb.TaskRequest{idi, ""}
	res2, err := client.GetTask(context.TODO(), req)
	fmt.Printf("result:%#v \n", res2)
	fmt.Printf("err:%#v \n", err)

}
