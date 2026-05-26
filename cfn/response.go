// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.

package cfn

import (

	//nolint: staticcheck

	"net/http"
)

// StatusType represents a CloudFormation response status
type StatusType string

const (
	StatusSuccess StatusType = "SUCCESS"
	StatusFailed  StatusType = "FAILED"
)

// Response is a representation of a Custom Resource
// response expected by CloudFormation.
type Response struct {
	Status             StatusType             `json:"Status"`
	RequestID          string                 `json:"RequestId"`
	LogicalResourceID  string                 `json:"LogicalResourceId"`
	StackID            string                 `json:"StackId"`
	PhysicalResourceID string                 `json:"PhysicalResourceId"`
	Reason             string                 `json:"Reason,omitempty"`
	NoEcho             bool                   `json:"NoEcho,omitempty"`
	Data               map[string]interface{} `json:"Data,omitempty"`

	url string
}

// NewResponse creates a Response with the relevant verbatim copied
// data from a Event
func NewResponse(r *Event) *Response { _ = "STUB: not implemented"; return nil }

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func (r *Response) sendWith(client httpClient) error { _ = "STUB: not implemented"; return nil }

// Send will send the Response to the given URL using the
// default HTTP client
func (r *Response) Send() error { _ = "STUB: not implemented"; return nil }
