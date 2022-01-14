package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, year, release_date from movies where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie

	err := row.Scan(&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate)

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (m *DBModel) All(id int) ([]*Movie, error) {
	return nil, nil
}
