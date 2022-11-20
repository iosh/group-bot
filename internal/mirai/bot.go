package mirai

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Bot struct {
	WSConnect *websocket.Conn
}

func (bot *Bot) Run(done <-chan any) {
	for {
		select {
		case <-done:
			fmt.Println("loop is done!")
			return
		default:
			messageType, p, err := bot.WSConnect.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(messageType, string(p))
		}
	}
}
