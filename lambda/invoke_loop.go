// Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved

package lambda

import (
	"context"
	"io"
	"time"

	"github.com/aws/aws-lambda-go/lambda/messages"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

const (
	msPerS  = int64(time.Second / time.Millisecond)
	nsPerMS = int64(time.Millisecond / time.Nanosecond)
)

// TODO: replace with time.UnixMillis after dropping version <1.17 from CI workflows
func unixMS(ms int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func doRuntimeAPILoop(ctx context.Context, client *runtimeAPIClient, handler *handlerOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// handleInvoke returns an error if the function panics, or some other non-recoverable error occurred
func handleInvoke(invoke *invoke, handler *handlerOptions) error {
	_ = "STUB: not implemented"
	// set the deadline
	return nil
}

// set the invoke metadata values

// set the trace id

// nolint:staticcheck

// call the handler, marshal any returned error

// if the response needs to be closed (ex: net.Conn, os.File), ensure it's closed before the next invoke to prevent a resource leak

// if the response defines a content-type, plumb it through

func reportFailure(invoke *invoke, invokeErr *messages.InvokeResponse_Error) error {
	_ = "STUB: not implemented"
	return nil
}

func callBytesHandlerFunc(ctx context.Context, payload []byte, handler handlerFunc) (response io.Reader, invokeErr *messages.InvokeResponse_Error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func parseDeadline(invoke *invoke) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func parseCognitoIdentity(invoke *invoke, out *lambdacontext.CognitoIdentity) error {
	_ = "STUB: not implemented"
	return nil
}

func parseClientContext(invoke *invoke, out *lambdacontext.ClientContext) error {
	_ = "STUB: not implemented"
	return nil
}

func safeMarshal(v interface{}) []byte { _ = "STUB: not implemented"; return nil }

// never reach

type xrayException struct {
	Type    string                                      `json:"type"`
	Message string                                      `json:"message"`
	Stack   []*messages.InvokeResponse_Error_StackFrame `json:"stack"`
}

type xrayError struct {
	WorkingDirectory string          `json:"working_directory"`
	Exceptions       []xrayException `json:"exceptions"`
	Paths            []string        `json:"paths"`
}

func makeXRayError(invokeResponseError *messages.InvokeResponse_Error) *xrayError {
	_ = "STUB: not implemented"
	return nil
}
