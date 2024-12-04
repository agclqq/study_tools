package register

import (
	"github.com/agclqq/prow-framework/event"

	"github.com/agclqq/study_tools/app/events"
)

type Demo struct {
}

func Register() {
	event.Register(&events.Demo{})
}
