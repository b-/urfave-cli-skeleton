package main

import (
	"io"
	"log"
	"os"

	"github.com/luthermonson/urfave-cli-skeleton/cmd"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "new-command",
		Usage:    "",
		Commands: cmd.Commands(),
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Turn on verbose debug logging",
			},
			&cli.BoolFlag{
				Name:    "quiet",
				Aliases: []string{"q"},
				Usage:   "Turn on off all logging",
			},
		},
		Before: func(ctx *cli.Context) error {
			if ctx.Bool("debug") {
				logrus.SetLevel(logrus.DebugLevel)
			} else {
				// treat logrus like fmt.Print
				logrus.SetFormatter(&easy.Formatter{
					LogFormat: "%msg%",
				})
			}
			if ctx.Bool("quiet") {
				logrus.SetOutput(io.Discard)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
