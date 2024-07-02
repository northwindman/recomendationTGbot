package main

import (
	"flag"
	"log"
	tgClient "northwindman_tgBot/clients/telegram"
	event_consumer "northwindman_tgBot/consumer/event-consumer"
	"northwindman_tgBot/events/telegram"
	"northwindman_tgBot/lib/storage/files"
)

// нужно сделать на подобии функции mustToken для большей гибкости и функциональности кода и приложения
const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

// 7230348904:AAHI2MTGCEjJ7TlkQQ7RWULLbjulR86YaR0

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("Service is stopped", err)
	}
}

// функция получения токена из флагов запуска строки
func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
