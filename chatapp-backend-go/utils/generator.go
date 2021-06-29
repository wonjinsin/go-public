package utils

import (
	"context"
	"time"
)

func CtxGenerator() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	return ctx, cancel
}
