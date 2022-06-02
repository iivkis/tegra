package tegra

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *TegraHandler) HandleMessage(update *tgbotapi.Update) {
	log.Printf("MESSAGE: @%s | `%s`", update.SentFrom().UserName, update.Message.Text)

	var upd = NewUpdate(update)
	if upd.Message.IsCommand() {
		h.ActionStorage.Clear(upd.SentID)
		h.HandleCommand(upd)
	} else {
		h.HandleAction(upd)
	}
}

func (h *TegraHandler) HandleCommand(upd *Update) {
	for _, handler := range h.allHandlers {
		if ok := handler.Pattern.MatchString(upd.Message.Text); ok {
			h.ExecFuncs(upd, handler.Funcs)
			return
		}
	}
	h.UseNotFound(upd)
}

func (h *TegraHandler) HandleAction(upd *Update) {
	funcs := h.allActions[h.ActionStorage.Get(upd.SentID)]
	h.ExecFuncs(upd, funcs)
}

func (h *TegraHandler) HandleCallback(update *tgbotapi.Update) {
	var upd = NewUpdate(update)
	for _, callback := range h.allCallbacks {
		if ok := callback.Pattern.MatchString(upd.CallbackData()); ok {
			h.ExecFuncs(upd, callback.Funcs)
			return
		}
	}
}
