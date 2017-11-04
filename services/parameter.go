package services

// CommonRequest is common request parameters.
type CommonRequest struct {
	AWSRegion    string `json:"awsRegion" form:"awsRegion" query:"awsRegion"`
	SessionToken string
}

// SearchLogRequest is SearchLog request paramters.
type SearchLogRequest struct {
	CommonRequest
}

// SearchLogResponse is SearchLog response paramters.
type SearchLogResponse struct {
	LogGroups []*LogGroup `json:"logGroups"`
}

// SearchLogEventRequest is SearchLogEvent request paramters
type SearchLogEventRequest struct {
	CommonRequest
	LogGroupName  string `json:"logGroupName" form:"logGroupName" query:"logGroupName"`
	LogStreamName string `json:"logStreamName" form:"logStreamName" query:"logStreamName"`
	StartTime     int64  `json:"startTime" form:"startTime" query:"startTime"`
	EndTime       int64  `json:"endTime" form:"endTime" query:"endTime"`
	NextToken     string `json:"nextToken" form:"nextToken" query:"nextToken"`
	StartFromHead bool   `json:"startFromHead" form:"startFromHead" query:"startFromHead"`
	FilterPattern string `json:"filterPattern" form:"filterPattern" query:"filterPattern"`
}

// SearchLogEventResponse is SearchLogEvent response paramters
type SearchLogEventResponse struct {
	LogEvents         []*LogEvent `json:"logEvents"`
	NextBackwardToken string      `json:"nextBackwardToken"`
	NextForwardToken  string      `json:"nextForwardToken"`
}

// LogGroup is log group parameters.
type LogGroup struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	LogStreams []*LogStream `json:"logStreams"`
}

// LogStream is log stream parameters.
type LogStream struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	FirstEventTimestamp int64  `json:"firstEventTimestamp"`
	LastEventTimestamp  int64  `json:"lastEventTimestamp"`
}

// LogEvent is log event parameters.
type LogEvent struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}
