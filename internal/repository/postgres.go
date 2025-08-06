package repository

import (
	"database/sql"
	"github.com/ivcDark/newsbot/internal/domain"
)

type PostgresNewsRepository struct {
	db *sql.DB
}

func NewPostgresNewsRepository(db *sql.DB) *PostgresNewsRepository {
	return &PostgresNewsRepository{db: db}
}

func (r *PostgresNewsRepository) Save(news *domain.News) error {
	query := `
		INSERT INTO news (title, subtitle, url, image_url, content, published)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (url) DO NOTHING;
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

func (r *PostgresNewsRepository) GetAll() ([]*domain.News, error) {
	rows, err := r.db.Query("SELECT id, title, subtitle, url, image_url, content, published, created_at FROM news")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*domain.News
	for rows.Next() {
		var newsItem domain.News
		err := rows.Scan(
			&newsItem.ID,
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

func (r *PostgresNewsRepository) GetById(id int64) (*domain.News, error) {
	query := `SELECT id, title, subtitle, url, image_url, content, published, created_at FROM news WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var newsItem domain.News
	err := row.Scan(
		&newsItem.ID,
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

func (r *PostgresNewsRepository) ExistsByURL(url string) (bool, error) {
	query := `SELECT COUNT(*) FROM news WHERE url = $1`
	var count int
	err := r.db.QueryRow(query, url).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *PostgresNewsRepository) GetUnpublished() ([]*domain.News, error) {
	rows, err := r.db.Query("SELECT id, title, subtitle, url, image_url, content, published, created_at FROM news WHERE is_published = false")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*domain.News
	for rows.Next() {
		var newsItem domain.News
		err := rows.Scan(
			&newsItem.ID,
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

func (r *PostgresNewsRepository) MarkAsPublished(id int64) error {
	_, err := r.db.Exec("UPDATE news SET is_published = true WHERE id = $1", id)
	return err
}
