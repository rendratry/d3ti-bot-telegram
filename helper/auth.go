package helper

import (
	"context"
	"d3ti-bot-telegram/app"
	"d3ti-bot-telegram/model/domain"
	"errors"
)

func AuthUser(username string) bool {
	db := app.GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "select username_telegram, no_hp from user where username_telegram = ?"
	rows, err := db.QueryContext(ctx, script, username)
	PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.UsernameTelegram, &user.NoHp)
		PanicIfError(err)
		return true
	} else {
		return false
	}
}

func VerifikasiChatbot(idverifikasi string) (domain.VerifikasiChatbot, error) {
	db := app.GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "select `id_registrasi`, `email`, `nama_lengkap`, `no_hp` from bot_registrasi where id_registrasi = ? and regis_status = 0"
	rows, err := db.QueryContext(ctx, script, idverifikasi)
	PanicIfError(err)
	defer rows.Close()

	verifikasi := domain.VerifikasiChatbot{}
	if rows.Next() {
		err := rows.Scan(&verifikasi.IdRegister, &verifikasi.Email, &verifikasi.NamaLengkap, &verifikasi.NoHp)
		PanicIfError(err)
		return verifikasi, nil
	} else {
		return verifikasi, errors.New("verifikasi gagal")
	}
}
