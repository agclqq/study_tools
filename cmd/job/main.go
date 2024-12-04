package main

import (
	prowjobreg "github.com/agclqq/prow-framework/prowjob/register"
	"github.com/agclqq/prowjob"

	"github.com/agclqq/study_tools/app/console/register"
)

func main() {
	jobEng := prowjob.New()
	prowjobreg.Register(jobEng)
	register.Register(jobEng)
	jobEng.Run()
}
