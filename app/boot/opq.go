package boot

import (
	"QQBot/app/def"
	"github.com/mcoo/OPQBot"
	"log"
)

func loadOPQ() {
	if def.Conf == nil || def.Conf.System == nil {
		log.Fatalln("config of system is not completed!")
	}
	def.OPQ = OPQBot.NewBotManager(def.Conf.System.QQ, def.Conf.System.OPQ)
	if def.OPQ == nil {
		log.Fatalln("opq boot failed!")
	}
}
