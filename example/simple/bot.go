package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/iivkis/tegra"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5320362080:AAEsjlreUe9Qwmm2UzlF2e2s3II_cFbrsGQ")
	if err != nil {
		panic(err)
	}

	handler := tegra.NewTegraHandler(&tegra.TegraHandlerConfig{
		Replicas: map[string]string{},
	})

	handler.Local.Set("bot", bot)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	handler.AddCommand("start", func(upd *tegra.Update) {
		msg := tgbotapi.NewMessage(upd.SentID, "hi!")
		bot.Send(msg)
	})

	for update := range updates {
		if update.Message != nil {
			go handler.HandleMessage(&update)
		}
	}
}
