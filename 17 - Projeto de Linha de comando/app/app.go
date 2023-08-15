package app

import "github.com/urfave/cli"

//Retornar a aplicação de linha de comando para ser executada
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application created in Go!"
	app.Description = "Search ips and server names on the internet"

	return app
}
