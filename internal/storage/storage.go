package storage

import (
	"fmt"
)


type Client struct {
	ID    int    `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"size:100;not null"`
	Email string `gorm:"size:100;unique;not null"`
}


type Storer interface {
	Add(client *Client) error
	GetAll() ([]*Client, error)
	GetByID(id int) (*Client, error)
	Update(id int, newName, newEmail string) error
	Delete(id int) error
}

var ErrClientNotFound = func(id int) error {
	return fmt.Errorf("Client avec l'ID %d non trouvé", id)
}
