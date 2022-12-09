package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	api "github.com/2rius/DiSys-Algorithms/tree/main/Elections/Bully/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type Peer struct {
	api.UnimplementedPeerServer
	id                string
	peers             map[string]api.PeerClient
	ctx               context.Context
	primaryId         string
	timeoutCoordinate chan bool
	expectingAnswer   bool
}

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

	p := &Peer{
		id:                myip,
		peers:             make(map[string]api.PeerClient),
		ctx:               ctx,
		primaryId:         myip,
		timeoutCoordinate: make(chan bool, 1),
		expectingAnswer:   false,
	}

	log.Printf("Opened on port: %s\n", p.id)

	grpcServer := grpc.NewServer()
	api.RegisterPeerServer(grpcServer, p)

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
		conn, err := grpc.Dial(fmt.Sprintf(":%v", peerPort), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}

		defer conn.Close()
		log.Printf("Succes connecting to: %v\n", peerPort)
		c := api.NewPeerClient(conn)
		ip := fmt.Sprintf("127.0.0.1:%d", peerPort)
		p.peers[ip] = c
	}

	log.Printf("I am connected to %d clients", len(p.peers))

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		p.CallElection()
	}
}
