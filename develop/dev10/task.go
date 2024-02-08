package dev

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

func telnetConnect() {
	host := flag.String("host", "", "Хост (IP или доменное имя)")
	port := flag.Int("port", 0, "Порт")
	timeout := flag.Duration("timeout", 10*time.Second, "Таймаут на подключение")
	flag.Parse()

	if *host == "" || *port == 0 {
		fmt.Println("Необходимо указать хост и порт.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	addr := fmt.Sprintf("%s:%d", *host, *port)

	conn, err := net.DialTimeout("tcp", addr, *timeout)
	if err != nil {
		fmt.Printf("Ошибка подключения: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Printf("Ошибка чтения из сокета: %v\n", err)
				os.Exit(1)
			}
			fmt.Print(string(buf[:n]))
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	<-sigCh
	fmt.Println("\nЗавершение программы.")
}
