package selector

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/nitin737/GoAutoPosts/go-daily/internal/model"
)

// LibrarySelector handles selection of libraries
type LibrarySelector struct {
	librariesPath string
	postedPath    string
	rand          *rand.Rand
}

// NewLibrarySelector creates a new library selector
func NewLibrarySelector(librariesPath, postedPath string) *LibrarySelector {
	return &LibrarySelector{
		librariesPath: librariesPath,
		postedPath:    postedPath,
		rand:          rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// SelectRandom selects a random library that hasn't been posted recently
func (s *LibrarySelector) SelectRandom() (*model.Library, error) {
	// Load all libraries
	libraries, err := s.loadLibraries()
	if err != nil {
		return nil, fmt.Errorf("failed to load libraries: %w", err)
	}

	// Load posted history
	posted, err := s.loadPostedHistory()
	if err != nil {
		return nil, fmt.Errorf("failed to load posted history: %w", err)
	}

	// Filter out recently posted libraries
	available := s.filterAvailable(libraries, posted)
	if len(available) == 0 {
		return nil, fmt.Errorf("no available libraries to post")
	}

	// Select random library
	selected := available[s.rand.Intn(len(available))]
	return &selected, nil
}

func (s *LibrarySelector) loadLibraries() ([]model.Library, error) {
	data, err := os.ReadFile(s.librariesPath)
	if err != nil {
		return nil, err
	}

	var libraries []model.Library
	if err := json.Unmarshal(data, &libraries); err != nil {
		return nil, err
	}

	return libraries, nil
}

func (s *LibrarySelector) loadPostedHistory() ([]model.PostedLibrary, error) {
	data, err := os.ReadFile(s.postedPath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.PostedLibrary{}, nil
		}
		return nil, err
	}

	var posted []model.PostedLibrary
	if err := json.Unmarshal(data, &posted); err != nil {
		return nil, err
	}

	return posted, nil
}

func (s *LibrarySelector) filterAvailable(libraries []model.Library, posted []model.PostedLibrary) []model.Library {
	// Create a map of posted library names
	postedMap := make(map[string]time.Time)
	for _, p := range posted {
		postedMap[p.Library.Name] = p.PostedAt
	}

	// Filter libraries that haven't been posted in the last 30 days
	cutoff := time.Now().AddDate(0, 0, -30)
	var available []model.Library

	for _, lib := range libraries {
		if postedAt, exists := postedMap[lib.Name]; !exists || postedAt.Before(cutoff) {
			available = append(available, lib)
		}
	}

	return available
}
