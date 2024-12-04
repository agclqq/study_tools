package demo

import (
	"context"
)

type DemoRepo struct {
}

func (d *DemoRepo) TableName() string {
	return "demo_table"
}
func (d *DemoRepo) Select(ctx context.Context, where any) *EntityA {
	return &EntityA{Id: 1, Name: "aa", Status: 1}
}
