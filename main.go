package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ritarock/ctj/lib/action"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "ctj",
		Usage: "Convert CSV to JSON",
		Action: func(c *cli.Context) error {
			if c.NArg() > 0 {
				action.ToJson(c.Args().Get(0))
			} else {
				fmt.Println("No file is selected.")
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
