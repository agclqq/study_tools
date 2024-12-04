package demo

import (
	"context"
)

type Demo interface {
	GetTest(ctx context.Context, id int) (*EntityA, error)
}
