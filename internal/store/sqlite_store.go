package store

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nitin737/GoAutoPosts/internal/model"
)

// SQLiteStore implements Repository using SQLite
type SQLiteStore struct {
	db *sql.DB
}

// NewSQLiteStore creates a new SQLite-based store
func NewSQLiteStore(dbPath string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	store := &SQLiteStore{db: db}
	if err := store.initialize(); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *SQLiteStore) initialize() error {
	query := `
	CREATE TABLE IF NOT EXISTS posted_libraries (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		library_data TEXT NOT NULL,
		posted_at DATETIME NOT NULL,
		post_id TEXT,
		image_path TEXT
	);
	CREATE INDEX IF NOT EXISTS idx_posted_at ON posted_libraries(posted_at);
	CREATE INDEX IF NOT EXISTS idx_name ON posted_libraries(name);
	`

	_, err := s.db.Exec(query)
	return err
}

// Save saves a posted library record
func (s *SQLiteStore) Save(posted *model.PostedLibrary) error {
	libraryData, err := json.Marshal(posted.Library)
	if err != nil {
		return err
	}

	query := `
	INSERT INTO posted_libraries (name, library_data, posted_at, post_id, image_path)
	VALUES (?, ?, ?, ?, ?)
	`

	_, err = s.db.Exec(query,
		posted.Library.Name,
		string(libraryData),
		posted.PostedAt,
		posted.PostID,
		posted.ImagePath,
	)

	return err
}

// GetAll retrieves all posted library records
func (s *SQLiteStore) GetAll() ([]model.PostedLibrary, error) {
	query := `SELECT library_data, posted_at, post_id, image_path FROM posted_libraries ORDER BY posted_at DESC`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []model.PostedLibrary
	for rows.Next() {
		var libraryData string
		var posted model.PostedLibrary

		if err := rows.Scan(&libraryData, &posted.PostedAt, &posted.PostID, &posted.ImagePath); err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(libraryData), &posted.Library); err != nil {
			return nil, err
		}

		records = append(records, posted)
	}

	return records, rows.Err()
}

// GetByName retrieves a posted library by name
func (s *SQLiteStore) GetByName(name string) (*model.PostedLibrary, error) {
	query := `SELECT library_data, posted_at, post_id, image_path FROM posted_libraries WHERE name = ?`

	var libraryData string
	var posted model.PostedLibrary

	err := s.db.QueryRow(query, name).Scan(&libraryData, &posted.PostedAt, &posted.PostID, &posted.ImagePath)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("library not found: %s", name)
		}
		return nil, err
	}

	if err := json.Unmarshal([]byte(libraryData), &posted.Library); err != nil {
		return nil, err
	}

	return &posted, nil
}

// Close closes the database connection
func (s *SQLiteStore) Close() error {
	return s.db.Close()
}
