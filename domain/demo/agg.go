package demo

import (
	"context"
)

var (
	_ Demo = (*DemoAgg)(nil)
)

type DemoAgg struct {
}

func NewAgg() *DemoAgg {
	return &DemoAgg{}
}
func (d *DemoAgg) GetTest(ctx context.Context, id int) (*EntityA, error) {
	repo := &DemoRepo{}
	return repo.Select(ctx, &EntityA{Id: id}), nil
}
