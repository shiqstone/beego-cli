package cmd

import (
	"github.com/astaxie/beego/logs"
	"github.com/urfave/cli/v2"
)

type BaseCmd struct {
	Name string
}

func init() {
	//name := "base"
	// Add(name, &BaseCmd{Name: name})
}

var method string

func (cmder *BaseCmd) Configure(cmd *cli.Command) {
	//setName

	//setDescription

	//addOptions
}

func (cmd *BaseCmd) Execute() {
	logs.Debug("base start Execute")
	//do something
	logs.Notice("base Success")
}
