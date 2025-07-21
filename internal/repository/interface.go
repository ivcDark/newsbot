package repository

import "github.com/ivcDark/newsbot/internal/domain"

type NewsRepository interface {
	Save(news *domain.News) error
	GetAll() ([]*domain.News, error)
	GetById(id int64) (*domain.News, error)
	ExistsByURL(url string) (bool, error)
}
