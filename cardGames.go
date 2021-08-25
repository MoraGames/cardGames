package main

import (
	"log"
	"strings"

	"github.com/MoraGames/cardGames/pkg/internal/languages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	//Default Settings
	//isSup := languages.isSupported("English")
	defaultLanguage := languages.Language("English")

	//Get the API_Token
	botToken, err := getToken("@MG_Telegram_bot")
	if err != nil {
		log.Panic(err)
	}
	log.Print(">> Token:", botToken, "\n\n")

	//Get the botAPI object
	botAPI, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}
	botAPI.Debug = true

	//Set the eventUpdate
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60
	updatesChannel, err := botAPI.GetUpdatesChan(update)
	if err != nil {
		log.Panic(err)
	}

	//Run the bot
	for update := range updatesChannel {
		if update.Message == nil {
			if update.CallbackQuery == nil {
				//Ignore any non-Message or non-CallbackQuery Updates
				continue
			}
			//manageCallbackQuery(update.CallbackQuery)
			continue
		}

		response := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		response.ReplyToMessageID = update.Message.MessageID

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				//Check if this is the first time this user has used /start
				_, userAlredyExist := usersSettings[user(update.Message.From.UserName)]
				if !userAlredyExist {
					//Create a new userSettings
					usersSettings[user(update.Message.From.UserName)] = settings{language: defaultLanguage}
				}

				//Get the language set for the user and write the welcome message
				settings := usersSettings[user(update.Message.From.UserName)]
				msg := messages["Welcoming"]
				if isSupported(settings.language) {
					response.Text = strings.Join(msg[settings.language], "")
				} else {
					//TODO: Implemets Logrus to make a log file and report the error "Language not supported".
					response.Text = strings.Join(msg[defaultLanguage], "")
				}
			}
		}

		//response.BaseChat.ReplyMarkup = standardButtonsMarkup()

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID

		botAPI.Send(response)
	}
}
