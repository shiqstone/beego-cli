package cmd

import (
	"beego-cli/models"

	"github.com/astaxie/beego/logs"
	"github.com/urfave/cli/v2"
)

type Job1Cmd struct {
	Name string
}

func init() {
	name := "job_exam_1"
	Add(name, &Job1Cmd{Name: name})
}

var uid string
var gid int

func (cmder *Job1Cmd) Configure(cmd *cli.Command) {
	//setName
	cmd.Name = cmder.Name
	//setDescription
	cmd.Usage = `cli job_exam_1 [options] [arguments...]`

	//addOptions
	cmd.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "method",
			Aliases: []string{"m"},
			Usage: `method to be execute 
	all - get all users cache
	print - print infomatin`,
			Value:       "all",
			Destination: &method,
			//DefaultText: "random",
		},
		&cli.StringFlag{
			Name:        "uid",
			Aliases:     []string{"u"},
			Usage:       "id of user",
			Value:       "",
			Destination: &uid,
		},
		&cli.IntFlag{
			Name:        "gid",
			Aliases:     []string{"g"},
			Usage:       "id of group",
			Value:       0,
			Destination: &gid,
		},
	}
}

func (cmd *Job1Cmd) Execute() {
	logs.Debug("job_exam_1 start Execute")
	if method == "all" {
		users := models.GetAllUsers()
		logs.Notice(users)
	} else if method == "print" {
		if len(uid) == 0 {
			logs.Error("uid can not be empty")
			return
		}
		user, err := models.GetUser(uid)
		if err != nil {
			logs.Error("print failed: ", err)
			return
		}
		logs.Notice(user)
	} else {
		logs.Error("method not exsited: ", method)
		return
	}
	logs.Notice("job_exam_1 Success")
}
