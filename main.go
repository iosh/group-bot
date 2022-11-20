package main

import (
	"fmt"
	"group-bot/internal/mirai"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"

	"github.com/gorilla/websocket"
)

func main() {

	// mirai server address
	MIRAI_ADDRESS := os.Getenv("MIRAI_ADDRESS")

	MIRAI_SERVER_PATH := "/all"

	MIRAI_VERIFY_KEY := os.Getenv("MIRAI_VERIFY_KEY")

	MIRAI_BOT_ID, err := strconv.Atoi(os.Getenv("MIRAI_BOT_ID"))

	if err != nil {
		log.Fatal("bot id can't convert to number")
	}

	wsURL := url.URL{Scheme: "ws", Host: MIRAI_ADDRESS, Path: MIRAI_SERVER_PATH, RawQuery: fmt.Sprintf("verifyKey=%s&qq=%d", MIRAI_VERIFY_KEY, MIRAI_BOT_ID)}
	wsConnect, _, wsErr := websocket.DefaultDialer.Dial(wsURL.String(), nil)

	if wsErr != nil {
		log.Fatal("dial", wsErr)
	}

	bot := mirai.Bot{
		WSConnect: wsConnect,
	}

	done := make(chan any)

	go bot.Run(done)
	interrupt := make(chan os.Signal, 1)
	go func() {
		signal.Notify(interrupt, os.Interrupt)
	}()
	<-interrupt
	close(done)
}
