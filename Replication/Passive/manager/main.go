package main

// From assignment 5

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	api "github.com/2rius/DiSys-Algorithms/tree/main/Replication/Passive/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var port = flag.Int("port", 5000, "port")

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var port = uint32(*port)

	// Create listener tcp on port ownPort
	var attempts = 0

	var list net.Listener

	for {
		var err error
		list, err = net.Listen("tcp", fmt.Sprintf(":%v", port))

		if err != nil {
			if attempts > 5 {
				log.Fatalf("Failed to listen on port: %v - with %d attempts", err, attempts)
			}

			attempts++
			port++

			continue
		}

		break
	}

	myip := fmt.Sprintf("127.0.0.1:%d", port)
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("address", myip))

	m := &Manager{
		id:                myip,
		peers:             make(map[string]api.ManagerClient),
		frontends:         make(map[string]api.FrontendClient),
		ctx:               ctx,
		primaryId:         "127.0.0.1:5002", // Default primary address, given there are 3 managers, "127.0.0.1:"+4999+Amount of managers
		data:              make(map[string]string),
		timeoutHeartbeat:  make(chan bool, 1),
		timeoutCoordinate: make(chan bool, 1),
		expectingAnswer:   false,
	}

	log.Printf("Opened on port: %s\n", m.id)

	grpcServer := grpc.NewServer()
	api.RegisterManagerServer(grpcServer, m)

	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to server %v", err)
		}
	}()

	for i := 0; i < 3; i++ {
		peerPort := uint32(5000 + i)

		if peerPort == port {
			continue
		}

		var conn *grpc.ClientConn
		log.Printf("Trying to dial: %v\n", peerPort)
		ip := fmt.Sprintf("127.0.0.1:%d", peerPort)

		conn, err := grpc.DialContext(ctx, ip, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}

		defer conn.Close()
		log.Printf("Succes connecting to: %v\n", peerPort)
		c := api.NewManagerClient(conn)
		m.peers[ip] = c
	}

	log.Printf("I am connected to %d clients", len(m.peers))

	<-time.After(3 * time.Second)

	m.MainLoop()
}
