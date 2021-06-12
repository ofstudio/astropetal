package config

import "crypto/tls"

type Config struct {
	BaseUrl   string
	TlsCert   *tls.Certificate
	BotApiKey string
	UserId    int64
}

func MustLoad() *Config {
	return &Config{
		BaseUrl:   "gemini://astrobotany.mozz.us",
		TlsCert:   mustReadTlsCert(),
		BotApiKey: mustGetBotApiKey(),
		UserId:    mustGetUserId(),
	}
}
