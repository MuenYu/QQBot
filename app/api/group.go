package api

import (
	"QQBot/app/def"
	"QQBot/app/model/msg"
	"QQBot/app/util"
	"encoding/json"
	"github.com/mcoo/OPQBot"
	"log"
)

// 群消息逻辑
func groupMessage(botQQ int64, packet *OPQBot.GroupMsgPack) {
	if packet.FromUserID != botQQ {
		var atContent = atCheck(botQQ, packet)
		if atContent != nil {
			var (
				instruct = util.InstructExtractor(atContent.Content)
				msgPack  = OPQBot.SendMsgPack{
					SendToType: OPQBot.SendToTypeGroup,
					ToUserUid:  packet.FromUserID,
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
				if v, ok := def.GroupTasks[instruct]; ok {
					v.Action(&msgPack)
				} else {
					printErrMsg(&msgPack)
				}
			} else {
				printMenuMsg(&msgPack)
			}

		}
	}
}

// 检测是否被 at, 并对消息内容进行解析
func atCheck(botQQ int64, packet *OPQBot.GroupMsgPack) *msg.AtContent {
	var (
		atContent = &msg.AtContent{}
		err       error
	)
	if packet.MsgType != "AtMsg" {
		return nil
	}
	if err = json.Unmarshal([]byte(packet.Content), atContent); err != nil {
		log.Println(err)
		return nil
	}
	for _, v := range atContent.UserID {
		if v == botQQ {
			return atContent
		}
	}
	return nil
}
