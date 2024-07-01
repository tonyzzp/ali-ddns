package main

import (
	"ali-ddns/config"
	"ali-ddns/dns"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func test(ctx *cli.Context) error {
	dns.ResolveDomain("pi.veikr.com")
	dns.ResolveDomain("qm.pi.veikr.com")
	dns.ResolveDomain("veikr.com")
	dns.ResolveDomain("www.izzp.me")
	return nil
}

func main() {
	log.SetOutput(os.Stdout)
	app := &cli.App{
		Usage: "操作阿里dns解析记录",
		Flags: []cli.Flag{flagConfig},
		Commands: []*cli.Command{
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
