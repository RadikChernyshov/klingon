package main

import (
	"github.com/RadikChernyshov/klingon/pkg/colorizer"
	"github.com/RadikChernyshov/klingon/pkg/recognizer"
	"github.com/RadikChernyshov/klingon/pkg/translator"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "Translate a name written in English to Klingon"
	app.Usage = "cli application"
	app.UsageText = "klingon name"
	app.Version = "1.0.0"
	app.Action = func(ctx *cli.Context) error {
		name := strings.TrimSpace(strings.Join(ctx.Args(), " "))
		t := translator.New()
		r := recognizer.New()

		t.In = name
		r.In = name

		err := t.ToKlingon()
		if err != nil {
			return cli.NewExitError(colorizer.Red(err.Error()), 1)
		}
		colorizer.PrintYellow(t.Out)
		err = r.Recognize()
		if err != nil {
			return cli.NewExitError(colorizer.Red(err.Error()), 1)
		}
		colorizer.PrintBlue(r.Out)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
