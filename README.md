# Astrobotany petal picker bot

Picks up all available petals in the [Astrobotany Community Garden](gemini://astrobotany.mozz.us/app/visit)
approximately every 25 hours.

## Requirements

- Go 1.16

## Usage

### Client certificate
Place your Astrobotany client certificate files under `embeded/client-cert` 
as `identity.crt` and `identity.key`.

## Run
`go run ./cmd/astropetal-bot.go`

## Telegram notifications

Set `TELEGRAM_API_KEY` and `TELEGRAM_USER_ID` environment variables
to enable notifications via [Telegram Bot API](https://core.telegram.org/bots/api).

## Credits

- Oleg Fomin <[ofstudio@gmail.com](mailto:ofstudio@gmail.com)>
