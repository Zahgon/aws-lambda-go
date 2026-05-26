// Package handlertrace allows middleware authors using lambda.NewHandler to
// instrument request and response events.
package handlertrace

import (
	"context"
)

// HandlerTrace allows handlers which wrap the return value of lambda.NewHandler
// to access to the request and response events.
type HandlerTrace struct {
	RequestEvent  func(context.Context, interface{})
	ResponseEvent func(context.Context, interface{})
}

func callbackCompose(f1, f2 func(context.Context, interface{})) func(context.Context, interface{}) {
	_ = "STUB: not implemented"
	return nil
}

type handlerTraceKey struct{}

// NewContext adds callbacks to the provided context which allows handlers which
// wrap the return value of lambda.NewHandler to access to the request and
// response events.
func NewContext(ctx context.Context, trace HandlerTrace) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// FromContext returns the HandlerTrace associated with the provided context.
func FromContext(ctx context.Context) HandlerTrace {
	_ = "STUB: not implemented"
	return *new(HandlerTrace)
}
