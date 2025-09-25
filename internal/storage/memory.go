package storage

import (
	"sync"
)

type MemoryStorage struct {
	mu      sync.Mutex
	clients []*Client
}


func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (m *MemoryStorage) Add(client *Client) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.clients) > 0 {
		client.ID = m.clients[len(m.clients)-1].ID + 1
	} else {
		client.ID = 1
	}

	m.clients = append(m.clients, client)
	return nil
}

func (m *MemoryStorage) GetAll() ([]*Client, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.clients, nil
}

func (m *MemoryStorage) GetByID(id int) (*Client, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, c := range m.clients {
		if c.ID == id {
			return c, nil
		}
	}
	return nil, ErrClientNotFound(id)
}

func (m *MemoryStorage) Update(id int, newName, newEmail string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, c := range m.clients {
		if c.ID == id {
			if newName != "" {
				c.Name = newName
			}
			if newEmail != "" {
				c.Email = newEmail
			}
			return nil
		}
	}
	return ErrClientNotFound(id)
}

func (m *MemoryStorage) Delete(id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.clients {
		if c.ID == id {
			m.clients = append(m.clients[:i], m.clients[i+1:]...)
			return nil
		}
	}
	return ErrClientNotFound(id)
}
