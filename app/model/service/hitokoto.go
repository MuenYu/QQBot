package service

import (
	"QQBot/app/model"
	"fmt"
	"github.com/mcoo/OPQBot"
)

type Hitokoto model.Function

func (h Hitokoto) Action(packet *OPQBot.SendMsgPack) {
	// TODO 发送一言
}

func (h Hitokoto) Intro() string {
	return fmt.Sprintf("[#%s]: %s\n", h.Name, h.Desc)
}
