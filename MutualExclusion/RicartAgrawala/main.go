package main

// From assignment 4

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	api "github.com/2rius/DiSys-Algorithms/tree/main/MutualExclusion/RicartAgrawala/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	p := &Peer{
		id:           port,
		clients:      make(map[uint32]api.RicartAgrawalaClient),
		defered:      make([]uint32, 0),
		ctx:          ctx,
		state:        RELEASED,
		clock:        0,
		requestsSent: 0,
	}

	log.Printf("Opened on port: %d\n", p.id)

	grpcServer := grpc.NewServer()
	api.RegisterRicartAgrawalaServer(grpcServer, p)

	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to server %v", err)
		}
	}()

	for i := 0; i < 3; i++ {
		PeerPort := uint32(5000 + i)

		if PeerPort == p.id {
			continue
		}

		var conn *grpc.ClientConn
		log.Printf("Trying to dial: %v\n", PeerPort)
		conn, err := grpc.Dial(fmt.Sprintf(":%v", PeerPort), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}

		defer conn.Close()
		log.Printf("Succes connecting to: %v\n", PeerPort)
		c := api.NewRicartAgrawalaClient(conn)
		p.clients[PeerPort] = c
	}

	log.Printf("I am connected to %d clients", len(p.clients))

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if p.state == RELEASED {
			log.Println("Requesting to work in critical section...")
			p.sendRequestToAll()
		} else {
			log.Println("Already working in critical section...")
		}
	}
}
