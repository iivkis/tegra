package tegra

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Update struct {
	*tgbotapi.Update

	SentID int64
	ChatID int64

	*Storage

	stopped bool
}

func NewUpdate(upd *tgbotapi.Update) *Update {
	return &Update{
		Update:  upd,
		SentID:  upd.SentFrom().ID,
		ChatID:  upd.FromChat().ID,
		Storage: NewStorage(),
	}
}

func (upd *Update) Next() {
	upd.stopped = false
}

func (upd *Update) stop() {
	upd.stopped = true
}

func (upd *Update) Stopped() bool {
	return upd.stopped
}
