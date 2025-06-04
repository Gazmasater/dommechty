package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_TOKEN не задан в .env")
	}

	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := botAPI.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil || update.Message.Text != "/start" {
			continue
		}

		// 1. Шапка (фото + подпись)
		header := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL("https://gazmasater.github.io/dommechty/header.jpg"))
		header.Caption = "🏡 Добро пожаловать в каталог домов"
		if _, err := botAPI.Send(header); err != nil {
			log.Println("Ошибка отправки шапки:", err)
		}

		// 2. Кнопка для перехода на Web App
		button := tgbotapi.NewMessage(update.Message.Chat.ID, "Нажмите кнопку ниже, чтобы открыть витрину:")
		button.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("🌐 Открыть витрину", "https://gazmasater.github.io/dommechty/"),
			),
		)

		if _, err := botAPI.Send(button); err != nil {
			log.Println("Ошибка отправки кнопки:", err)
		}
	}
}
