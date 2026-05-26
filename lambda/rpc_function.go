// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

//go:build !lambda.norpc
// +build !lambda.norpc

package lambda

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda/messages"
)

func init() {
	// Register `startFunctionRPC` to be run if the _LAMBDA_SERVER_PORT environment variable is set.
	// This happens when the runtime for the function is configured as `go1.x`.
	// The value of the environment variable will be passed as the first argument to `startFunctionRPC`.
	// This allows users to save a little bit of coldstart time in the download, by the dependencies brought in for RPC support.
	// The tradeoff is dropping compatibility with the RPC mode of the go1.x runtime.
	// To drop the rpc dependencies, compile with `-tags lambda.norpc`
	startFunctions = append([]*startFunction{{
		env: "_LAMBDA_SERVER_PORT",
		f:   startFunctionRPC,
	}}, startFunctions...)
}

func startFunctionRPC(port string, handler Handler) error { _ = "STUB: not implemented"; return nil }

// Function struct which wrap the Handler
//
// Deprecated: The Function type is public for the go1.x runtime internal use of the net/rpc package
type Function struct {
	handler *handlerOptions
}

// NewFunction which creates a Function with a given Handler
//
// Deprecated: The Function type is public for the go1.x runtime internal use of the net/rpc package
func NewFunction(handler Handler) *Function { _ = "STUB: not implemented"; return nil }

// Ping method which given a PingRequest and a PingResponse parses the PingResponse
func (fn *Function) Ping(req *messages.PingRequest, response *messages.PingResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// Invoke method try to perform a command given an InvokeRequest and an InvokeResponse
func (fn *Function) Invoke(req *messages.InvokeRequest, response *messages.InvokeResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// nolint:staticcheck

func (fn *Function) baseContext() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
