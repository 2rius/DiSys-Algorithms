package main

// From assignment 5

import (
	"context"
	"log"
	"time"

	api "github.com/2rius/DiSys-Algorithms/tree/main/Replication/Passive/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type Manager struct {
	api.UnimplementedManagerServer
	id                string
	peers             map[string]api.ManagerClient
	frontends         map[string]api.FrontendClient
	ctx               context.Context
	primaryId         string
	data              map[string]string
	timeoutHeartbeat  chan bool
	timeoutCoordinate chan bool
	expectingAnswer   bool
}

/*
Get data from primary
*/
func (m *Manager) Get(ctx context.Context, req *api.Void) (*api.Data, error) {
	log.Println("[PRIMARY] A client requested data")

	rep := &api.Data{
		Data: m.data,
	}

	return rep, nil
}

/*
Get data from frontend

	Update all backup managers
*/
func (m *Manager) Set(ctx context.Context, req *api.Value) (*api.Void, error) {
	log.Println("[PRIMARY] A client updated data")

	var addr string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		addr = md.Get("address")[0]
	}

	if _, ok := m.frontends[addr]; !ok {
		conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

		if err != nil {
			log.Fatalf("could not reconnect: %v", err)
		}

		c := api.NewFrontendClient(conn)
		m.frontends[addr] = c
	}

	m.UpdatePeers(req)

	m.data[req.Key] = req.Value

	rep := &api.Void{}

	return rep, nil
}

func (m *Manager) SendHeartbeatToAll() {
	log.Println("[PRIMARY] Sending heartbeat")

	for ip, frontendClient := range m.frontends {
		_, err := frontendClient.Heartbeat(m.ctx, &api.Primary{})

		if err != nil {
			delete(m.frontends, ip)
		}
	}

	for ip, peer := range m.peers {
		_, err := peer.Heartbeat(m.ctx, &api.Void{})

		if err != nil {
			delete(m.peers, ip)
		}
	}

	<-time.After(2 * time.Second)
}

/*
listens to heartbeat from primary

	reset timeout to 5000ms

--------

TimeoutEnds

	HoldElection
*/
func (m *Manager) Heartbeat(ctx context.Context, req *api.Void) (*api.Void, error) {
	// call heartbeat channel
	m.timeoutHeartbeat <- true

	rep := &api.Void{}
	return rep, nil
}

func (m *Manager) HeartbeatTimeout(reset <-chan bool, ms int) {
	select {
	case <-reset:
		// heartbeat
		log.Println("[HB] .")
	case <-time.After(time.Duration(ms) * time.Millisecond):
		// Prim dead
		log.Println("[HB] Prim Dead - do election")
		delete(m.peers, m.primaryId)
		m.CallElection()
	}
}

/*
Updates peers when data is updated
*/
func (m *Manager) UpdatePeers(req *api.Value) error {
	log.Println("[PRIMARY] Updating peers on latest data")

	ips := make([]string, 0, len(m.frontends))
	for k := range m.frontends {
		ips = append(ips, k)
	}

	for ip, client := range m.peers {
		attempts := 0
		for {
			_, err := client.Update(m.ctx, &api.UpdateData{
				Key:       req.Key,
				Value:     req.Value,
				Frontends: ips,
			})
			if err == nil {
				break
			} else if attempts == 3 {
				delete(m.peers, ip)
				break
			}
			attempts++
		}
	}
	return nil
}

/*
Update data from primary
*/
func (m *Manager) Update(ctx context.Context, req *api.UpdateData) (*api.Void, error) {
	log.Println("I'm getting updated to latest info!")

	m.data[req.Key] = req.Value

	for _, ip := range req.Frontends {
		conn, err := grpc.DialContext(ctx, ip, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

		if err != nil {
			continue
		}

		c := api.NewFrontendClient(conn)

		m.frontends[ip] = c
	}

	return &api.Void{}, nil
}

/*
Update data from old primary
*/
func (m *Manager) UpdateLeader(ctx context.Context, req *api.UpdateLeaderData) (*api.Void, error) {
	log.Println("I'm getting updated to latest info by old leader!")

	m.data = req.Data

	for _, ip := range req.Frontends {
		conn, err := grpc.DialContext(ctx, ip, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

		if err != nil {
			continue
		}

		c := api.NewFrontendClient(conn)

		m.frontends[ip] = c
	}

	return &api.Void{}, nil
}

/*
NotifyClient, I'm new primary
*/
func (m *Manager) NotifyClients() error {
	for ip, client := range m.frontends {
		var _, err = client.Heartbeat(m.ctx, &api.Primary{})

		if err != nil {
			delete(m.frontends, ip)
		}
	}

	return nil
}

func (m *Manager) MainLoop() {
	for {
		if m.id == m.primaryId {
			m.Leader()
		} else {
			m.Backup()
		}
	}
}

func (m *Manager) Leader() {
	log.Println("leader loop")

	m.SendCoordinate()

	for {
		// Am I still leader?
		if m.id != m.primaryId {
			return
		}

		// send heartbeat to all: peers and clients
		m.SendHeartbeatToAll()
	}
}

func (m *Manager) Backup() {
	log.Println("backup loop")

	for {
		if m.id == m.primaryId {
			m.timeoutHeartbeat <- true
			return
		}

		m.HeartbeatTimeout(m.timeoutHeartbeat, 5_000)
	}
}
