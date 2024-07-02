package main

import (
	"xddns/dns"

	"github.com/urfave/cli/v2"
)

func cmdDelAction(ctx *cli.Context) error {
	t := flagRecordType.Get(ctx)
	domain := flagDomain.Get(ctx)
	return obtainClient(domain).DelRecord(dns.DelRecordParams{
		Domain: domain,
		Type:   t,
	})
}

var cmdDel = &cli.Command{
	Name:  "del",
	Usage: "del dns record",
	Flags: []cli.Flag{
		flagDomain,
		flagRecordType,
	},
	Action: cmdDelAction,
}
