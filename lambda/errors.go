// Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved

package lambda

import (
	"github.com/aws/aws-lambda-go/lambda/messages"
)

func getErrorType(err interface{}) string { _ = "STUB: not implemented"; return "" }

func lambdaErrorResponse(invokeError error) *messages.InvokeResponse_Error {
	_ = "STUB: not implemented"
	return nil
}

func lambdaPanicResponse(err interface{}) *messages.InvokeResponse_Error {
	_ = "STUB: not implemented"
	return nil
}
