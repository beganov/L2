// L2.8
// Создать программу, печатающую точное текущее время с использованием NTP-сервера.
// Реализовать проект как модуль Go.
// Использовать библиотеку ntp для получения времени.
// Программа должна выводить текущее время, полученное через NTP (Network Time Protocol).
// Необходимо обрабатывать ошибки библиотеки: в случае ошибки вывести её текст в STDERR и вернуть ненулевой код выхода.
// Код должен проходить проверки (vet и golint), т.е. быть написан идиоматически корректно.

// Package ntpclient предоставляет утилиту для получения точного времени через NTP-сервер.
package ntpclient

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

// GetExactTime возвращает точное время с NTP-сервера.
func GetExactTime() {
	response, err := ntp.Query("0.ru.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(time.Now().Add(response.ClockOffset).Format("02.01.2006 15:04:05.000 MST"))
}
