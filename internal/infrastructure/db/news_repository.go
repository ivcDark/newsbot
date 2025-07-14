package db

import (
	"github.com/ivcDark/newsbot/internal/domain"
	"gorm.io/gorm"
)

type GormNewsRepository struct {
	db *gorm.DB
}

func NewGormNewsRepository(db *gorm.DB) *GormNewsRepository {
	return &GormNewsRepository{db: db}
}

func (r *GormNewsRepository) Save(news *domain.News) error {
	return r.db.Save(news).Error
}

func (r *GormNewsRepository) FindByURL(url string) (*domain.News, error) {
	var news domain.News
	err := r.db.Where("url = ?", url).First(&news).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &news, err
}
