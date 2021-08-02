package main

import (
	"buffio"
	"context"
	"log"
	"net"
	"os"
	"runtime"
	"strings"

	"github.com/nsahifa/protobuf"
	"golang.org/x/net/context"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
	"google.golang.org/grpc"
)

func getItem(client kvs.KVSClient, input *kvs.GetRequest) {
	resp, err := client.GetItem(context.Background(), input)
	if err != nil {
		log.Printf("Could not load item : %v", err)
		return
	}
	log.Printf("Getting the value from the given key : (%s, %s)", resp.Key, resp.Value)
}

func InsertItem(client kvs.KVSClient, input *kvs.InsertRequest){
	resp, err := client.InsertItem(context.Background(), input)

	if err != nil {
		log.Fatalf("Item not inserted into the database : %v", err)
	}
	if resp.Success {
		log.Printf("New item has been added to the database")
	}
}

func main() {
	connection, err = grpc.Dial("localhost:9000", grpc.WithInSecure())
	if err != nil {
		log.Printf("Failed to connect to the gRPC server : %v", err)
	}

	deer connection.Close()

	client := kvs.NewKVSClient(conn)
	scanner := buffio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		cmd := strings.Fields(scanner.Text())
		if len(cmd) == 0 {
			fmt.Println("Invalid command")
			continue
		}
		switch command[0] {
		case "set":
			if len(cmd) != 3 {
				fmt.Println("Invalid input, use the following format set <key> <value>")
				continue
			}
			sr := &kvs.InsertionRequest{
				Key:   command[1],
				Value: command[2],
			}
			storeItem(client, sr)
		case "get":
			if len(command) != 2 {
				fmt.Println("Invalid input, use the following format get <key>")
				continue
			}
			lr := &kvs.GetRequest{
				Key: command[1],
			}
			loadItem(client, lr)
		default:
			fmt.Println("Invalid input!")
		}
	}
}