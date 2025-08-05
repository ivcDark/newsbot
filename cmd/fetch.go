package cmd

import (
	"database/sql"
	"fmt"
	"github.com/ivcDark/newsbot/internal/repository"
	"github.com/ivcDark/newsbot/internal/service"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var sourceURL string

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Получить и сохранить новости",
	Run: func(cmd *cobra.Command, args []string) {
		driver := os.Getenv("DB_DRIVER")
		dsn := os.Getenv("DB_DSN")

		if driver == "" || dsn == "" {
			log.Fatal("DB_DRIVER и DB_DSN должны быть установлены в переменных окружения")
		}

		db, err := sql.Open(driver, dsn)
		if err != nil {
			log.Fatalf("Ошибка при подключении к БД: %v", err)
		}
		defer db.Close()

		newsRepo, err := repository.NewRepository(driver, db)
		if err != nil {
			log.Fatalf("Ошибка инициализации репозитория: %v", err)
		}

		newsService := service.NewNewsService(newsRepo)

		if sourceURL == "" {
			sourceURL = "https://63.ru/text/"
		}

		err = newsService.FetchAndSaveNews(sourceURL)
		if err != nil {
			log.Fatalf("Ошибка при получении и сохранении новостей: %v", err)
		}

		fmt.Println("Новости успешно получены и сохранены")
	},
}

func init() {
	fetchCmd.Flags().StringVarP(&sourceURL, "source", "s", "", "URL источника новостей")
	rootCmd.AddCommand(fetchCmd)
}
