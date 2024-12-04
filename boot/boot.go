package boot

import (
	"fmt"
	"strings"
	"sync"

	"github.com/agclqq/prow-framework/db"
	"github.com/agclqq/prow-framework/event"
	"github.com/spf13/cast"
	"gorm.io/gorm"

	"github.com/agclqq/study_tools/app/events/register"
	"github.com/agclqq/study_tools/config"
)

var (
	onceDbW   = &sync.Once{}
	onceEvent = &sync.Once{}
	dbW       *gorm.DB
)

func GetDbW() *gorm.DB {
	onceDbW.Do(func() {
		// 初始化mysql读写连接
		dbW = db.GetConn("demo", config.GetDb("demo"))
	})
	return dbW
}
func StartEvent() {
	onceEvent.Do(func() {
		vs := []string{}
		for _, v := range config.GetAllEvent() {
			err := event.InitEvent(v["name"], cast.ToInt(v["capacity"]))
			if err != nil {
				fmt.Println("init event error ", err)
				continue
			}
			vs = append(vs, v["name"])
		}
		register.Register()
		event.Run()
		fmt.Printf("%d events are running. \n detail:%s", len(vs), strings.Join(vs, ","))
	})
}
