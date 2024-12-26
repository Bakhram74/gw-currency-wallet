package grpcClient

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New(port string) *grpc.ClientConn {

	conn, err := grpc.Dial("localhost"+fmt.Sprintf(":%s", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	return conn

}
