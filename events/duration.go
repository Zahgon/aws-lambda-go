package events

import (
	"time"
)

type DurationSeconds time.Duration

// UnmarshalJSON converts a given json to a DurationSeconds
func (duration *DurationSeconds) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// MarshalJSON converts a given DurationSeconds to json
func (duration DurationSeconds) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DurationMinutes time.Duration

// UnmarshalJSON converts a given json to a DurationMinutes
func (duration *DurationMinutes) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// MarshalJSON converts a given DurationMinutes to json
func (duration DurationMinutes) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
