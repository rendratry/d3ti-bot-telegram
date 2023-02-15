package main

import (
	"d3ti-bot-telegram/helper"
	"d3ti-bot-telegram/model/domain"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	mytoken := "5836520934:AAGQ_iQvY7Hbm5goVRPbwH-k57p25dA-Gns"
	bot, err := tgbotapi.NewBotAPI(mytoken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			log.Println(update.Message.From.ID)

			auth := helper.AuthUser(update.Message.From.UserName)

			if auth {
				mess := helper.Messages(strings.ToLower(update.Message.Text))
				log.Println(mess)
				if mess != "" {
					if strings.Contains(mess, "%s") {
						welcomeMsg := fmt.Sprintf(mess, update.Message.From.UserName)
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, welcomeMsg)
						bot.Send(msg)
					} else {
						msg1 := tgbotapi.NewMessage(update.Message.Chat.ID, mess)
						msg1.ReplyToMessageID = update.Message.MessageID
						bot.Send(msg1)
					}
				} else {
					sayError := helper.Greeting("error")
					welcomeMsg := fmt.Sprintf(sayError, update.Message.From.UserName)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, welcomeMsg)
					bot.Send(msg)
				}

			} else {
				if strings.Contains(update.Message.Text, "Jangan merubah pesan ini! silahkan dikirim ke chatbot, validasi:") {
					verifikasibot := domain.VerifikasiChatbot{}

					resultSplit := helper.SplitResult{}
					resultSplit = helper.Spliter(update.Message.Text)

					verifikasibot, err = helper.VerifikasiChatbot(resultSplit.IdRegis)
					if err != nil {
						greeting := helper.Greeting("greeting")
						if strings.Contains(greeting, "%s") {
							welcomeMsg := fmt.Sprintf(greeting, update.Message.From.UserName)
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, welcomeMsg)
							bot.Send(msg)
						} else {
							msg1 := tgbotapi.NewMessage(update.Message.Chat.ID, greeting)
							msg1.ReplyToMessageID = update.Message.MessageID
							bot.Send(msg1)
						}
					}

					passwordgen := helper.GeneratePassword(verifikasibot.NamaLengkap)
					passwordhash, err := helper.HashPassword(passwordgen)
					helper.PanicIfError(err)

					sendmsg := "Yayyy... Akun berhasil diverifikasi..🎉🎉\nDengan akses akun :\n"
					helper.UpdateDataMahasiswa(verifikasibot.NoHp, passwordhash, update.Message.From.UserName, strconv.Itoa(int(update.Message.From.ID)), resultSplit.Uuid)
					helper.PanicIfError(err)

					msg1 := tgbotapi.NewMessage(update.Message.Chat.ID, sendmsg+"\nNama : "+verifikasibot.NamaLengkap+"\nEmail : "+verifikasibot.Email+"\nNomor Hp : "+verifikasibot.NoHp+"\nPassword : "+passwordgen+"\n\nSilahkan masuk dengan akun diatas kemudian ganti password anda agar lebih aman dan mudah diingat")
					msg1.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg1)
					msg2 := tgbotapi.NewMessage(update.Message.Chat.ID, "Sekarang kamu bisa berinteraksi dengan saya☺️")
					bot.Send(msg2)
					helper.UpdateRegisStatus(resultSplit.IdRegis, verifikasibot.Email)

				} else {
					greeting := helper.Greeting("greeting")
					if strings.Contains(greeting, "%s") {
						welcomeMsg := fmt.Sprintf(greeting, update.Message.From.UserName)
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, welcomeMsg)
						bot.Send(msg)
					} else {
						msg1 := tgbotapi.NewMessage(update.Message.Chat.ID, greeting)
						msg1.ReplyToMessageID = update.Message.MessageID
						bot.Send(msg1)
					}
				}

			}

		}
	}
}
