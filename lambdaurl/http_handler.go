//go:build go1.18
// +build go1.18

// Copyright 2023 Amazon.com, Inc. or its affiliates. All Rights Reserved.

// Package lambdaurl serves requests from Lambda Function URLs using http.Handler.
package lambdaurl

import (
	"context"
	"io"
	"net/http"
	"sync"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type detectContentTypeContextKey struct{}

// WithDetectContentType sets the behavior of content type detection when the Content-Type header is not already provided.
// When true, the first Write call will pass the intial bytes to http.DetectContentType.
// When false, and if no Content-Type is provided, no Content-Type will be sent back to Lambda,
// and the Lambda Function URL will fallback to it's default.
//
// Note: The http.ResponseWriter passed to the handler is unbuffered.
// This may result in different Content-Type headers in the Function URL response when compared to http.ListenAndServe.
//
// Usage:
//
//	lambdaurl.Start(
//	        http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
//	                w.Write("<!DOCTYPE html><html></html>")
//	        }),
//	        lambdaurl.WithDetectContentType(true)
//	)
func WithDetectContentType(detectContentType bool) lambda.Option {
	_ = "STUB: not implemented"
	return *new(lambda.Option)
}

type httpResponseWriter struct {
	detectContentType bool
	header            http.Header
	writer            io.Writer
	once              sync.Once
	ready             chan<- header
}

type header struct {
	code   int
	header http.Header
}

func (w *httpResponseWriter) Header() http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

func (w *httpResponseWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *httpResponseWriter) WriteHeader(statusCode int) { _ = "STUB: not implemented"; return }

func (w *httpResponseWriter) writeHeader(statusCode int, initialPayload []byte) {
	_ = "STUB: not implemented"
	return
}

func detectContentType(p []byte) string {
	_ = "STUB: not implemented"
	// http.DetectContentType returns "text/plain; charset=utf-8" for nil and zero-length byte slices.
	// This is a weird behavior, since otherwise it defaults to "application/octet-stream"! So we'll do that.
	// This differs from http.ListenAndServe, which set no Content-Type when the initial Flush body is empty.
	return ""
}

type requestContextKey struct{}

// RequestFromContext returns the *events.LambdaFunctionURLRequest from a context.
func RequestFromContext(ctx context.Context) (*events.LambdaFunctionURLRequest, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Wrap converts an http.Handler into a Lambda request handler.
//
// Only Lambda Function URLs configured with `InvokeMode: RESPONSE_STREAM` are supported with the returned handler.
// The response body of the handler will conform to the content-type `application/vnd.awslambda.http-integration-response`.
func Wrap(handler http.Handler) func(context.Context, *events.LambdaFunctionURLRequest) (*events.LambdaFunctionURLStreamingResponse, error) {
	_ = "STUB: not implemented"
	return nil
}

// Signals when it's OK to start returning the response body to Lambda

// TODO: recover and CloseWithError the any panic value once the runtime API client supports plumbing fatal errors through the reader
//nolint:errcheck
// force default status, headers, content type detection, if none occurred during the execution of the handler

// Start wraps a http.Handler and calls lambda.StartHandlerFunc
// Only supports:
//   - Lambda Function URLs configured with `InvokeMode: RESPONSE_STREAM`
//   - Lambda Functions using the `provided` or `provided.al2` runtimes.
//   - Lambda Functions using the `go1.x` runtime when compiled with `-tags lambda.norpc`
func Start(handler http.Handler, options ...lambda.Option) { _ = "STUB: not implemented"; return }
