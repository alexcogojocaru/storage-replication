package node

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"os"
	"path/filepath"
	"storageReplication/pb"
)

type Service struct {
	pb.UnimplementedNodeServer

	NodeName         string
	storageDirectory string
	kvstore          KVStore
}

func (s Service) Receive(ctx context.Context, data *pb.Data) (*pb.Response, error) {
	id := uuid.New().String()
	fileStoragePath := filepath.Join(s.storageDirectory, id)

	err := os.WriteFile(fileStoragePath, data.Chunk, os.ModePerm)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%s-%s", s.NodeName, id)

	content, err := proto.Marshal(data.Metadata)
	if err != nil {
		log.Printf("[WARNING] cannot marshal metadata: %v\n", err)
	} else {
		if err := s.kvstore.Set(key, content); err != nil {
			log.Printf("[WARNING] cannot set redis key: %v\n", err)
		}
	}

	return &pb.Response{
		Status: true,
	}, nil
}

func (s Service) Delete(ctx context.Context, metadata *pb.Metadata) (*pb.Response, error) {
	fileStoragePath := filepath.Join(s.storageDirectory, metadata.Uuid)

	if _, err := os.Stat(fileStoragePath); !os.IsNotExist(err) {
		err = os.Remove(fileStoragePath)
		if err != nil {
			log.Printf("cannot remove %s\n", fileStoragePath)
			return nil, err
		}
	} else {
		return nil, errors.New(fmt.Sprintf("file %s does not exist on disk", metadata.Uuid))
	}

	return &pb.Response{
		Status: true,
	}, nil
}

func (s Service) Healthcheck(stream pb.Node_HealthcheckServer) error {
	for {
		metadata, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Empty{})
		}
		if err != nil {
			return err
		}

		log.Printf("received healthcheck from %s\n", metadata.Name)
	}
}

func (s Service) mustEmbedUnimplementedNodeServer() {
	//TODO implement me
	panic("implement me")
}

func NewService(nodename, storageDirectory string, store KVStore) *Service {
	return &Service{
		NodeName:         nodename,
		storageDirectory: storageDirectory,
		kvstore:          store,
	}
}
