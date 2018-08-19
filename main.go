//usr/bin/env go run $0 $@; exit;
package main

import (
	"fmt"
	"os"

	"github.com/quantumew/aufait/app"
	"github.com/docopt/docopt-go"
)

func main() {
	doc := `Au Fait
Command line interface that assists in setting up Au Fait, the tool that keeps your NPM projects up to date. Set it and forget it (pre)

Usage:
	aufait <repository>...
Options:
	-h --help			Show this message
	--version			Show version info
Arguments:
	<repository>  		Path to repository to setup with Au Fait.`

	arguments, _ := docopt.ParseDoc(doc)
	cli := app.NewCLI(arguments["<repository>"].([]string))
	err := cli.ConfigureRepositories()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
