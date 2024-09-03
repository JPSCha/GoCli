package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Will return the cli app ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Gli"
	app.Usage = "Search IPs and server names on the web"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Usage: "host to search",
			Value: "",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Search for IP adresses on the web",
			Flags: flags,

			Action: SearchIp,

		},
		{
			Name: "servers",
			Usage: "Search for server names on the web",
			Flags: flags,

			Action: SearchServers,
		},
	}

	return app
}

func SearchIp(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func SearchServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host) //NS = name server
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}