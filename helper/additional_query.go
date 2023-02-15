package helper

import (
	"context"
	"d3ti-bot-telegram/app"
)

func UpdateDataMahasiswa(Nohp string, Password string, UsernameTelegram string, IdUserTelegram string, UUID string) {
	db := app.GetConnection()
	ctx := context.Background()
	script := "update user set no_hp = ?, password = ?, username_telegram = ?, id_user_telegram = ? where uuid = ? "
	_, err := db.ExecContext(ctx, script, Nohp, Password, UsernameTelegram, IdUserTelegram, UUID)
	PanicIfError(err)

}

func UpdateRegisStatus(idRegistrasi string, email string) {
	db := app.GetConnection()
	ctx := context.Background()
	script := "update bot_registrasi set regis_status = true where id_registrasi = ? and email = ?"
	_, err := db.ExecContext(ctx, script, idRegistrasi, email)
	PanicIfError(err)
}
