package main

import (
	"context"
	"log"
	"sync"
	"time"

	api "github.com/2rius/DiSys-Algorithms/tree/main/MutualExclusion/RicartAgrawala/grpc"
)

type STATE int64

const (
	RELEASED STATE = iota
	HELD
	WANTED
)

type Peer struct {
	api.UnimplementedRicartAgrawalaServer
	id           uint32
	clients      map[uint32]api.RicartAgrawalaClient
	defered      []uint32
	ctx          context.Context
	state        STATE
	clock        uint32
	requestsSent int
	mu           sync.Mutex
}

func (p *Peer) Request(ctx context.Context, req *api.Request) (*api.RequestBack, error) {
	/*
		I received a request from a Peer

		if my state is HELD or (WANTED and I have higher priority than my Peer)
			then i defer with a REPLY until later(queue)
		else
			I send them a REPLY right away
	*/

	log.Printf("(L: %d) RECV: Received request from %d\n", p.clock, req.Id)
	if p.state == HELD || (p.state == WANTED && (p.clock < req.Clock || (p.clock == req.Clock && p.id > req.Id))) {
		log.Printf("(L: %d) Deferring request from %v\n", p.clock, req.Id)
		p.defered = append(p.defered, req.Id)

		p.LamportMax(req.Clock)
	} else {
		if p.state == WANTED {
			p.requestsSent++
			p.incrementClock()
			log.Printf("(L: %d) SEND: Rerequesting critical section to higher priority Peer\n", p.clock)
			p.clients[req.Id].Request(ctx, &api.Request{Id: p.id, Clock: p.clock})
		}
		p.incrementClock()
		log.Printf("(L: %d) SEND: Sending reply right away to %v\n", p.clock, req.Id)
		p.clients[req.Id].Reply(ctx, &api.Reply{
			Clock: p.clock,
			Id:    p.id,
		})
	}

	rep := &api.RequestBack{}
	return rep, nil
}

func (p *Peer) Reply(ctx context.Context, req *api.Reply) (*api.ReplyBack, error) {
	/*
		I received REPLY from other Peer

		if i received N-1 REPLIES(1 for each Request send) then
			i take CS(Critical Sec)
		else
			i wait for next reply
	*/

	p.LamportMax(req.Clock)

	p.requestsSent--
	log.Printf("(L: %d) RECV: Got a reply. Missing %d replies.\n", p.clock, p.requestsSent)

	if p.requestsSent == 0 {
		log.Printf("(L: %d) RECV: All requests replied.\n", p.clock)
		go p.criticalSection()
	}

	rep := &api.ReplyBack{}
	return rep, nil
}

func (p *Peer) criticalSection() {
	/*
		I have critical section
		When done I call sendReplyToAllDefered
	*/

	p.incrementClock()
	log.Printf("(L: %d) CS: Started working\n", p.clock)
	p.state = HELD
	time.Sleep(5 * time.Second)
	p.state = RELEASED
	p.incrementClock()
	log.Printf("(L: %d) CS: Done working\n", p.clock)

	p.sendReplyToAllDefered()
}

func (p *Peer) sendRequestToAll() {
	p.state = WANTED

	p.incrementClock()

	request := &api.Request{
		Id:    p.id,
		Clock: p.clock,
	}

	p.requestsSent = len(p.clients)

	log.Printf("(L: %d) SEND: Sending request to all Peers. Missing %d replies.\n", p.clock, p.requestsSent)

	for id, client := range p.clients {
		_, err := client.Request(p.ctx, request)

		if err != nil {
			log.Printf("Something went wrong with %v\n", id)
		}
	}
}

func (p *Peer) sendReplyToAllDefered() {

	p.incrementClock()

	if len(p.defered) != 0 {
		log.Printf("(L: %d) SEND: Sending reply to all deferred Peers.\n", p.clock)
	} else {
		log.Printf("(L: %d) SEND: No deferred Peers.\n", p.clock)
	}

	reply := &api.Reply{
		Clock: p.clock,
		Id:    p.id,
	}

	/*
		Reply all defered

		then clear p.defered
	*/

	for _, id := range p.defered {
		_, err := p.clients[id].Reply(p.ctx, reply)

		if err != nil {
			log.Printf("Something went wrong with %v\n", id)
		}
	}

	p.defered = make([]uint32, 0)
}
