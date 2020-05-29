package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var commanderMap map[string]Command

func init() {
	commanderMap = make(map[string]Command)
}

var command Command

func Add(name string, command Command) {
	commanderMap[name] = command
}

func Run(app *cli.App) {
	app.Name = "beego"
	app.Version = "v0.1"
	app.Commands = []*cli.Command{
		{
			Name:  "cli",
			Usage: "cli [command options] [arguments...]",
			Action: func(c *cli.Context) error {
				fmt.Println(c.Command.Name)
				return nil
			},
		},
	}

	subcmds := []*cli.Command{
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "version  show version",
			Action: func(c *cli.Context) error {
				fmt.Printf("%s version=%s\n", c.App.Name, c.App.Version)
				return nil
			},
		},
	}
	for _, commander := range commanderMap {
		subcmd := new(cli.Command)
		subcmd.Action = func(c *cli.Context) error {
			commanderMap[c.Command.Name].Execute()
			return nil
		}
		commander.Configure(subcmd)
		subcmds = append(subcmds, subcmd)
	}
	app.Commands[0].Subcommands = subcmds

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
