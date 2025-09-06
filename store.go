package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Note struct {
	ID    int64
	Title string
	Body  string
}

type Store struct {
	conn *sql.DB
}

func (s *Store) Init() error {
	var err error

	s.conn, err = sql.Open("sqlite3", "../notes.db")

	if err != nil {
		return err
	}

	createTableStmt := `
  CREATE TABLE IF NOT EXISTS notes (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    body TEXT NOT NULL
  );
  `
	_, err = s.conn.Exec(createTableStmt)

	return err
}

func (s *Store) GetAllNotes() ([]Note, error) {
	rows, err := s.conn.Query("SELECT id, title, body FROM notes")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	notes := []Note{}

	for rows.Next() {
		var note Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Body); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (s *Store) AddNote(note Note) error {
	if note.ID == 0 {
		note.ID = time.Now().UTC().UnixNano()
	}

	upsertQuery := `
  INSERT INTO notes (id, title, body) VALUES (?, ?, ?)
  ON CONFLICT(id) DO UPDATE SET title = excluded.title, body = excluded.body;
  `
	_, err := s.conn.Exec(upsertQuery, note.ID, note.Title, note.Body)

	return err
}
