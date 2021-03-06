package main

import (
	"log"
	"os"

	"github.com/sadayuki-matsuno/git-credential-substitute/internal"
	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
)

func main() {
	app := &cli.App{
		Name:                 "git-credential-substitute",
		Usage:                "git credential helper, which switches git credentials for the username/organization.",
		Action:               internal.Action,
		Version:              version,
		EnableBashCompletion: true,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
