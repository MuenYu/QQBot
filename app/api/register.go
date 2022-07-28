package api

import (
	"QQBot/app/def"
	"github.com/mcoo/OPQBot"
	"log"
)

// EventRegister 事件注册
func EventRegister() {
	var (
		opqBot = def.OPQ
		err    error
	)

	if _, err = opqBot.AddEvent(OPQBot.EventNameOnGroupMessage, groupMessage); err != nil {
		log.Println(err)
	}

	if _, err = opqBot.AddEvent(OPQBot.EventNameOnFriendMessage, userMessage); err != nil {
		log.Println(err)
	}
}
