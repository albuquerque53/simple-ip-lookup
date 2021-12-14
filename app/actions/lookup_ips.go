package actions

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// LookupIPs will search public IPs based on given host
func LookupIPs(context *cli.Context) {
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
