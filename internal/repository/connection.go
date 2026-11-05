package repository

import "sync"

type ConnectionRepository struct {
	connections map[string][]string
	mu          sync.Mutex
}

func NewConnectionRepository() *ConnectionRepository {
	return &ConnectionRepository{
		connections: make(map[string][]string),
	}
}

func (r *ConnectionRepository) IsExist(collection string) bool {
	_, ok := r.connections[collection]
	if !ok {
		return false
	}

	return true
}

func (r *ConnectionRepository) Create(collection, addr string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.connections[collection] = append(r.connections[collection], addr)
}

func (r *ConnectionRepository) Delete(collection, addr string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for col, addrs := range r.connections {
		if col == collection {
			if len(addrs) == 0 {
				delete(r.connections, collection)
				return
			}

			r.deleteAddr(addrs, addr)
			return
		}
	}
}

func (r *ConnectionRepository) deleteAddr(addrs []string, addr string) {
	var newSlice []string
	for i := range addrs {
		if addrs[i] != addr {
			newSlice = append(newSlice, addrs[i])
		}
	}

	addrs = newSlice
}
