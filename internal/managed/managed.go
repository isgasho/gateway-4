package managed

import (
	"google.golang.org/grpc"
	"sync"
)

type ManagedChannel struct {
	conns map[string]*grpc.ClientConn
	rw    sync.RWMutex
}

func NewManagedChannel() *ManagedChannel {

	return &ManagedChannel{
		conns: make(map[string]*grpc.ClientConn),
		rw:    sync.RWMutex{},
	}
}

func (m *ManagedChannel) Create(endpoint string) *grpc.ClientConn {
	conn := m.getConn(endpoint)
	if conn != nil {
		return conn
	}
	return m.doCreate(endpoint)
}

func (m *ManagedChannel) Release(endpoint string) {

	m.rw.Lock()
	defer m.rw.Unlock()
	delete(m.conns, endpoint)
}

func (m *ManagedChannel) getConn(endpoint string) *grpc.ClientConn {

	m.rw.RLock()
	defer m.rw.RUnlock()
	if conn, ok := m.conns[endpoint]; ok {
		return conn
	}
	return nil
}

func (m *ManagedChannel) doCreate(endpoint string) *grpc.ClientConn {
	m.rw.RLock()
	defer m.rw.RUnlock()
	if conn, err := grpc.Dial(endpoint, grpc.WithInsecure()); err == nil {
		m.conns[endpoint] = conn
		return conn
	}
	return nil
}
