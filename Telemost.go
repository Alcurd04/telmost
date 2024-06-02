package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

var goQuotes = []string{
	"Simple, Poetic, Pithy",
	"Concurrency is not parallelism.",
	"Don't panic.",
	"Documentation is for users.",
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Сервер запущен. Ожидание подключений...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при принятии подключения:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		rand.Seed(time.Now().UnixNano())
		randIndex := rand.Intn(len(goQuotes))
		quote := goQuotes[randIndex]

		_, err := conn.Write([]byte(quote + "\n"))
		if err != nil {
			fmt.Println("Ошибка при отправке сообщения клиенту:", err)
			return
		}

		time.Sleep(3 * time.Second)
	}
}
