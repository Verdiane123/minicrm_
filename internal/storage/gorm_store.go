package storage

import (
	"fmt"
    "gorm.io/gorm"
    "github.com/glebarez/sqlite"
)


type GORMStore struct {
	db *gorm.DB
}


func NewGORMStore(dbPath string) (*GORMStore, error) {
    db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("erreur ouverture DB: %w", err)
    }

    if err := db.AutoMigrate(&Client{}); err != nil {
        return nil, fmt.Errorf("erreur migration: %w", err)
    }

    return &GORMStore{db: db}, nil
}


// Implémentations

func (s *GORMStore) Add(client *Client) error {
	return s.db.Create(client).Error
}

func (s *GORMStore) GetAll() ([]*Client, error) {
	var clients []*Client
	if err := s.db.Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

func (s *GORMStore) GetByID(id int) (*Client, error) {
	var client Client
	if err := s.db.First(&client, id).Error; err != nil {
		return nil, ErrClientNotFound(id)
	}
	return &client, nil
}

func (s *GORMStore) Update(id int, newName, newEmail string) error {
	var client Client
	if err := s.db.First(&client, id).Error; err != nil {
		return ErrClientNotFound(id)
	}

	if newName != "" {
		client.Name = newName
	}
	if newEmail != "" {
		client.Email = newEmail
	}

	return s.db.Save(&client).Error
}

func (s *GORMStore) Delete(id int) error {
	if err := s.db.Delete(&Client{}, id).Error; err != nil {
		return ErrClientNotFound(id)
	}
	return nil
}
