package main

import (
	"log"
	"net"
	"fmt"
	"errors"

	"github.com/nsahifa/protobuf"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"github.com/orcaman/concurrent-map"
)

type kv struct {
	cmap cmap.ConurrentMap
}

func (server *kv) insertItem(ctx context.Context, input *kvs.InsertionRequest) (*kvs.InsertionResponse, error) {
	
	server.cmap.Set(input.Key, input.Value)

	log.Printf("the pair (%s, %s) is stored succesffuly", input.Key, input.Value)

	return &kvs.InsertionResponse{Success: true}, nil
}

func (server *kv) getItem(ctx context.Context, input *kvs.GetItem) (*kvs.GetResponse, error) {

	value, status := server.cmap.Get(input.Key)

	if (status == false) {
		log.Printf("Loading failed for the value from the key %s", input.Key)
		return nil, errors.New(fmt.Sprintf("Invalid key : %s", input.Key))
	}

	value, status = value.(string)
	log.Printf("Loading successed for the value from the key %s", input.Key)

	return &kvs.GetResponse{Key: input.Key, Value: value.(string)}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Listenning failed : %v", err)
	}

	server := grpc.NewServer()
	kvs.RegisterKVSServer(server, &kv{cmap: cmap.New()})

	s.Serve(listener)
}