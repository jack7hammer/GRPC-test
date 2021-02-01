package main

import (
	"context"
	
	"C:/Users/jackh/Desktop/g1/github.com/monkrus/grpc-from0"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp","404")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterBMIServer(srv,&server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e!=nil{
		panic(err)
	}

}

func (s *server) getBMI(ctx context.Context, in *protos.wahRequest) (*protos.bmiResponse, error) {
	h,w := wahRequest.GetHeight() , wahRequest.GetWeight();
	result := w/(h*h)
	return &protos.bmiResponse{bmi: result}, nil
}
