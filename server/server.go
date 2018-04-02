package main

import (
	"context"
	"net"

	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/ynishi/tasks"
	"google.golang.org/grpc"
)

type tasksServer struct {
}

func (s *tasksServer) AddTask(ctx context.Context, in *pb.Task) (*pb.TaskResponse, error) {
	return &pb.TaskResponse{0, &google_protobuf.Timestamp{}}, nil
}

func (s *tasksServer) GetTask(ctx context.Context, taskRequest *pb.TaskRequest) (*pb.Task, error) {
	return &pb.Task{}, nil
}

func (s *tasksServer) DoneTask(ctx context.Context, task *pb.Task) (*pb.TaskResponse, error) {
	return &pb.TaskResponse{}, nil
}

func (s *tasksServer) ListTask(taskRequest *pb.TaskRequest, lts pb.TaskService_ListTaskServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":1315")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	ts := tasksServer{}
	pb.RegisterTaskServiceServer(grpcServer, &ts)
	grpcServer.Serve(lis)
}
