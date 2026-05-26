// Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved
//
// Runtime API documentation: https://docs.aws.amazon.com/lambda/latest/dg/runtimes-api.html

package lambda

import (
	"bytes"
	"context"
	"io" //nolint: staticcheck
	"net/http"
	"sync"
)

const (
	headerAWSRequestID       = "Lambda-Runtime-Aws-Request-Id"
	headerDeadlineMS         = "Lambda-Runtime-Deadline-Ms"
	headerTraceID            = "Lambda-Runtime-Trace-Id"
	headerCognitoIdentity    = "Lambda-Runtime-Cognito-Identity"
	headerClientContext      = "Lambda-Runtime-Client-Context"
	headerInvokedFunctionARN = "Lambda-Runtime-Invoked-Function-Arn"
	headerTenantID           = "Lambda-Runtime-Aws-Tenant-Id"
	headerXRayErrorCause     = "Lambda-Runtime-Function-Xray-Error-Cause"
	trailerLambdaErrorType   = "Lambda-Runtime-Function-Error-Type"
	trailerLambdaErrorBody   = "Lambda-Runtime-Function-Error-Body"
	contentTypeJSON          = "application/json"
	contentTypeBytes         = "application/octet-stream"
	apiVersion               = "2018-06-01"
	xrayErrorCauseMaxSize    = 1024 * 1024
)

type runtimeAPIClient struct {
	baseURL    string
	userAgent  string
	httpClient *http.Client
	pool       *sync.Pool
}

func newRuntimeAPIClient(address string) *runtimeAPIClient { _ = "STUB: not implemented"; return nil }

// connections to the runtime API are never expected to time out

type invoke struct {
	id      string
	payload *bytes.Buffer
	headers http.Header
	client  *runtimeAPIClient
}

// success sends the response payload for an in-progress invocation.
// Notes:
//   - An invoke is not complete until next() is called again!
func (i *invoke) success(body io.Reader, contentType string) error {
	_ = "STUB: not implemented"
	return nil
}

// failure sends the payload to the Runtime API. This marks the function's invoke as a failure.
// Notes:
//   - The execution of the function process continues, and is billed, until next() is called again!
//   - A Lambda Function continues to be re-used for future invokes even after a failure.
//     If the error is fatal (panic, unrecoverable state), exit the process immediately after calling failure()
func (i *invoke) failure(body io.Reader, contentType string, causeForXRay []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// next connects to the Runtime API and waits for a new invoke Request to be available.
// Note: After a call to Done() or Error() has been made, a call to next() will complete the in-flight invoke.
func (c *runtimeAPIClient) next(ctx context.Context) (*invoke, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *runtimeAPIClient) post(url string, body io.Reader, contentType string, xrayErrorCause []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func newErrorCapturingReader(r io.Reader) *errorCapturingReader {
	_ = "STUB: not implemented"
	return nil
}

type errorCapturingReader struct {
	reader  io.Reader
	Trailer http.Header
}

func (r *errorCapturingReader) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
