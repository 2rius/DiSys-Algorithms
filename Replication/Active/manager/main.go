package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	api "github.com/2rius/DiSys-Algorithms/tree/main/Replication/Active/api"
	"google.golang.org/grpc"
)

var (
	defaultPort = 5000
	port        = flag.Int("port", defaultPort, "port")
)

func main() {
	flag.Parse()

	var port = uint32(*port)

	// Create listener tcp on port ownPort
	var attempts = 0

	var list net.Listener

	for {
		var err error
		list, err = net.Listen("tcp", fmt.Sprintf(":%v", port))

		if err != nil {
			if attempts > 5 { // Allows 5 managers by just execution 'go run .' without adding flag --port
				log.Fatalf("Failed to listen on port: %v - with %d attempts", err, attempts)
			}

			attempts++
			port++

			continue
		}

		break
	}

	myip := fmt.Sprintf("127.0.0.1:%d", port)

	m := &Manager{
		id:   myip,
		data: make(map[string]string),
	}

	log.Printf("Opened on port: %s\n", m.id)

	grpcServer := grpc.NewServer()
	api.RegisterManagerServer(grpcServer, m)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
