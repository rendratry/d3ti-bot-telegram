package main

import (
	"d3ti-bot-telegram/helper"
	"fmt"
	"regexp"
	"strings"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestString(t *testing.T) {
	str1 := helper.EncodeToString(4)
	str2 := helper.EncodeToString(5)
	sentence := "Jangan merubah pesan ini! kode verifikasi anda jihi-huyio"
	re := regexp.MustCompile(`\d{4}-\d{5}`)
	phone := re.FindString(sentence + str1 + "-" + str2)
	fmt.Println("Phone number:", phone)
}

func TestGeneratePW(t *testing.T) {
	pw := helper.GeneratePassword("rendra tri kusuma")
	fmt.Println(strings.ToLower(pw))
}

func TestCreateUUID(t *testing.T) {
	uuidstr := uuid.NewV4()
	fmt.Println(uuidstr)
}

func TestEkstrakString(t *testing.T) {
	s := "berikut ini adalah string nya:1839-94893.93d0cff1-4768-49a2-ad05-679168d218a9"
	parts := strings.Split(s, ":")
	if len(parts) > 1 {
		result := parts[1]
		fmt.Println(result)
	} else {
		fmt.Println("Tidak ditemukan karakter ':' dalam string")
	}
}
