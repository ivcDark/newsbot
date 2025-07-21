package service

import (
	"github.com/ivcDark/newsbot/internal/parser"
	"github.com/ivcDark/newsbot/internal/repository"
	"log"
)

type NewsService struct {
	repo repository.NewsRepository
}

func NewNewsService(repo repository.NewsRepository) *NewsService {
	return &NewsService{repo: repo}
}

func (s *NewsService) FetchAndSaveNews(sourceURL string) error {
	newsItems, err := parser.FetchHeadlines(sourceURL)
	if err != nil {
		return err
	}

	for _, newsItem := range newsItems {
		domainNews, err := newsItem.ToDomain()
		if err != nil {
			log.Printf("Ошибка преобразования новости: %v", err)
			continue
		}
		exists, err := s.repo.ExistsByURL(domainNews.URL)
		if err != nil {
			log.Printf("Ошибка проверки новости по URL: %v", err)
			return err
		}
		if exists {
			log.Printf("Новость уже существует: %s", domainNews.URL)
			continue
		}

		err = s.repo.Save(domainNews)
		if err != nil {
			log.Printf("Ошибка сохранения новости: %v", err)
			return err
		} else {
			log.Printf("Сохранена новость: %s", domainNews.Title)
		}
	}

	return nil
}
