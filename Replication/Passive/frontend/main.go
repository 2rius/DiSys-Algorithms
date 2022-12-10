package main

// From assignment 5

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	api "github.com/2rius/DiSys-Algorithms/tree/main/Replication/Passive/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func (f *FrontendListener) Heartbeat(ctx context.Context, req *api.Primary) (*api.Void, error) {
	// when heartbeat received, update new primary

	var addr string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		addr = md.Get("address")[0]
	}

	if addr != f.address {
		f.address = addr

		conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}

		f.conn = conn
		f.primaryClient = api.NewManagerClient(conn)
	}

	return &api.Void{}, nil
}

type FrontendListener struct {
	api.UnimplementedFrontendServer
	address       string
	primaryClient api.ManagerClient
	conn          *grpc.ClientConn
	ctx           context.Context
}

var address = flag.String("primary", "127.0.0.1:5002", "primary ip")
var startPort = flag.Int("port", 5003, "port")

func main() {
	flag.Parse()

	var myPort = uint32(*startPort)

	myip := fmt.Sprintf("127.0.0.1:%d", myPort)

	ctx, _ := context.WithCancel(context.Background())
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("address", myip))

	// Dial initial primary manager
	conn, err := grpc.DialContext(ctx, *address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to client: %v", err)
	}

	f := &FrontendListener{
		address:       *address,
		primaryClient: api.NewManagerClient(conn),
		conn:          conn,
		ctx:           ctx,
	}

	// Start heartbeat server(frontendListener)
	list, err := net.Listen("tcp", myip)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterFrontendServer(grpcServer, f)

	// Start listening for heartbeat
	go func() {
		fmt.Printf("\nlistening on %s\n", myip)

		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to server %v", err)
		}

	}()

	// Start console for listening to commands
	fmt.Println("Commands:")
	fmt.Println("\"set <key> <value>\" - set new value for key")
	fmt.Println("\"get <key>\" - get value from key")
	fmt.Println("\"get\" - get all keys and values")
	fmt.Println("\"exit\" - exit the program")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter command: ")

	for scanner.Scan() {
		//read in command
		command := scanner.Text()

		//split command into args
		args := strings.Split(command, " ")

		//switch on command
		switch args[0] {
		case "set":
			key := args[1]
			value := args[2]

			//send set request to primary
			_, err := f.primaryClient.Set(ctx, &api.Value{Key: key, Value: value})
			if err != nil {
				log.Println("Error sending set request")
			}

			log.Println("Value sucessfully set")

		case "get":
			data, err := f.primaryClient.Get(ctx, &api.Void{})
			if err != nil {
				log.Println("Error sending get request")
			}

			values := data.Data

			if len(args) == 1 {
				//print all values
				for key, value := range values {
					fmt.Printf("%s: %s\n", key, value)
				}
			} else {
				//print value for key
				key := args[1]
				value := values[key]
				if value != "" {
					fmt.Printf("%s: %s\n", key, value)
				} else {
					fmt.Printf("Key %s not found\n", key)
				}
			}

		case "exit":
			os.Exit(0)
		}

		fmt.Print("Enter command: ")
	}
}
