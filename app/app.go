package app

import "github.com/urfave/cli"

//Will return the cli app ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Gli"
	app.Usage = "Search IPs and server names"
	return app
}