package cmd

import (
	"database/sql"
	"github.com/ivcDark/newsbot/internal/repository"
	"github.com/ivcDark/newsbot/internal/telegram"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"time"
)

var tgPublishCmd = &cobra.Command{
	Use:   "tg_publish",
	Short: "Публикация новостей в Telegram",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite3", "./newsbot.db")
		if err != nil {
			log.Fatalf("Ошибка при подключении БД: %v", err)
		}
		defer db.Close()

		newsRepo := repository.NewSQLiteNewsRepository(db)

		idStr, _ := cmd.Flags().GetString("id")

		if idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Fatalf("Неверный id: %v", err)
			}

			news, err := newsRepo.GetById(int64(id))
			if err != nil {
				log.Fatalf("Ошибка получения новости: %v", err)
			}
			if news == nil {
				log.Printf("Новость с id %d не найдена", id)
				return
			}

			err = telegram.PublishToChannel(news)
			if err != nil {
				log.Fatalf("Ошибка публикации: %v", err)
			}

			newsRepo.MarkAsPublished(news.ID)
			log.Printf("✅ Опубликована новость: %s", news.Title)
		} else {
			newsList, err := newsRepo.GetUnpublished()
			if err != nil {
				log.Fatalf("Ошибка получения новостей: %v", err)
			}

			for _, news := range newsList {
				err := telegram.PublishToChannel(news)
				if err != nil {
					log.Printf("Ошибка публикации: %v", err)
					continue
				}
				newsRepo.MarkAsPublished(news.ID)
				log.Printf("✅ Опубликована новость: %s", news.Title)
				time.Sleep(5 * time.Second)
			}
		}
	},
}

func init() {
	tgPublishCmd.Flags().String("id", "", "ID новости для публикации")
	rootCmd.AddCommand(tgPublishCmd)
}
