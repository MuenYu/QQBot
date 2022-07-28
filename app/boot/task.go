package boot

import (
	"QQBot/app/def"
	"QQBot/app/model"
	"QQBot/app/model/service"
)

func loadTasks() {
	for _, v := range def.Conf.Function {
		var task model.Task
		if task = func2Task(v); task == nil {
			continue
		}
		if v.AdminOnly {
			def.AdminTasks[v.Name] = task
		} else {
			switch v.Trigger {
			case "all":
				def.UserTasks[v.Name] = task
				def.GroupTasks[v.Name] = task
			case "user":
				def.UserTasks[v.Name] = task
			case "group":
				def.GroupTasks[v.Name] = task
			}
		}
	}
}

func func2Task(f model.Function) model.Task {
	switch f.Type {
	case "text":
		return service.Text(f)
	case "moetu":
		return service.Moetu(f)
	case "hitokoto":
		return service.Hitokoto(f)
	}
	return nil
}
