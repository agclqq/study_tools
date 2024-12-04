package command

import (
	"fmt"

	"github.com/agclqq/prowjob"
)

type Demo struct {
}

func (d *Demo) GetCommand() string {
	return "command:demo"
}
func (d *Demo) Usage() string {
	return `Usage of command:demo:
	  command:demo
	`
}
func (d *Demo) Handle(ctx *prowjob.Context) {
	fmt.Println("this is command demo")
}
