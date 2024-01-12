package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"storageReplication/node"
	"storageReplication/pb"
)

func main() {
	lis, err := net.Listen("tcp", ":43000")
	if err != nil {
		log.Fatal(err)
	}

	gs := grpc.NewServer()
	ns := node.NewService()

	pb.RegisterNodeServer(gs, ns)
	if err := gs.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
