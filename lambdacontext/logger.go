//go:build go1.21
// +build go1.21

// Copyright 2026 Amazon.com, Inc. or its affiliates. All Rights Reserved.

package lambdacontext

import (
	"context"
	"log/slog"
	"os"
)

// logFormat is the log format from AWS_LAMBDA_LOG_FORMAT (TEXT or JSON)
var logFormat = os.Getenv("AWS_LAMBDA_LOG_FORMAT")

// logLevel is the log level from AWS_LAMBDA_LOG_LEVEL
var logLevel = os.Getenv("AWS_LAMBDA_LOG_LEVEL")

// field represents a Lambda context field to include in log records.
type field struct {
	key   string
	value func(*LambdaContext) string
}

// logOptions holds configuration for the Lambda log handler.
type logOptions struct {
	fields []field
}

// LogOption is a functional option for configuring the Lambda log handler.
type LogOption func(*logOptions)

// WithFunctionARN includes the invoked function ARN in log records.
func WithFunctionARN() LogOption { _ = "STUB: not implemented"; return *new(LogOption) }

// WithTenantID includes the tenant ID in log records (for multi-tenant functions).
func WithTenantID() LogOption { _ = "STUB: not implemented"; return *new(LogOption) }

// NewLogHandler returns a [slog.Handler] for AWS Lambda structured logging.
// It reads AWS_LAMBDA_LOG_FORMAT and AWS_LAMBDA_LOG_LEVEL from environment,
// and injects requestId from Lambda context into each log record.
//
// By default, only requestId is injected. Use WithFunctionARN or WithTenantID to include more.
// See the package examples for usage.
func NewLogHandler(opts ...LogOption) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

// NewLogger returns a [*slog.Logger] configured for AWS Lambda structured logging.
// This is a convenience function equivalent to slog.New(NewLogHandler(opts...)).
func NewLogger(opts ...LogOption) *slog.Logger { _ = "STUB: not implemented"; return nil }

// ReplaceAttr maps slog's default keys to AWS Lambda's log format (time->timestamp, msg->message).
func ReplaceAttr(groups []string, attr slog.Attr) slog.Attr {
	_ = "STUB: not implemented"
	return *new(slog.Attr)
}

// lambdaHandler wraps a slog.Handler to inject Lambda context fields.
type lambdaHandler struct {
	handler slog.Handler
	fields  []field
}

// Enabled implements slog.Handler.
func (h *lambdaHandler) Enabled(ctx context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

// Handle implements slog.Handler.
func (h *lambdaHandler) Handle(ctx context.Context, r slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}

// WithAttrs implements slog.Handler.
func (h *lambdaHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

// WithGroup implements slog.Handler.
func (h *lambdaHandler) WithGroup(name string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func parseLogLevel() slog.Level { _ = "STUB: not implemented"; return *new(slog.Level) }
