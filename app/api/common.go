package api

import (
	"QQBot/app/def"
	"QQBot/app/model"
	"QQBot/app/util"
	"github.com/mcoo/OPQBot"
	"strings"
)

// 打印菜单
func printMenuMsg(packet *OPQBot.SendMsgPack) {
	var (
		contentBuilder = strings.Builder{}
	)

	contentBuilder.WriteString(def.Conf.System.HeadMsg)
	contentBuilder.WriteString("\n指令列表:\n")
	var taskList map[string]model.Task
	switch packet.SendToType {
	case OPQBot.SendToTypeFriend:
		taskList = def.UserTasks
	case OPQBot.SendToTypeGroup:
		taskList = def.GroupTasks
	default:
		return
	}

	for _, v := range taskList {
		contentBuilder.WriteString(v.Intro())
	}

	if packet.ToUserUid == def.Conf.System.Admin {
		contentBuilder.WriteString("\n管理员指令:\n")
		for _, v := range def.AdminTasks {
			contentBuilder.WriteString(v.Intro())
		}
	}

	util.SendText(contentBuilder.String(), packet)
}

func printErrMsg(packet *OPQBot.SendMsgPack) {
	util.SendText(def.Conf.System.ErrMsg, packet)
}
