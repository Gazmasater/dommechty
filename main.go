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
		if update.Message != nil && update.Message.Text == "/start" {
			// Шапка
			headerPhoto := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL("https://gazmasater.github.io/dommechty/header.jpg"))
			headerPhoto.Caption = "Добро пожаловать в каталог домов"
			if _, err := botAPI.Send(headerPhoto); err != nil {
				log.Println("Ошибка отправки шапки:", err)
			}

			// Простая ссылка вместо WebApp-кнопки
			webAppLink := tgbotapi.NewMessage(update.Message.Chat.ID, "🌐 Витрина домов: https://gazmasater.github.io/dommechty/")
			if _, err := botAPI.Send(webAppLink); err != nil {
				log.Println("Ошибка отправки ссылки:", err)
			}

			// Список домов
			houses := []struct {
				Name, Description, PhotoURL string
			}{
				{"🏡 Дом 120 м²", "2 этажа, участок 6 соток", "https://terem-dom.ru/d/cimg6172.jpg"},
				{"🏠 Дом 95 м²", "Компактный и тёплый", "https://terem-dom.ru/d/cimg6177.jpg"},
				{"🏘 Дом с террасой", "С видом на реку", "https://terem-dom.ru/d/cimg6169.jpg"},
				{"🏕 Коттедж", "Для семьи и отдыха", "https://terem-dom.ru/d/cimg6170.jpg"},
			}

			for _, h := range houses {
				msg := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(h.PhotoURL))
				msg.Caption = h.Name + "\n" + h.Description
				if _, err := botAPI.Send(msg); err != nil {
					log.Println("Ошибка отправки дома:", err)
				}
			}
		}
	}
}
