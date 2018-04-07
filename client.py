import grpc

import tasks_pb2
import tasks_pb2_grpc

chan = grpc.insecure_channel(':1315')
stub = tasks_pb2_grpc.TaskServiceStub(chan)

task = stub.GetTask(tasks_pb2.TaskRequest(id=1,tag=""))

print(task)
