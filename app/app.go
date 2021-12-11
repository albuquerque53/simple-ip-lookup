package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will generate the CLI app ready to execute
func Generate() *cli.App {
    app := cli.NewApp()

    app.Name = "Simple IP Loopkup"
    app.Usage = "Look for IPs and nameservers on web"

    app.Commands = []cli.Command{
        {
            Name: "ip",
            Usage: "Look for public IP addresses on the internet",
            Flags: []cli.Flag {
                cli.StringFlag{
                    Name: "host",
                    Value: "duck.com",
                },
            },
            Action: lookupIps,
        },
    }

    return app
}

func lookupIps(context *cli.Context) {
    host := context.String("host")

    fmt.Printf("Looking for %s...\n", host)
    ips, error := net.LookupIP(host)

    if error != nil {
        log.Fatal(error)
    }

    fmt.Print("\nRESULTS::\n")
    for _, ip := range ips {
        fmt.Println(ip)
    }
}
