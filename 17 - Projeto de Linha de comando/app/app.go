package app

import (
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
