package providers

import (
	"context"
)

type BooleanProvider func(ctx context.Context) (bool, error)
type FloatProvider func(ctx context.Context) (float64, error)
type IntProvider func(ctx context.Context) (int, error)
type StringProvider func(ctx context.Context) (string, error)
