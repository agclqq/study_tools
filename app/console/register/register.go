package register

import (
	"github.com/agclqq/prowjob"

	"github.com/agclqq/study_tools/app/console/command"
)

func Register(eng *prowjob.CommandEngine) {
	eng.Add(&command.Demo{})
	eng.Add(&command.PrimaryThreeMath{})
}
