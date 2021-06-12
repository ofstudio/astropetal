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

func (t *Telegram) Send(text string) {
	if t.bot == nil {
		return
	}
	msg := tgbotapi.NewMessage(t.userId, text)
	_, err := t.bot.Send(msg)
	if err != nil {
		log.Print(err)
	}
}

func (t *Telegram) SendReport(report *Report) {
	text := ""
	if report.Items == nil || len(report.Items) == 0 {
		return
	}
	for _, v := range report.Items {
		if v.Status != StatusInfo {
			text += string(v.Status) + " "
		}
		text += v.Message + "\n"
	}
	t.Send(text)
}
