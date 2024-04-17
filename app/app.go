package app

import (
	"github.com/urfave/cli"
)

// Provides a Cli Application
func Cli() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nomes de servidor na internet"

	return app
}
