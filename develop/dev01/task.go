package dev

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

func getTime() {
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		log.Fatal("Ошибка при получении времени с NTP сервера:", err)
	}
	fmt.Println("Точное время с использованием NTP:", ntpTime)
}
