package app

import (
	"simple-ip-lookup/app/actions"

	"github.com/urfave/cli"
)

// Generate will generate the CLI app ready to execute
func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "Simple IP Loopkup"
	app.Usage = "Look for IPs and nameservers on web"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "duck.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Look for public IP addresses on the internet",
			Flags:  flags,
			Action: actions.LookupIPs,
		},
		{
			Name:   "server",
			Usage:  "Look for hostnames on the internet",
			Flags:  flags,
			Action: actions.LookupHosts,
		},
	}

	return app
}
