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

    flags := []cli.Flag {
        cli.StringFlag{
            Name: "host",
            Value: "duck.com",
        },
    }

    app.Commands = []cli.Command{
        {
            Name: "ip",
            Usage: "Look for public IP addresses on the internet",
            Flags: flags,
            Action: lookupIps,
        },
        {
            Name: "server",
            Usage: "Look for hostnames on the internet",
            Flags: flags,
            Action: lookupHosts,
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

func lookupHosts(context *cli.Context) {
    host := context.String("host")

    fmt.Printf("Looking for hosts in %s...\n", host)
    hosts, error := net.LookupNS(host)

    if error != nil {
        log.Fatal(error)
    }

    fmt.Print("\nRESULTS::\n")
    for _, host := range hosts {
        fmt.Println(host.Host)
    }
}
