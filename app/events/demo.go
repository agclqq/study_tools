package events

import (
	"context"
	"fmt"
)

type Demo struct {
}

func (d *Demo) ListenName() string {
	return "test"
}
func (d *Demo) Concurrence() int64 {
	return 1
}
func (d *Demo) Handle(ctx context.Context, data []byte) {
	fmt.Println("demo event", string(data))
}
