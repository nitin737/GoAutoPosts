package store

import "github.com/nitin737/GoAutoPosts/go-daily/internal/model"

// Repository defines the interface for storing posted library history
type Repository interface {
	// Save saves a posted library record
	Save(posted *model.PostedLibrary) error

	// GetAll retrieves all posted library records
	GetAll() ([]model.PostedLibrary, error)

	// GetByName retrieves a posted library by name
	GetByName(name string) (*model.PostedLibrary, error)
}
