package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"storageReplication/node"
	"storageReplication/pb"
)

var (
	RedisHost     = os.Getenv("REDISHOST")
	RedisPassword = os.Getenv("REDISPASSWORD")
)

func main() {
	var nodeName string
	flag.StringVar(&nodeName, "nodename", "", "Set the node's name")
	flag.Parse()

	lis, err := net.Listen("tcp", ":43000")
	if err != nil {
		log.Fatal(err)
	}

	store := node.NewRedisKVStore(RedisHost, RedisPassword)

	gs := grpc.NewServer()
	ns := node.NewService(nodeName, "temp", store)

	pb.RegisterNodeServer(gs, ns)
	if err := gs.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
