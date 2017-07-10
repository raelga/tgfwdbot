package main

import (
	"log"
	"os"
	"strconv"

	"strings"

	tg "gopkg.in/telegram-bot-api.v4"
)

var fwdGroupIP int64 = 430838894
var privateChats []int64

// AppendIfMissing appends an element to slice if the newElement
// doesn't already exists, otherwise, returns slice unmodified
func AppendIfMissing(slice []int64, newElement int64) []int64 {

	for _, element := range slice {
		if element == newElement {
			return slice
		}
	}
	return append(slice, newElement)
}

// RemoveIfExisting removes an element from slice if the newElement
// already exists, otherwise, returns slice unmodified
func RemoveIfExisting(slice []int64, newElement int64) []int64 {

	for index, element := range slice {
		if element == newElement {
			return append(slice[:index], slice[index+1:]...)
		}
	}
	return slice
}

func main() {

	bot, err := tg.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Starting %s", bot.Self.UserName)

	updateCfg := tg.NewUpdate(0)
	updateCfg.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateCfg)

	fwdGroupIP, err = strconv.ParseInt(os.Getenv("TELEGRAM_GROUP_ID"), 10, 64)

	if err != nil {
		log.Panic(err)
	}

	for update := range updates {

		msg := update.Message

		if msg == nil {
			continue
		}

		if msg.Chat.Type == "private" {

			privateMessageHandler(bot, update.Message)

		} else {

			for _, privateChat := range privateChats {
				bot.Send(tg.NewForward(privateChat, msg.Chat.ID, msg.MessageID))
			}

		}
	}
}

func privateMessageHandler(bot *tg.BotAPI, msg *tg.Message) {

	if strings.HasPrefix("/start", msg.Text) {

		privateChats = AppendIfMissing(privateChats, msg.Chat.ID)

		bot.Send(tg.NewMessage(fwdGroupIP, msg.From.UserName+" started following the group in a private chat."))

	} else if strings.HasPrefix("/stop", msg.Text) {

		privateChats = RemoveIfExisting(privateChats, msg.Chat.ID)

		bot.Send(tg.NewMessage(fwdGroupIP, msg.From.UserName+" stopped following the group in private chat."))

	} else {
		_, err := bot.Send(tg.NewForward(fwdGroupIP, msg.Chat.ID, msg.MessageID))

		if err != nil {
			bot.Send(tg.NewMessage(msg.Chat.ID, "Unable to FWD the message"))
		}
	}

}
