package main

// From assignment 5

import (
	"context"
	"fmt"
	"log"
	"time"

	api "github.com/2rius/DiSys-Algorithms/tree/main/Replication/Passive/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

/*
send election to all higher IDs

	if no higher IDs call coordinator(i won)

If timeout on ALL higher IDs for answer

	i'm also winner, coordinate

if timeout on coordinate

	call new election

New process calls coordinate immediately
*/
func (m *Manager) CallElection() {
	votes := len(m.peers)

	log.Printf("Calling election: need %d votes\n", votes)

	for ip, o := range m.peers {
		log.Printf("[ME] %s > %s : %t\n", m.id, ip, m.id > ip)
		if m.id > ip {
			votes--
			continue
		}

		// REQUIRE ANSWER
		var _, err = o.Election(m.ctx, &api.Elect{})

		if err != nil {
			votes--
			delete(m.peers, ip)
		}
	}

	log.Printf("Calling election: need %d votes\n", votes)

	if votes == 0 {
		log.Println("Peers died or no one has higher id than me, I must be new leader then")
		m.primaryId = m.id
		// m.SendCoordinate()
	} else {
		m.WaitForCoordinationTimeout(5_000)
	}
}

func (m *Manager) WaitForCoordinationTimeout(ms int) {
	// [TODO] use states instead
	m.expectingAnswer = true

	select {
	case <-m.timeoutCoordinate:
		log.Println("New prim lives")
	case <-time.After(time.Duration(ms) * time.Millisecond):
		log.Println("Coordination timeout, call new election")
		m.CallElection()
	}

	m.expectingAnswer = false
}

/*
OnElection
-----------------
if I have highest id

	coordinate

for all ID's higher than me

	call election, wait for answer timeout
	on timeouts answer, bullys dead
		i win, coordinate
	# Maybe
	if answer but no coordinate for higher ID
		recall election
*/
func (m *Manager) Election(ctx context.Context, in *api.Elect) (*api.Void, error) {

	log.Println("Somebody wants election, I'll ask the higher ups if i got any")

	m.CallElection()

	return &api.Void{}, nil

}

/*
Send I'm leader to all
*/
func (m *Manager) SendCoordinate() {
	log.Println("I'm the new primary, coordinating with everybody")

	for ip, o := range m.peers {
		var _, err = o.Coordinate(m.ctx, &api.Coord{})

		if err != nil {
			fmt.Printf("Err: %v", err)
			fmt.Println("[SendCoordinate] could not connect to peer")
			delete(m.peers, ip)
		}
	}
}

/*
OnCoordinate
---------------

	i'm actually higher ID
		do Election
	set new prim id
*/
func (m *Manager) Coordinate(ctx context.Context, in *api.Coord) (*api.Void, error) {

	log.Println("New primary calling coordinate")

	var addr string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		addr = md.Get("address")[0]
	}

	// Was I the old leader?
	wasLeader := (m.id == m.primaryId && m.id != addr)

	m.primaryId = addr

	select {
	case m.timeoutCoordinate <- true:
	default:
	}

	// Save new client/boss
	if _, ok := m.peers[addr]; !ok {
		log.Println("Saving new primary")

		conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

		if err != nil {
			log.Fatalf("could not reconnect: %v", err)
		}

		defer conn.Close()

		c := api.NewManagerClient(conn)
		m.peers[addr] = c
	}

	if wasLeader {
		log.Println("[OLD LEADER] I will update new leader with my data as the old leader")

		// [TODO] send data to new leader
	}

	return &api.Void{}, nil
}
