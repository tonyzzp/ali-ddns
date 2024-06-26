package main

import (
	"ali-ddns/ali"
	"ali-ddns/config"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func test(ctx *cli.Context) error {
	ali.ResolveDomain("pi.veikr.com")
	ali.ResolveDomain("qm.pi.veikr.com")
	ali.ResolveDomain("veikr.com")
	ali.ResolveDomain("www.izzp.me")
	return nil
}

func main() {
	log.SetOutput(os.Stdout)
	app := &cli.App{
		Usage: "操作阿里dns解析记录",
		Flags: []cli.Flag{flagConfig},
		Commands: []*cli.Command{
			cmdAdd,
			cmdSet,
			cmdUpdate,
			cmdDel,
			cmdDump,
			{
				Name:   "test",
				Action: test,
			},
		},
		Before: func(ctx *cli.Context) error {
			configFile := ctx.String("config")
			log.Println("configFile", configFile)
			return config.Init(configFile)
		},
	}
	e := app.Run(os.Args)
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
