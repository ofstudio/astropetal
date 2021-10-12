package main

import (
	"astropetal/config"
	"astropetal/notify"
	"astropetal/picker"
	"astropetal/timing"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	// Schedule pickBot
	c := cron.New()
	entryId, err := c.AddFunc(cfg.Cron, func() {
		time.Sleep(timing.Approx(1*time.Hour, 1*time.Hour))
		report := pickBot.Pick()
		notificator.SendReport(report)
	})
	if err != nil {
		log.Fatal("cron error: ", err)
	}

	// Start schedule
	c.Start()
	logAndNotify(fmt.Sprintf("Cron: '%s' should start after %v", cfg.Cron, c.Entry(entryId).Next))

	// Wait for interrupt
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	logAndNotify("astropetal-bot stopped")
}
