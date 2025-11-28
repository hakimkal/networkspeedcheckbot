package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

//	func CmdKeyboard() tgbotapi.ReplyKeyboardMarkup {
//		var cmdKeyboard = tgbotapi.NewReplyKeyboard(
//			tgbotapi.NewKeyboardButtonRow(
//				tgbotapi.NewKeyboardButton("/speed-check"),
//			),
//			tgbotapi.NewKeyboardButtonRow(
//				tgbotapi.NewKeyboardButton("/download-speed-check"),
//			),
//			tgbotapi.NewKeyboardButtonRow(
//				tgbotapi.NewKeyboardButton("/upload-speed-check"),
//			),
//		)
//		return cmdKeyboard
//	}
func CmdKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/speedCheck"),
			tgbotapi.NewKeyboardButton("/downloadSpeedCheck"),
			tgbotapi.NewKeyboardButton("/uploadSpeedCheck"),
		),
	)

	keyboard.ResizeKeyboard = true   // fit screen width
	keyboard.OneTimeKeyboard = false // keep it visible after use

	return keyboard
}
