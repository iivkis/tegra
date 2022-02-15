package app

import (
	"fmt"
	"log"
	"tegra/internal/repository"
	"tegra/pkg/tegra"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Launch() {
	log.Println("App launching..")

	bot, err := tgbotapi.NewBotAPI("your token")
	if err != nil {
		panic(err)
	}

	//create repo
	repo := repository.NewRepository("tegra.sqlite")

	//create handler with replicas
	handler := tegra.NewHandler(replicas)

	//set vars in storage
	handler.Storage.Set("bot", bot)
	handler.Storage.Set("repo", repo)

	//set handlers
	setHandlerFuncs(handler)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	log.Println("Wait messages...")
	for update := range bot.GetUpdatesChan(u) {
		user := repository.UserModel{
			TelegramID: update.SentFrom().ID,
		}
		repo.Users.Create(&user)

		if update.Message != nil && update.Message.Text != "" {
			go handler.HandleMessage(&update)
		}
	}
}

func setHandlerFuncs(h *tegra.Handler) {
	//get repo
	repo, ok := h.Storage.Get("repo").(*repository.Repository)
	if !ok {
		panic(ok)
	}

	//get bot
	bot, ok := h.Storage.Get("bot").(*tgbotapi.BotAPI)
	if !ok {
		panic(ok)
	}

	//create middleware
	withAdmin := func(upd *tegra.Update) {
		if user, _ := repo.Users.GetOne(upd.SentID); user.IsAdmin() {
			upd.Next()
		}
	}

	//set commands
	{
		h.AddCommand("^/start$", func(upd *tegra.Update) {
			bot.Send(tgbotapi.NewMessage(upd.SentID, h.Replicas.Get("cmd_start")))
		})

		h.AddCommand("^/repeat$", func(upd *tegra.Update) {
			bot.Send(tgbotapi.NewMessage(upd.SentID, "enter word:"))
			h.Actions.Set(upd.SentID, ACT_REPEAT_WORD)
		})

		h.AddCommand("^/id$", func(upd *tegra.Update) {
			msg := tgbotapi.NewMessage(upd.SentID, h.Replicas.Get("cmd_id", upd.SentID))
			msg.ParseMode = "html"
			bot.Send(msg)
		})

		h.AddCommand("^/actionsCount", withAdmin, func(upd *tegra.Update) {
			msg := fmt.Sprintf("count: %d", h.Actions.Count())
			bot.Send(tgbotapi.NewMessage(upd.SentID, msg))
		})
	}

	//set actions
	{
		h.AddAction(ACT_DEFAULT, func(upd *tegra.Update) {
			bot.Send(tgbotapi.NewMessage(upd.SentID, "i don't understand u"))
		})

		h.AddAction(ACT_REPEAT_WORD, func(upd *tegra.Update) {
			bot.Send(tgbotapi.NewMessage(upd.SentID, upd.Message.Text))
			h.Actions.Clear(upd.SentID)
		})
	}
}
