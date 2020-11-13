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

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search ip internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com.br",
				},
			},
			Action: searchIp,
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
