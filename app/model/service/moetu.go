package service

import (
	"QQBot/app/model"
	"fmt"
	"github.com/mcoo/OPQBot"
)

type Moetu model.Function

func (m Moetu) Action(packet *OPQBot.SendMsgPack) {
	// TODO 发送萌图
}

func (m Moetu) Intro() string {
	return fmt.Sprintf("[#%s]: %s\n", m.Name, m.Desc)
}
