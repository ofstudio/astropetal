package main

import (
	"astropetal/config"
	"astropetal/notify"
	"astropetal/picker"
	"astropetal/timing"
	"log"
	"time"
)

const (
	msgStart = "astropetal-bot started"
)

func main() {
	cfg := config.MustLoad()
	notificator := notify.NewTelegram(cfg.BotApiKey, cfg.UserId)
	bot := picker.NewPicker(cfg.BaseUrl, cfg.TlsCert)

	log.Print("* * *")
	log.Print(msgStart)
	notificator.Send(notify.NewReportSingle(notify.StatusOk, msgStart))

	for ; ; {
		report := bot.Pick()
		notificator.Send(report)
		time.Sleep(timing.Approx(25*time.Hour, 20*time.Minute))
	}
}
