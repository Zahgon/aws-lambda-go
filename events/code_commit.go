package events

import (
	"time"
)

// CodeCommitEvent represents a CodeCommit event
type CodeCommitEvent struct {
	Records []CodeCommitRecord `json:"Records"`
}

// String returns a string representation of this object.
// Useful for testing and debugging.
func (e CodeCommitEvent) String() string { _ = "STUB: not implemented"; return "" }

type CodeCommitEventTime time.Time

// https://golang.org/pkg/time/#Parse
const codeCommitEventTimeReference = "\"2006-01-2T15:04:05.000-0700\""

func (t *CodeCommitEventTime) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *CodeCommitEventTime) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// CodeCommitRecord represents a CodeCommit record
type CodeCommitRecord struct {
	EventID              string               `json:"eventId"`
	EventVersion         string               `json:"eventVersion"`
	EventTime            CodeCommitEventTime  `json:"eventTime"`
	EventTriggerName     string               `json:"eventTriggerName"`
	EventPartNumber      uint64               `json:"eventPartNumber"`
	CodeCommit           CodeCommitCodeCommit `json:"codecommit"`
	EventName            string               `json:"eventName"`
	EventTriggerConfigId string               `json:"eventTriggerConfigId"` //nolint: staticcheck
	EventSourceARN       string               `json:"eventSourceARN"`
	UserIdentityARN      string               `json:"userIdentityARN"`
	EventSource          string               `json:"eventSource"`
	AWSRegion            string               `json:"awsRegion"`
	EventTotalParts      uint64               `json:"eventTotalParts"`
	CustomData           string               `json:"customData,omitempty"`
}

// String returns a string representation of this object.
// Useful for testing and debugging.
func (r CodeCommitRecord) String() string { _ = "STUB: not implemented"; return "" }

// CodeCommitCodeCommit represents a CodeCommit object in a record
type CodeCommitCodeCommit struct {
	References []CodeCommitReference `json:"references"`
}

// String returns a string representation of this object.
// Useful for testing and debugging.
func (c CodeCommitCodeCommit) String() string { _ = "STUB: not implemented"; return "" }

// CodeCommitReference represents a Reference object in a CodeCommit object
type CodeCommitReference struct {
	Commit  string `json:"commit"`
	Ref     string `json:"ref"`
	Created bool   `json:"created,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}

// String returns a string representation of this object.
// Useful for testing and debugging.
func (r CodeCommitReference) String() string { _ = "STUB: not implemented"; return "" }
