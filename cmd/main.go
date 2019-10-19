package main

import (
	"github.com/RadikChernyshov/klingon/pkg/recognizer"
	"github.com/RadikChernyshov/klingon/pkg/translator"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "Klingon name transliteration"
	app.Usage = "cli application"
	app.UsageText = "go_build_main_go name"
	app.Version = "1.0.0"
	app.Action = func(ctx *cli.Context) error {
		name := strings.TrimSpace(ctx.Args().Get(0))
		t := translator.New()
		r := recognizer.New()

		t.In = name
		r.In = name

		err := t.ToKlingon()
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		err = r.Recognize()
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		color.Yellow("%s", t.Out)
		color.Blue("%s", r.Out)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
