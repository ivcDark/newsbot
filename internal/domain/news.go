package domain

import "time"

type News struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	URL       string    `gorm:"unique" json:"url"`
	Image     string    `gorm:"column:image_url"`
	Content   string    `json:"content"`
	Published time.Time `json:"published"`
	CreatedAt time.Time `json:"created_at"`
}
