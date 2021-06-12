package config

import (
	"log"
	"os"
	"strconv"
)

func mustGetBotApiKey() string {
	return os.Getenv("TELEGRAM_API_KEY")
}

func mustGetUserId() int64 {
	val := os.Getenv("TELEGRAM_USER_ID")
	if val == "" {
		return 0
	}
	id, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Fatal("invalid TELEGRAM_USER_ID")
	}
	return id
}
