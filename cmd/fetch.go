package cmd

import (
	"database/sql"
	"fmt"
	"github.com/ivcDark/newsbot/internal/repository"
	"github.com/ivcDark/newsbot/internal/service"
	"github.com/spf13/cobra"
	"log"
)

var sourceURL string

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Получить и сохранить новости",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite3", "./newsbot.db")
		if err != nil {
			log.Fatalf("Failed to open db: %v", err)
		}
		defer db.Close()

		if sourceURL == "" {
			sourceURL = "https://63.ru/text/"
		}

		newsRepo := repository.NewSQLiteNewsRepository(db)
		newsService := service.NewNewsService(newsRepo)

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
