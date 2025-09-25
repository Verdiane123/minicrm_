package storage

import (
	"encoding/json"
	"os"
	"sync"
)


type JSONStorage struct {
	file   string
	mu     sync.Mutex
	clients []*Client
}


func NewJSONStorage(filename string) (*JSONStorage, error) {
	js := &JSONStorage{file: filename}

	
	if _, err := os.Stat(filename); err == nil {
		data, err := os.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		if len(data) > 0 {
			if err := json.Unmarshal(data, &js.clients); err != nil {
				return nil, err
			}
		}
	}

	return js, nil
}


func (js *JSONStorage) persist() error {
	data, err := json.MarshalIndent(js.clients, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(js.file, data, 0644)
}


func (js *JSONStorage) Add(client *Client) error {
	js.mu.Lock()
	defer js.mu.Unlock()

	
	if len(js.clients) > 0 {
		client.ID = js.clients[len(js.clients)-1].ID + 1
	} else {
		client.ID = 1
	}

	js.clients = append(js.clients, client)
	return js.persist()
}


func (js *JSONStorage) GetAll() ([]*Client, error) {
	js.mu.Lock()
	defer js.mu.Unlock()

	return js.clients, nil
}


func (js *JSONStorage) GetByID(id int) (*Client, error) {
	js.mu.Lock()
	defer js.mu.Unlock()

	for _, c := range js.clients {
		if c.ID == id {
			return c, nil
		}
	}
	return nil, ErrClientNotFound(id)
}


func (js *JSONStorage) Update(id int, newName, newEmail string) error {
	js.mu.Lock()
	defer js.mu.Unlock()

	for _, c := range js.clients {
		if c.ID == id {
			c.Name = newName
			c.Email = newEmail
			return js.persist()
		}
	}
	return ErrClientNotFound(id)
}


func (js *JSONStorage) Delete(id int) error {
	js.mu.Lock()
	defer js.mu.Unlock()

	for i, c := range js.clients {
		if c.ID == id {
			
			js.clients = append(js.clients[:i], js.clients[i+1:]...)
			return js.persist()
		}
	}
	return ErrClientNotFound(id)
}
