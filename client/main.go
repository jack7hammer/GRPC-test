package main

import (
	"google.golang.org/grpc"
	"C:\Users\jackh\Desktop\g1\github.com\monkrus\grpc-from0"
	"github.com/gin-gonic/gin"
)
func main(){
	conn, err := grpc.Dial("Localhost:9000",grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewAddServiceClient(conn)
	g := gin.Default()
	
}