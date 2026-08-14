package notes

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Note struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type Store struct {
	pool *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{pool: pool}
}

func (s *Store) List(ctx context.Context) ([]Note, error) {
	rows, err := s.pool.Query(ctx, `
		SELECT id, text, created_at
		FROM notes
		ORDER BY id
	`)
	if err != nil {
		return nil, fmt.Errorf("list notes: %w", err)
	}
	defer rows.Close()

	notes := make([]Note, 0)
	for rows.Next() {
		var n Note
		if err := rows.Scan(&n.ID, &n.Text, &n.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan note: %w", err)
		}
		notes = append(notes, n)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate notes: %w", err)
	}
	return notes, nil
}

func (s *Store) Create(ctx context.Context, text string) (Note, error) {
	var n Note
	err := s.pool.QueryRow(ctx, `
		INSERT INTO notes (text)
		VALUES ($1)
		RETURNING id, text, created_at
	`, text).Scan(&n.ID, &n.Text, &n.CreatedAt)
	if err != nil {
		return Note{}, fmt.Errorf("create note: %w", err)
	}
	return n, nil
}
