package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Retornar a aplicação de linha de comando para ser executada
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application created in Go!"
	app.Description = "Search ips and server names on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	//cli.Command é um array de structs dos comandos da aplicação CLI
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search ip from addresses on the internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "server",
			Usage:  "Search server names from addresses on the internet",
			Flags:  flags,
			Action: searchServerNames,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("Found %d IP's for %s host: \n", len(ips), host)

	for _, ip := range ips {
		fmt.Println("* ", ip)
	}
}

func searchServerNames(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("Found %d Server Names for %s host: \n", len(servers), host)

	for _, server := range servers {
		fmt.Println("* ", server.Host)
	}
}
