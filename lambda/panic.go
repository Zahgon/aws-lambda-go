// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

package lambda

import (
	"runtime"

	"github.com/aws/aws-lambda-go/lambda/messages"
)

type panicInfo struct {
	Message    string                                      // Value passed to panic call, converted to string
	StackTrace []*messages.InvokeResponse_Error_StackFrame // Stack trace of the panic
}

func getPanicInfo(value interface{}) panicInfo { _ = "STUB: not implemented"; return *new(panicInfo) }

func getPanicMessage(value interface{}) string { _ = "STUB: not implemented"; return "" }

var defaultErrorFrameCount = 32

func getPanicStack() []*messages.InvokeResponse_Error_StackFrame {
	_ = "STUB: not implemented"
	return nil
}

// this (getPanicStack) -> getPanicInfo -> handler defer func

func convertStack(s []uintptr) []*messages.InvokeResponse_Error_StackFrame {
	_ = "STUB: not implemented"
	return nil
}

func formatFrame(inputFrame runtime.Frame) *messages.InvokeResponse_Error_StackFrame {
	_ = "STUB: not implemented"
	return nil
}

// Strip GOPATH from path by counting the number of seperators in label & path
//
// For example given this:
//     GOPATH = /home/user
//     path   = /home/user/src/pkg/sub/file.go
//     label  = pkg/sub.Type.Method
//
// We want to set:
//     path  = pkg/sub/file.go
//     label = Type.Method

// Something went wrong and path has less seperators than we expected
// Abort and leave i as -1 to counteract the +1 below

// Trim the initial /

// Strip the path from the function name as it's already in the path

// Likewise strip the package name
