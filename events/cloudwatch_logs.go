package events

// CloudwatchLogsEvent represents raw data from a cloudwatch logs event
type CloudwatchLogsEvent struct {
	AWSLogs CloudwatchLogsRawData `json:"awslogs"`
}

// CloudwatchLogsRawData contains gzipped base64 json representing the bulk
// of a cloudwatch logs event
type CloudwatchLogsRawData struct {
	Data string `json:"data"`
}

// Parse returns a struct representing a usable CloudwatchLogs event
func (c CloudwatchLogsRawData) Parse() (d CloudwatchLogsData, err error) {
	_ = "STUB: not implemented"
	return *new(CloudwatchLogsData), nil
}

// CloudwatchLogsData is an unmarshal'd, ungzip'd, cloudwatch logs event
type CloudwatchLogsData struct {
	Owner               string                   `json:"owner"`
	LogGroup            string                   `json:"logGroup"`
	LogStream           string                   `json:"logStream"`
	SubscriptionFilters []string                 `json:"subscriptionFilters"`
	MessageType         string                   `json:"messageType"`
	LogEvents           []CloudwatchLogsLogEvent `json:"logEvents"`
}

// CloudwatchLogsLogEvent represents a log entry from cloudwatch logs
type CloudwatchLogsLogEvent struct {
	ID        string `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}
