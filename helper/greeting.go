package helper

import (
	"context"
	"d3ti-bot-telegram/app"
)

func Greeting(key string) string {

	db := app.GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "select kunci, messages from bot_telegram_greeting where kunci = ?"
	rows, err := db.QueryContext(ctx, script, key)
	PanicIfError(err)
	defer rows.Close()

	var kunci string
	var messages string
	if rows.Next() {
		rows.Scan(&kunci, &messages)
	}

	return messages
}
