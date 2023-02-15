package helper

import (
	"context"
	"d3ti-bot-telegram/app"
	"d3ti-bot-telegram/model/domain"
	"log"
	"strings"
)

func Messages(kunci string) string {

	db := app.GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "select tag1, tag2, tag3, tag4, tag5, messages from bot_telegram_messages"
	rows, err := db.QueryContext(ctx, script)
	PanicIfError(err)
	defer rows.Close()

	log.Println(rows)
	var pesan string
	var pesannil string
	msg := domain.MessagesModel{}
	for rows.Next() {
		err := rows.Scan(&msg.Tag1, &msg.Tag2, &msg.Tag3, &msg.Tag4, &msg.Tag5, &pesan)
		PanicIfError(err)
		if strings.Contains(kunci, msg.Tag1) || strings.Contains(kunci, msg.Tag2) || strings.Contains(kunci, msg.Tag3) || strings.Contains(kunci, msg.Tag4) || strings.Contains(kunci, msg.Tag5) {
			return pesan
		}
	}
	return pesannil
}
