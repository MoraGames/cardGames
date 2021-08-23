package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func standardButtonsMarkup() tgbotapi.InlineKeyboardMarkup {
	upVoteButton := tgbotapi.NewInlineKeyboardButtonData("🔼 Up Vote", "[DATA] Upvoted message.")
	downVoteButton := tgbotapi.NewInlineKeyboardButtonData("🔽 Down Vote", "[DATA] Downvoted message.")
	reportButton := tgbotapi.NewInlineKeyboardButtonData("❗️ Report", "[DATA] Reported message.")
	replyButton := tgbotapi.NewInlineKeyboardButtonData("💬 Reply", "[DATA] Wants to reply to a message.")

	firstButtonsRow := tgbotapi.NewInlineKeyboardRow(upVoteButton, downVoteButton)
	secondButtonsRow := tgbotapi.NewInlineKeyboardRow(replyButton)
	thirdButtonsRow := tgbotapi.NewInlineKeyboardRow(reportButton)

	return tgbotapi.NewInlineKeyboardMarkup(firstButtonsRow, secondButtonsRow, thirdButtonsRow)
}

// REQUIRED Google Translate APIs Plan (45.00 USD/hour)
// func SupportedLanguages() {
// 	ctx := context.Background()
// 	client, err := translate.NewClient(ctx)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	langs, err := client.SupportedLanguages(ctx, language.English)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	log.Println(langs)
// }
