package main

import (
	"flag"
	"log"
	tgClient "tgBot/clients/telegram"
	"tgBot/consumer/event-consumer"
	"tgBot/events/telegram"
	"tgBot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	////token = flags.Get(token) запуск через флаг чтобы не палить токен в гите
	//tgClient = telegram.New(token) токен для связки бота и языка
	//fetcher = fetcher.New(tgClient) для получения данных нужен фетчер
	//processor = processor.New(tgClient) для обработки данных нужен процессор
	//consumer.Start(fetcher, processor) получает события и обрабатывает их

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

}

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
