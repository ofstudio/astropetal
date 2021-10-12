package config

import (
	"astropetal/embeded"
	"crypto/tls"
	"io/fs"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Cron      string
	BaseUrl   string
	TlsCert   *tls.Certificate
	BotApiKey string
	UserId    int64
}

func MustLoad() *Config {
	return &Config{
		Cron:      mustReadCron(),
		BaseUrl:   "gemini://astrobotany.mozz.us",
		TlsCert:   mustReadTlsCert(),
		BotApiKey: mustGetBotApiKey(),
		UserId:    mustGetUserId(),
	}
}

func mustReadCron() string {
	val := os.Getenv("CRON")
	if val == "" {
		val = "0 0 * * *"
	}
	return val
}

func mustReadTlsCert() *tls.Certificate {
	certPemBlock, err := fs.ReadFile(embeded.ClientCertFS, "identity.crt")
	if err != nil {
		log.Fatal(err)
	}
	keyPemBlock, err := fs.ReadFile(embeded.ClientCertFS, "identity.key")
	if err != nil {
		log.Fatal(err)
	}
	tlsCert, err := tls.X509KeyPair(certPemBlock, keyPemBlock)
	if err != nil {
		log.Fatal(err)
	}
	return &tlsCert
}

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
