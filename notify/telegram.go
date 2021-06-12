package notify

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Telegram struct {
	bot    *tgbotapi.BotAPI
	userId int64
}

func NewTelegram(botToken string, userId int64) *Telegram {
	if botToken == "" {
		return &Telegram{}
	}
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}
	return &Telegram{bot, userId}
}

func (t *Telegram) Send(report *Report) {
	if t.bot == nil {
		return
	}
	msgText := ""
	if report.Items == nil || len(report.Items) == 0 {
		return
	}
	for _, v := range report.Items {
		if v.Status != StatusInfo {
			msgText += string(v.Status) + " "
		}
		msgText += v.Message + "\n"
	}
	msg := tgbotapi.NewMessage(t.userId, msgText)
	_, err := t.bot.Send(msg)
	if err != nil {
		log.Print(err)
	}
}
