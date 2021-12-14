package actions

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// LookupHosts will search for hostnames based on specific host
func LookupHosts(context *cli.Context) {
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
