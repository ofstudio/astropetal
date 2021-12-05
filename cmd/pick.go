package main

import (
	"astropetal/config"
	"astropetal/notify"
	"astropetal/picker"
	"log"
)

func main() {
	cfg := config.MustLoad()
	notificator := notify.NewTelegram(cfg.BotApiKey, cfg.UserId)
	pickBot := picker.NewPicker(cfg.BaseUrl, cfg.TlsCert)
	logAndNotify := func(msg string) {
		log.Print(msg)
		notificator.Send(msg)
	}

	logAndNotify("astropetal-bot started")
	report := pickBot.Pick()
	notificator.SendReport(report)

	logAndNotify("astropetal-bot stopped")
}
