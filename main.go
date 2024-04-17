package main

import (
	"app-cli/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Init")

	app := app.Cli()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
