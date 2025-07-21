package repository

import (
	"database/sql"
	"github.com/ivcDark/newsbot/internal/domain"
)

type SQLiteNewsRepository struct {
	db *sql.DB
}

func NewSQLiteNewsRepository(db *sql.DB) *SQLiteNewsRepository {
	return &SQLiteNewsRepository{db: db}
}

func (r *SQLiteNewsRepository) Save(news *domain.News) error {
	query := `
		INSERT INTO news (title, subtitle, url, image_url, content, published)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query,
		news.Title,
		news.Subtitle,
		news.URL,
		news.Image,
		news.Content,
		news.Published,
	)

	return err
}

func (r *SQLiteNewsRepository) GetAll() ([]*domain.News, error) {
	rows, err := r.db.Query("SELECT title, subtitle, url, image_url, content, published, created_at FROM news")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*domain.News
	for rows.Next() {
		var newsItem domain.News
		err := rows.Scan(
			&newsItem.Title,
			&newsItem.Subtitle,
			&newsItem.URL,
			&newsItem.Image,
			&newsItem.Content,
			&newsItem.Published,
			&newsItem.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, &newsItem)
	}

	return results, nil
}

func (r *SQLiteNewsRepository) GetById(id int64) (*domain.News, error) {
	query := `SELECT title, subtitle, url, image_url, content, published, created_at FROM news WHERE id=?`
	row := r.db.QueryRow(query, id)

	var newsItem domain.News
	err := row.Scan(
		&newsItem.Title,
		&newsItem.Subtitle,
		&newsItem.URL,
		&newsItem.Image,
		&newsItem.Content,
		&newsItem.Published,
		&newsItem.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &newsItem, nil
}

func (r *SQLiteNewsRepository) ExistsByURL(url string) (bool, error) {
	query := `SELECT count(*) FROM news WHERE url=?`
	var count int
	err := r.db.QueryRow(query, url).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
