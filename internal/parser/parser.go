package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/ivcDark/newsbot/internal/domain"
	"log"
	"net/http"
	"time"
)

type NewsItem struct {
	Title    string
	Subtitle string
	Link     string
	Image    string
	Date     string
	Content  string
}

func ParseNews() ([]NewsItem, error) {
	url := "https://63.ru/text/"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Request error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Bad status: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Html parse error: %s", err)
	}

	var news []NewsItem

	doc.Find("div.wrap_RL97A").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a.header_RL97A").Text()
		link, _ := s.Find("a.header_RL97A").Attr("href")
		subtitle := s.Find("span.subtitle_RL97A").Text()
		imgSrc, _ := s.Find("img.image_RL97A").Attr("src")
		date := s.Find("div.statistic_RL97A span.text_eiDCU").First().Text()
		content := getArticleText(link)

		news = append(news, NewsItem{
			Title:    title,
			Subtitle: subtitle,
			Link:     link,
			Image:    imgSrc,
			Date:     date,
			Content:  content,
		})
	})

	return news, nil
}

func getArticleText(link string) string {
	resp, err := http.Get(link)
	if err != nil {
		log.Printf("Ошибка загрузки статьи %s: %v", link, err)
		return ""
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Ошибка: статья %s вернула статус %d", link, resp.StatusCode)
		return ""
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Ошибка парсинга html статьи: %v", err)
		return ""
	}

	content := ""

	doc.Find("div.uiArticleBlockText_5xJo1").Each(func(i int, s *goquery.Selection) {
		content += s.Text()
	})

	return content
}

func (ni NewsItem) ToDomain() *domain.News {
	publishedTime, _ := time.Parse("02 января 2006", ni.Date) // можешь поменять формат если сайт другой
	return &domain.News{
		Title:     ni.Title,
		Subtitle:  ni.Subtitle,
		URL:       ni.Link,
		Image:     ni.Image,
		Content:   ni.Content,
		Published: publishedTime,
	}
}
