package services

// Service is interface of various log storage service.
type Service interface {
	SearchLogs(req *SearchLogRequest) (res *SearchLogResponse, err error)
	SearchLogEvents(req *SearchLogEventRequest) (res *SearchLogEventResponse, err error)
}
