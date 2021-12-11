package app

import "github.com/urfave/cli"

// Generate will generate the CLI app ready to execute
func Generate() *cli.App {
    app := cli.NewApp()

    app.Name = "Simple IP Loopkup"
    app.Usage = "Look for IPs and nameservers on web"

    return app
}
