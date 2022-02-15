package tegra

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Update struct {
	*tgbotapi.Update

	SentID int64
	ChatID int64

	stopped bool
	store   map[string]interface{}
}

func newUpdate(upd *tgbotapi.Update) *Update {
	return &Update{
		Update: upd,
		SentID: upd.SentFrom().ID,
		ChatID: upd.FromChat().ID,
		store:  map[string]interface{}{},
	}
}

func (upd *Update) Next() {
	upd.stopped = false
}

func (upd *Update) stop() {
	upd.stopped = true
}

func (upd *Update) isStopped() bool {
	return upd.stopped
}

func (upd *Update) SetItem(name string, v interface{}) {
	upd.store[name] = v
}

func (upd *Update) GetItem(name string) (v interface{}) {
	return upd.store[name]
}
