package mysql

import (
	models "allcran_wsx/gameplatform/pkg"
	"database/sql"
	"errors"
)

type GameModel struct {
	DB *sql.DB
}

func (m *GameModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
			 VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *GameModel) Get(id int) (*models.Game, error) {
	stmt := `SELECT id, title, description, src, created FROM games
			 WHERE id = ?`
	s := &models.Game{}

	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Description, &s.Src, &s.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *GameModel) Lastest() ([]*models.Game, error) {
	stmt := `SELECT id, title, description, src, created FROM games
			 ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var snippets []*models.Game

	for rows.Next() {
		s := &models.Game{}
		err = rows.Scan(&s.ID, &s.Title, &s.Description, &s.Src, &s.Created)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
