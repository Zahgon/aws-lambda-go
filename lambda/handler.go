// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

package lambda

import (
	"bytes"
	"context"
	"io" // nolint:staticcheck
	"reflect"
	"sync"
)

type Handler interface {
	Invoke(ctx context.Context, payload []byte) ([]byte, error)
}

type handlerOptions struct {
	handlerFunc
	baseContext                      context.Context
	contextValues                    map[interface{}]interface{}
	jsonRequestUseNumber             bool
	jsonRequestDisallowUnknownFields bool
	jsonResponseEscapeHTML           bool
	jsonResponseIndentPrefix         string
	jsonResponseIndentValue          string
	enableSIGTERM                    bool
	sigtermCallbacks                 []func()
	jsonOutBufferPool                *sync.Pool // contains *jsonOutBuffer
}

type Option func(*handlerOptions)

// WithContext is a HandlerOption that sets the base context for all invocations of the handler.
func WithContext(ctx context.Context) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContextValue adds a value to the handler context.
// If a base context was set using WithContext, that base is used as the parent.
func WithContextValue(key interface{}, value interface{}) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSetEscapeHTML sets the SetEscapeHTML argument on the underlying json encoder
func WithSetEscapeHTML(escapeHTML bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSetIndent sets the SetIndent argument on the underling json encoder
func WithSetIndent(prefix, indent string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUseNumber sets the UseNumber option on the underlying json decoder
func WithUseNumber(useNumber bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDisallowUnknownFields sets the DisallowUnknownFields option on the underlying json decoder
func WithDisallowUnknownFields(disallowUnknownFields bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEnableSIGTERM enables SIGTERM behavior within the Lambda platform on container spindown.
// SIGKILL will occur ~500ms after SIGTERM.
// Optionally, an array of callback functions to run on SIGTERM may be provided.
func WithEnableSIGTERM(callbacks ...func()) Option { _ = "STUB: not implemented"; return *new(Option) }

// handlerTakesContext returns whether the handler takes a context.Context as its first argument.
func handlerTakesContext(handler reflect.Type) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// handlers like func(event any) are valid.

func validateReturns(handler reflect.Type) error { _ = "STUB: not implemented"; return nil }

// NewHandler creates a base lambda handler from the given handler function. The
// returned Handler performs JSON serialization and deserialization, and
// delegates to the input handler function. The handler function parameter must
// satisfy the rules documented by Start. If handlerFunc is not a valid
// handler, the returned Handler simply reports the validation error.
func NewHandler(handlerFunc interface{}) Handler { _ = "STUB: not implemented"; return *new(Handler) }

// NewHandlerWithOptions creates a base lambda handler from the given handler function. The
// returned Handler performs JSON serialization and deserialization, and
// delegates to the input handler function. The handler function parameter must
// satisfy the rules documented by Start. If handlerFunc is not a valid
// handler, the returned Handler simply reports the validation error.
func NewHandlerWithOptions(handlerFunc interface{}, options ...Option) Handler {
	_ = "STUB: not implemented"
	return *new(Handler)
}

func newHandler(handlerFunc interface{}, options ...Option) *handlerOptions {
	_ = "STUB: not implemented"
	return nil
}

type handlerFunc func(context.Context, []byte) (io.Reader, error)

// back-compat for the rpc mode
func (h handlerFunc) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if the response needs to be closed (ex: net.Conn, os.File), ensure it's closed before the next invoke to prevent a resource leak

// optimization: if the response is a *bytes.Buffer, a copy can be eliminated

func errorHandler(err error) handlerFunc { _ = "STUB: not implemented"; return *new(handlerFunc) }

type jsonOutBuffer struct {
	pool *sync.Pool
	*bytes.Buffer
}

func (j *jsonOutBuffer) ContentType() string { _ = "STUB: not implemented"; return "" }

func (j *jsonOutBuffer) Close() error { _ = "STUB: not implemented"; return nil }

func reflectHandler(f interface{}, h *handlerOptions) handlerFunc {
	_ = "STUB: not implemented"
	return *new(handlerFunc)
}

// back-compat: types with reciever `Invoke(context.Context, []byte) ([]byte, error)` need the return bytes wrapped

// If the final return value is not our buffer, reset and return it to the pool.
// The caller of the handlerFunc does this otherwise.

// construct arguments

// return the error, if any

// set the response value, if any

// encode to JSON

// if response is not JSON serializable, but the response type is a reader, return it as-is

// if response value is an io.Reader, return it as-is

// back-compat, don't return the reader if the value serialized to a non-empty json

// back-compat, strip the encoder's trailing newline unless WithSetIndent was used
