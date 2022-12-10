package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	api "github.com/2rius/DiSys-Algorithms/tree/main/LogicalTime/Lamport/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	api.UnimplementedChatServer
	name        string
	port        string
	mu          sync.Mutex
	clock       uint32
	clients     map[string]api.Chat_ConnectServer
	clientNames map[string]string
}

func valueInMap(m map[string]string, value string) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

func (s *Server) Connect(stream api.Chat_ConnectServer) error {
	log.Println("Client connected.")

	m, err := stream.Recv()
	if err != nil {
		log.Printf("Could not receive the message %v", err)
	}

	if valueInMap(s.clientNames, m.Msg) {
		log.Printf("Username already exists: %v. Client will disconnect.", m.Msg)
		return errors.New("username already exists")
	}

	p, _ := peer.FromContext(stream.Context())
	s.clients[p.Addr.String()] = stream
	log.Printf("%s (%s) connected to the chat\n", m.Msg, p.Addr.String())

	s.clientNames[p.Addr.String()] = m.Msg

	timeConnect := time.Now().Local()

	message := &api.MsgServer{
		Name:      "Server",
		Msg:       fmt.Sprintf("Welcome to ChittyChat %v", m.Msg),
		Lclock:    s.clock,
		Timestamp: timestamppb.New(timeConnect),
	}

	log.Println("Broadcast new client connection to all connected clients.")
	for _, st := range s.clients {
		st.Send(message)
	}

	for {
		log.Printf("Listen for messages from client: %s\n", p.Addr.String())
		msg, err := stream.Recv()

		if err != nil {
			//if error is EOF, the client has disconnected
			if status.Code(err).String() == "Canceled" || status.Code(err).String() == "EOF" {
				log.Printf("%s (%s) disconnected from the chat\n", m.Msg, p.Addr.String())
				break
			} else {
				log.Printf("Could not receive the message %v", err)
				break
			}
		}

		s.LamportMax(msg.Lclock)

		log.Printf("Lampert: %d - Received message from %s: %s", s.clock, msg.Name, msg.Msg)

		time := time.Now().Local()

		s.incrementClock()

		message := &api.MsgServer{
			Name:      s.clientNames[p.Addr.String()],
			Msg:       msg.Msg,
			Lclock:    s.clock,
			Timestamp: timestamppb.New(time),
		}

		log.Printf("Broadcast message from %s to all other clients with new server clock lampert(%d).\n", p.Addr.String(), s.clock)
		for _, st := range s.clients {
			st.Send(message)
		}
	}

	timeLeave := time.Now().Local()

	messageLeave := &api.MsgServer{
		Name:      "Server",
		Msg:       fmt.Sprintf("Goodbye %s", s.clientNames[p.Addr.String()]),
		Lclock:    s.clock,
		Timestamp: timestamppb.New(timeLeave),
	}

	log.Printf("Broadcast the leave of %s to other clients.\n", p.Addr.String())
	for _, st := range s.clients {
		st.Send(messageLeave)
	}

	delete(s.clients, p.Addr.String())
	delete(s.clientNames, p.Addr.String())
	return nil
}

var name = flag.String("name", "localhost", "The server name")
var port = flag.String("port", "5000", "The server port")

func startServer(s *Server) {
	listen, err := net.Listen("tcp", s.name+":"+s.port)
	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}

	grpcServer := grpc.NewServer()

	log.Printf("Starting server on port %v", s.port)

	api.RegisterChatServer(grpcServer, s)
	serveError := grpcServer.Serve(listen)

	if serveError != nil {
		log.Fatalf("Could not start the server %v", serveError)
	}
}

func main() {
	flag.Parse() //Get the port from the command line

	server := &Server{
		name:        *name,
		port:        *port,
		clock:       0,
		clients:     make(map[string]api.Chat_ConnectServer),
		clientNames: make(map[string]string),
	}

	go startServer(server)

	for {
		time.Sleep(1 * time.Second)
	}
}
