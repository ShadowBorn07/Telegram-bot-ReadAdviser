package main

import (
	tgClient "ReadAdviser_bot/clients/telegram"
	event_consumer "ReadAdviser_bot/consumer/event-consumer"
	"ReadAdviser_bot/events/telegram"
	"ReadAdviser_bot/storage/files"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

//bot -tg-bot-token ''

func main() {
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
		"5357332269:AAEoWdOPo3JbaUUQohF3nOQy3P2jcFt3KF4",
		"token for access to telegram token bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
