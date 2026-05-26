package lambda

import (
	"net/http"
)

const (
	headerExtensionName       = "Lambda-Extension-Name"
	headerExtensionIdentifier = "Lambda-Extension-Identifier"
	extensionAPIVersion       = "2020-01-01"
)

type extensionAPIEventType string

const (
	extensionInvokeEvent   extensionAPIEventType = "INVOKE"   //nolint:deadcode,unused,varcheck
	extensionShutdownEvent extensionAPIEventType = "SHUTDOWN" //nolint:deadcode,unused,varcheck
)

type extensionAPIClient struct {
	baseURL    string
	httpClient *http.Client
}

func newExtensionAPIClient(address string) *extensionAPIClient {
	_ = "STUB: not implemented"
	return nil
}

// connections to the extensions API are never expected to time out

func (c *extensionAPIClient) register(name string, events ...extensionAPIEventType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type extensionEventResponse struct {
	EventType extensionAPIEventType
	// ... the rest not implemented
}

func (c *extensionAPIClient) next(id string) (response extensionEventResponse, err error) {
	_ = "STUB: not implemented"
	return *new(extensionEventResponse), nil
}
