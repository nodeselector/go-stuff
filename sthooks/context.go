package sthooks

import "context"

type hookContext struct {
	err error
}

func WithError(ctx context.Context, err error) context.Context {
	return context.WithValue(ctx, hookContext{}, hookContext{
		err: err,
	})
}
