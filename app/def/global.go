package def

import (
	"QQBot/app/model"
	"github.com/mcoo/OPQBot"
)

var (
	Path string
	Conf = &model.Config{}

	OPQ *OPQBot.BotManager

	UserTasks  = make(map[string]model.Task)
	GroupTasks = make(map[string]model.Task)
	AdminTasks = make(map[string]model.Task)
)
