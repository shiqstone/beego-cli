package main

import (
	"beego-cli/cmd"
	_ "beego-cli/routers"
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/urfave/cli/v2"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	mode := "web"
	app := cli.NewApp()
	app.Action = func(ctx *cli.Context) error {
		if ctx.NArg() > 0 {
			if ctx.Args().Get(0) == "cli" {
				mode = "cli"
				cmd.Run(app)
			}
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}

	if mode != "cli" {
		beego.Run()
	}
}
