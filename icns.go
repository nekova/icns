package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "icns"
	app.Version = Version
	app.Usage = ""
	app.Author = "nekova"
	app.Email = "nekova07@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
