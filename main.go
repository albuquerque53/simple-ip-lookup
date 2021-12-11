package main

import (
	"fmt"
	"log"
	"os"
	"simple-ip-lookup/app"
)

func main() {
    fmt.Print("\n=-=-=-= SIMPLE IP Lookup =-=-=-=\n\n\n")

    simpleIpLookup := app.Generate()

    if error := simpleIpLookup.Run(os.Args); error != nil {
        log.Fatal(error)
    }
}

