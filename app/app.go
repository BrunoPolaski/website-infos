package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line app"
	app.Usage = "Search for website IPs and names"
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search for IPs",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: searchIps,
		},
		{
			Name:  "server",
			Usage: "Search for server name",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: searchServerNames,
		},
	}
	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServerNames(c *cli.Context) {
	host := c.String("host")

	names, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range names {
		fmt.Println(name.Host)
	}
}
