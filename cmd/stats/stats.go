package main

import (
	"fmt"
	"io"
	"os"

	"korsaj.io/rootme/internal/http"
	"korsaj.io/rootme/pkg/printer"

	"github.com/urfave/cli/v2"
)

func run(args []string, stdout, stderr io.Writer) int {
	pr := printer.NewTreePrinter(stdout, stderr)

	app := &cli.App{
		Name:      "root-me stats",
		Usage:     "formation of a list of solved tasks",
		Writer:    stdout,
		ErrWriter: stderr,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "api",
				Usage:    "root-me api url",
				Required: true,
				Aliases:  []string{"a"},
				EnvVars:  []string{"ROOT_ME_API_URL"},
			},
			&cli.StringFlag{
				Name:     "key",
				Usage:    "root-me api key",
				Required: true,
				Aliases:  []string{"k"},
				EnvVars:  []string{"ROOT_ME_API_KEY"},
			},
			&cli.StringFlag{
				Name:     "uid",
				Usage:    "root-me api uid",
				Required: true,
				Aliases:  []string{"u"},
				EnvVars:  []string{"ROOT_ME_API_UID"},
			},
		},
		Action: func(c *cli.Context) error {
			http.InitAPI(c.String("api"),
				c.String("key"),
				c.String("uid"),
			)

			pr.PrintText("process...")

			profile, err := http.UserProfile(c.Context)
			if err != nil {
				return err
			}

			pr.PrintProfile(profile)
			return nil
		},
	}

	if err := app.Run(args); err != nil {
		pr.PrintError(fmt.Sprintf("%v\n", err))
		return 127
	}
	return 0

}

func main() { os.Exit(run(os.Args, os.Stdout, os.Stderr)) }
