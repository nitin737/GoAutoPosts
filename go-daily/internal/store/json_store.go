package store

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/nitin737/GoAutoPosts/go-daily/internal/model"
)

// JSONStore implements Repository using a JSON file
type JSONStore struct {
	filePath string
	mu       sync.RWMutex
}

// NewJSONStore creates a new JSON-based store
func NewJSONStore(filePath string) *JSONStore {
	return &JSONStore{
		filePath: filePath,
	}
}

// Save saves a posted library record
func (s *JSONStore) Save(posted *model.PostedLibrary) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Load existing records
	records, err := s.loadRecords()
	if err != nil {
		return err
	}

	// Append new record
	records = append(records, *posted)

	// Save back to file
	return s.saveRecords(records)
}

// GetAll retrieves all posted library records
func (s *JSONStore) GetAll() ([]model.PostedLibrary, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.loadRecords()
}

// GetByName retrieves a posted library by name
func (s *JSONStore) GetByName(name string) (*model.PostedLibrary, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	records, err := s.loadRecords()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		if record.Library.Name == name {
			return &record, nil
		}
	}

	return nil, fmt.Errorf("library not found: %s", name)
}

func (s *JSONStore) loadRecords() ([]model.PostedLibrary, error) {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.PostedLibrary{}, nil
		}
		return nil, err
	}

	var records []model.PostedLibrary
	if err := json.Unmarshal(data, &records); err != nil {
		return nil, err
	}

	return records, nil
}

func (s *JSONStore) saveRecords(records []model.PostedLibrary) error {
	data, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}
