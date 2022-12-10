package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	api "github.com/2rius/DiSys-Algorithms/tree/main/LogicalTime/Lamport/api"
	"google.golang.org/grpc"
)

type Client struct {
	name  string
	port  string
	clock uint32
	mu    sync.Mutex
}

var name = flag.String("name", "localhost", "The server name")
var port = flag.String("port", "5000", "The server port")

func getInput() (string, error) {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	return input, nil
}

func startClient(c *Client) {
	fmt.Printf("Enter your username: ")

	var connectClient api.Chat_ConnectClient

	for {
		conn, err := grpc.Dial(c.name+":"+c.port, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connect to the server %v", err)
		}

		chatClient := api.NewChatClient(conn)

		connectClient, err = chatClient.Connect(context.Background())
		if err != nil {
			log.Fatalf("Could not connect to the server %v", err)
		}

		username, err := getInput()

		if err != nil {
			log.Fatalf("Could not get username: %v", err)
		}

		connectClient.Send(&api.MsgClient{
			Name:   c.name,
			Msg:    username,
			Lclock: c.clock,
		})

		msg, err1 := connectClient.Recv()

		if err1 == nil {
			fmt.Printf("%s [%s] (%d): %s\n", msg.Timestamp.AsTime().Local().Format("15:04:05"), msg.Name, c.clock, msg.Msg)
			break
		}
		fmt.Println("Username already in use, please enter a new username")
	}

	go listenforMsg(connectClient, c)

	for {
		input, err := getInput()
		if err != nil {
			log.Fatalf("Could not retrieve user input")
		}

		if input == "exit" {
			break
		}

		c.incrementClock()

		connectClient.Send(&api.MsgClient{
			Name:   c.name,
			Msg:    input,
			Lclock: c.clock,
		})

	}

	connectClient.CloseSend()
}

func listenforMsg(connectClient api.Chat_ConnectClient, c *Client) {
	for {
		msg, err := connectClient.Recv()
		if err != nil {
			log.Fatalf("Could not receive the message %v", err)
		}

		c.LamportMax(msg.Lclock)

		fmt.Printf("%s [%s] (%d): %s\n", msg.Timestamp.AsTime().Local().Format("15:04:05"), msg.Name, c.clock, msg.Msg)
	}
}

func main() {
	flag.Parse() //Get the port from the command line

	client := &Client{
		name:  *name,
		port:  *port,
		clock: 0,
	}

	startClient(client)
}
