// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

package events

import (
	"time"
)

// RFC3339EpochTime serializes a time.Time in JSON as an ISO 8601 string.
type RFC3339EpochTime struct {
	time.Time
}

// SecondsEpochTime serializes a time.Time in JSON as a UNIX epoch time in seconds
type SecondsEpochTime struct {
	time.Time
}

// MilliSecondsEpochTime serializes a time.Time in JSON as a UNIX epoch time in milliseconds.
type MilliSecondsEpochTime struct {
	time.Time
}

const secondsToNanoSecondsFactor = 1000000000
const milliSecondsToNanoSecondsFactor = 1000000

func (e SecondsEpochTime) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// UnixNano() returns the epoch in nanoseconds
	return nil, nil
}

func (e *SecondsEpochTime) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// time.Unix(sec, nsec) expects the epoch integral seconds in the first parameter
// and remaining nanoseconds in the second parameter

func (e MilliSecondsEpochTime) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// UnixNano() returns the epoch in nanoseconds
	return nil, nil
}

func (e *MilliSecondsEpochTime) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e RFC3339EpochTime) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *RFC3339EpochTime) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
