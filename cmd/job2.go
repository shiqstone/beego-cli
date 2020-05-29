package cmd

import (
	"beego-cli/models"

	"github.com/astaxie/beego/logs"
	"github.com/urfave/cli/v2"
)

type Job2Cmt struct {
	Name string
}

func init() {
	name := "job_exam_2"
	Add(name, &Job2Cmt{Name: name})
}

func (cmder *Job2Cmt) Configure(cmd *cli.Command) {
	//setName
	cmd.Name = cmder.Name
	//setDescription
	cmd.Usage = `cli job_exam_2 [options] [arguments...]`

	//addOptions
	cmd.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "method",
			Aliases: []string{"m"},
			Usage: `method to be execute 
	all - get all users cache`,
			Value:       "all",
			Destination: &method,
		},
	}
}

func (cmd *Job2Cmt) Execute() {
	logs.Debug("job_exam_2 start Execute")
	if method == "all" {
		objs := models.GetAll()
		logs.Notice(objs)
	} else {
		logs.Error("key not correct: ", method)
		return
	}
	logs.Notice("job_exam_2 Success")
}
