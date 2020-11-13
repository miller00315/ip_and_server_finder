package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicaçõ de linha de comando a ser executada
func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "Command line aplication"

	app.Usage = "Search ip and and servers name"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
		cli.StringFlag{},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search ip internet",
			Flags:  flags,
			Action: searchIp,
		},
		{
			Name:   "servers",
			Usage:  "Search server name",
			Flags:  flags,
			Action: searchName,
		},
	}

	return app
}

func searchIp(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchName(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
