// Copyright 2022 Amazon.com, Inc. or its affiliates. All Rights Reserved.

package lambda

// enableSIGTERM configures an optional list of sigtermHandlers to run on process shutdown.
// This non-default behavior is enabled within Lambda using the extensions API.
func enableSIGTERM(sigtermHandlers []func()) {
	_ = "STUB: not implemented"
	// for fun, we'll also optionally register SIGTERM handlers
	return
}

// detect if we're actually running within Lambda

// Now to do the AWS Lambda specific stuff.
// The default Lambda behavior is for functions to get SIGKILL at the end of lifetime, or after a timeout.
// Any use of the Lambda extension register API enables SIGTERM to be sent to the function process before the SIGKILL.
// We'll register an extension that does not listen for any lifecycle events named "GoLangEnableSIGTERM".
// The API will respond with an ID we need to pass in future requests.

// We didn't actually register for any events, but we need to call /next anyways to let the API know we're done initalizing.
// Because we didn't register for any events, /next will never return, so we'll do this in a go routine that is doomed to stay blocked.
