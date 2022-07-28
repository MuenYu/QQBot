package util

import (
	"QQBot/app/def"
	"github.com/mcoo/OPQBot"
)

// 消息发送工具类

// SendText 发送文本消息
func SendText(content string, packet *OPQBot.SendMsgPack) {
	packet.Content = OPQBot.SendTypeTextMsgContent{
		Content: content,
	}
	def.OPQ.Send(*packet)
}
