package main

import (
	"context"
	"sync"

	api "github.com/2rius/DiSys-Algorithms/tree/main/Raft/api"
)

type Peer struct {
	api.UnimplementedPeerServer
	mu          sync.Mutex
	id          string
	peerIds     []int
	peerClients map[int]api.PeerClient
	ctx         context.Context
}
