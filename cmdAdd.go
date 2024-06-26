package main

import (
	"ali-ddns/ali"

	"github.com/urfave/cli/v2"
)

var cmdAdd = &cli.Command{
	Name:  "add",
	Usage: "添加一条新的记录",
	Flags: []cli.Flag{
		flagDomain,
		flagRecordType,
		flagValue,
	},
	Action: actionAdd,
}

func actionAdd(ctx *cli.Context) error {
	domain := flagDomain.Get(ctx)
	t := flagRecordType.Get(ctx)
	value := flagValue.Get(ctx)
	return ali.AddRecord(ali.AddRecordParams{
		Domain: domain,
		Type:   t,
		Value:  value,
	})
}
