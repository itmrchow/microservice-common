package context

import (
	"context"
	"errors"
)

var ErrEmpty = errors.New("empty")

func GetTraceID(ctx context.Context) (string, error) {
	traceID, ok := ctx.Value(TraceIDKey{}).(string)
	if !ok {
		return "", ErrEmpty
	}
	return traceID, nil
}
