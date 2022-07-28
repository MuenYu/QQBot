package model

import "github.com/mcoo/OPQBot"

type Task interface {
	// Action 任务执行
	Action(packet *OPQBot.SendMsgPack)
	// Intro 介绍任务
	Intro() string
}
