package main

import (
	"QQBot/app/api"
	"QQBot/app/boot"
	"QQBot/app/def"
	"flag"
	"log"
	"os"
)

func main() {
	var (
		err    error
		opqBot = def.OPQ
	)

	if err = opqBot.Start(); err != nil {
		log.Fatalln(err)
	}
	defer opqBot.Stop()
	api.EventRegister()
	opqBot.Wait()
}

func init() {
	flag.StringVar(&def.Path, "c", "", "the path of boot file, it cannot be null.")
	flag.Parse()
	if def.Path == "" {
		flag.Usage()
		os.Exit(0)
	}
	boot.Boot()
}
