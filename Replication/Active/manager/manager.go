package main

import (
	"context"
	"log"

	api "github.com/2rius/DiSys-Algorithms/tree/main/Replication/Active/api"
)

type Manager struct {
	api.UnimplementedManagerServer
	id   string
	data map[string]string
}

/*
Get data from the manager
*/
func (m *Manager) Get(ctx context.Context, req *api.Void) (*api.Data, error) {
	log.Println("[PRIMARY] A client requested data")

	rep := &api.Data{
		Data: m.data,
	}

	return rep, nil
}

/*
Update data in the manager
*/
func (m *Manager) Set(ctx context.Context, req *api.Value) (*api.Void, error) {
	log.Println("[PRIMARY] A client updated data")

	m.data[req.Key] = req.Value

	rep := &api.Void{}

	return rep, nil
}
