package service

import (
	"QQBot/app/model"
	"QQBot/app/util"
	"fmt"
	"github.com/mcoo/OPQBot"
)

type Text model.Function

func (t Text) Action(packet *OPQBot.SendMsgPack) {
	value := t.Config["value"].(string)
	util.SendText(value, packet)
}

func (t Text) Intro() string {
	return fmt.Sprintf("[#%s]: %s\n", t.Name, t.Desc)
}
