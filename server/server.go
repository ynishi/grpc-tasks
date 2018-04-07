package main

import (
	"context"
	"net"

	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/ynishi/grpc-tasks"
	"google.golang.org/grpc"
)

type tasksServer struct {
	Tasks []*pb.Task
}

func (s *tasksServer) AddTask(ctx context.Context, in *pb.Task) (*pb.TaskResponse, error) {
	s.Tasks = append(s.Tasks, in)
	return &pb.TaskResponse{in.Id, &google_protobuf.Timestamp{}}, nil
}

func (s *tasksServer) GetTask(ctx context.Context, taskRequest *pb.TaskRequest) (*pb.Task, error) {
	for _, task := range s.Tasks {
		if task.Id == taskRequest.Id {
			return task, nil
		}
	}
	ret := s.Tasks[len(s.Tasks)-1]
	return ret, nil
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
