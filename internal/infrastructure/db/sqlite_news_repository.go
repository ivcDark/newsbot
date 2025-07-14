package db

import (
	"database/sql"
	"github.com/ivcDark/newsbot/internal/domain"
	"time"
)

type SQLiteNewsRepository struct {
	db *sql.DB
}

func (r *SQLiteNewsRepository) Save(news *domain.News) error {
	query := `
		INSERT INTO news (title, subtitle, url, image_url, content, published, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query, news.Title, news.Subtitle, news.URL, news.Image, news.Content, news.Published, time.Now())

	return err
}

func (r *SQLiteNewsRepository) FindByURL(url string) (*domain.News, error) {
	query := `SELECT id, title, subtitle, url, image_url, content, published, created_at FROM news WHERE url = ?`
	row := r.db.QueryRow(query, url)

	var news domain.News
	err := row.Scan(
		&news.ID,
		&news.Title,
		&news.Subtitle,
		&news.URL,
		&news.Image,
		&news.Content,
		&news.Published,
		&news.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &news, nil
}

func (r *SQLiteNewsRepository) FindByID(id int64) (*domain.News, error) {
	query := `SELECT id, title, subtitle, url, image_url, content, published, created_at FROM news WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var news domain.News
	err := row.Scan(
		&news.ID,
		&news.Title,
		&news.Subtitle,
		&news.URL,
		&news.Image,
		&news.Content,
		&news.Published,
		&news.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &news, nil
}

func (r *SQLiteNewsRepository) GetLatest(limit int) ([]domain.News, error) {
	query := `SELECT id, title, subtitle, url, image_url, content, published, created_at FROM news ORDER BY published DESC LIMIT ?`
	rows, err := r.db.Query(query, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var newsList []domain.News

	for rows.Next() {
		var news domain.News
		err := rows.Scan(
			&news.ID,
			&news.Title,
			&news.Subtitle,
			&news.URL,
			&news.Image,
			&news.Content,
			&news.Published,
			&news.CreatedAt,
		)

		if err != nil {
			return nil, err
		}
		newsList = append(newsList, news)
	}

	return newsList, nil
}
