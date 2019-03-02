package store

import (
	"fmt"
	"strconv"
	"sync"
)

// MemoryStorer implements a in-memory key-value store
type MemoryStorer struct {
	store map[string]string
	mux   sync.Mutex
}

func (m *MemoryStorer) Get(key string) (string, error) {
	m.mux.Lock()
	defer m.mux.Unlock()
	if value, ok := m.store[key]; ok {
		return value, nil
	}
	return "", fmt.Errorf("key %s not found", key)
}

func (m *MemoryStorer) Set(key string, value string) error {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.store[key] = value
	return nil
}

func (m *MemoryStorer) Increment(key string) error {
	m.mux.Lock()
	defer m.mux.Unlock()
	i, _ := strconv.Atoi(m.store[key])
	m.store[key] = strconv.FormatInt(int64(i+1), 10)
	return nil
}

// NewMemoryStorer creates a new MemoryStorer
func NewMemoryStorer() *MemoryStorer {
	return &MemoryStorer{
		store: make(map[string]string),
	}
}
