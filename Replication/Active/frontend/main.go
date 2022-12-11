package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	api "github.com/2rius/DiSys-Algorithms/tree/main/Replication/Active/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	defaultPort   = 5000
	managerAmount = flag.Int("managers", 3, "managers")
	ctx, cancel   = context.WithCancel(context.Background())
	managers      = make(map[string]api.ManagerClient)
	conns         = make(map[string]*grpc.ClientConn)
)

func main() {
	flag.Parse()

	ports := flag.Args()

	defer cancel()

	if len(ports) > 0 {
		for _, portS := range ports {
			port, err1 := strconv.Atoi(portS)
			if err1 != nil {
				log.Fatalf("%s illegal port\n", portS)
			}

			var conn *grpc.ClientConn
			log.Printf("Trying to dial: %v\n", port)
			ip := fmt.Sprintf("127.0.0.1:%d", port)

			conn, err2 := grpc.DialContext(ctx, ip, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
			if err2 != nil {
				log.Fatalf("Could not connect: %s", err2)
			}

			defer conn.Close()
			log.Printf("Succes connecting to: %v\n", port)
			conns[ip] = conn
			c := api.NewManagerClient(conn)
			managers[ip] = c
		}
	} else {
		for i := 0; i < *managerAmount; i++ {
			var port = defaultPort + i

			var conn *grpc.ClientConn
			log.Printf("Trying to dial: %v\n", port)
			ip := fmt.Sprintf("127.0.0.1:%d", port)

			conn, err := grpc.DialContext(ctx, ip, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

			if err != nil {
				log.Fatalf("Could not connect: %s", err)
			}

			defer conn.Close()
			log.Printf("Succes connecting to: %v\n", port)
			conns[ip] = conn
			c := api.NewManagerClient(conn)
			managers[ip] = c
		}
	}

	mainloop()
}

func mainloop() {
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
			callSet(key, value)

		case "get":
			callGet(args)

		case "exit":
			os.Exit(0)
		}

		fmt.Print("Enter command: ")
	}
}

func callGet(args []string) {
	dataAcquired := false
	var values = make(map[string]string)

	//send get request to all managers
	for ip, m := range managers {
		data, err := m.Get(ctx, &api.Void{})
		if err != nil {
			log.Printf("Manager %s not responding, deleting manager from managers\n", ip)
			delete(managers, ip)
		} else if !dataAcquired {
			/*
				Saves data from first responding client.
				Better solution would be to include lamport clock,
				and save data from highest lamport clock manager.
			*/
			dataAcquired = true
			values = data.Data
		}
	}

	if !dataAcquired {
		log.Fatalln("No managers responded, assuming no managers alive")
	}

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
}

func callSet(key string, value string) {
	managerResponded := false

	//send set request to all managers
	for ip, m := range managers {
		_, err := m.Set(ctx, &api.Value{Key: key, Value: value})
		if err != nil {
			log.Printf("Manager %s not responding, deleting manager from managers\n", ip)
			delete(managers, ip)
		} else if !managerResponded {
			managerResponded = true
		}
	}

	if managerResponded {
		log.Println("Value sucessfully set")
	} else {
		log.Fatalln("No managers responded, assuming no managers alive")
	}
}
