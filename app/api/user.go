package api

import (
	"QQBot/app/def"
	"QQBot/app/util"
	"github.com/mcoo/OPQBot"
	"log"
)

// 好友消息逻辑
func userMessage(botQQ int64, packet *OPQBot.FriendMsgPack) {
	if botQQ != packet.FromUin {
		def.OPQ.Session.SessionStart(packet.FromUin)
		var (
			instruct = util.InstructExtractor(packet.Content)
			msgPack  = OPQBot.SendMsgPack{
				SendToType: OPQBot.SendToTypeFriend,
				ToUserUid:  packet.FromUin,
				Content:    nil,
				CallbackFunc: func(Code int, Info string, record OPQBot.MyRecord) {
					log.Println(record)
					if Code == 0 {
						log.Println("发送成功")
					} else {
						log.Println("发送失败 错误代码", Code, Info)
					}
				},
			}
		)

		if len(instruct) > 0 {
			if v, ok := def.UserTasks[instruct]; ok {
				v.Action(&msgPack)
			} else {
				printErrMsg(&msgPack)
			}
		} else {
			printMenuMsg(&msgPack)
		}
	}
}
