package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ivcDark/newsbot/internal/domain"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var bot *tgbotapi.BotAPI
var chatID string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env файл не найден, продолжаем без него")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN не найден")
	}

	chatID = os.Getenv("TELEGRAM_CHAT_ID")
	if chatID == "" {
		log.Fatal("TELEGRAM_CHAT_ID не найден")
	}

	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Ошибка инициализации бота: %v", err)
	}
}

func PublishToChannel(news *domain.News) error {
	message := fmt.Sprintf("*%s*\n\n%s\n\n[Читать полностью](%s)", news.Title, news.Content, news.URL)

	msg := tgbotapi.NewMessageToChannel(chatID, message)
	msg.ParseMode = "Markdown"

	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("ошибка отправки в Telegram: %w", err)
	}

	return nil
}
