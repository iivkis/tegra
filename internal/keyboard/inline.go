package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func InlineStart() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Okey!", "/okey"),
		),
	)
}
