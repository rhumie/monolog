package services

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/pkg/errors"
)

var sess = session.Must(session.NewSession())

// CloudwatchService provides the API operation methods for making requests to
// Amazon CloudWatch Logs.
type CloudwatchService struct {
}

// SearchLogEvents lists the specified log events.
func (s *CloudwatchService) SearchLogEvents(req *SearchLogEventRequest) (*SearchLogEventResponse, error) {
	// Create cloudwatch logs client.
	client := s.createAWSClient(&req.CommonRequest)
	if req.FilterPattern == "" {
		return s.searchLogEvents(client, req)
	}
	return s.searchLogEventsWithFilter(client, req)
}

// SearchLogEvent lists the specified log events.
func (s *CloudwatchService) searchLogEvents(client *cloudwatchlogs.CloudWatchLogs, req *SearchLogEventRequest) (*SearchLogEventResponse, error) {
	// Create GetLogEventsInput.
	input := new(cloudwatchlogs.GetLogEventsInput).
		SetLogGroupName(req.LogGroupName).
		SetLogStreamName(req.LogStreamName).
		SetStartFromHead(req.StartFromHead)
	if req.NextToken != "" {
		input.SetNextToken(req.NextToken)
	}
	if req.StartTime != 0 {
		input.SetStartTime(req.StartTime)
	}
	if req.EndTime != 0 {
		input.SetEndTime(req.EndTime)
	}
	fmt.Println(input)
	// Execute API.
	o, err := client.GetLogEvents(input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case credentials.ErrNoValidProvidersFoundInChain.Code():
				return nil, newAuthenticationError(err)
			case sts.ErrCodeExpiredTokenException:
				return nil, newAuthenticationError(err)
			}

		}
		return nil, errors.Wrap(err, "Failed to execute aws cloudwatchlogs GetLogEvents API.")
	}

	// Create response.
	res := &SearchLogEventResponse{
		LogEvents:         []*LogEvent{},
		NextBackwardToken: aws.StringValue(o.NextBackwardToken),
		NextForwardToken:  aws.StringValue(o.NextForwardToken),
	}
	for _, v := range o.Events {
		res.LogEvents = append(res.LogEvents, &LogEvent{
			Message:   aws.StringValue(v.Message),
			Timestamp: aws.Int64Value(v.Timestamp),
		})
	}
	return res, nil
}

// SearchLogEvent lists the specified log events.
func (s *CloudwatchService) searchLogEventsWithFilter(client *cloudwatchlogs.CloudWatchLogs, req *SearchLogEventRequest) (*SearchLogEventResponse, error) {
	// Create GetLogEventsInput.
	input := new(cloudwatchlogs.FilterLogEventsInput).
		SetLogGroupName(req.LogGroupName).
		SetLogStreamNames([]*string{&req.LogStreamName}).
		SetFilterPattern(req.FilterPattern)
	if req.NextToken != "" {
		input.SetNextToken(req.NextToken)
	}
	if req.StartTime != 0 {
		input.SetStartTime(req.StartTime)
	}
	if req.EndTime != 0 {
		input.SetEndTime(req.EndTime)
	}
	// Execute API.
	o, err := client.FilterLogEvents(input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case credentials.ErrNoValidProvidersFoundInChain.Code():
				return nil, newAuthenticationError(err)
			case sts.ErrCodeExpiredTokenException:
				return nil, newAuthenticationError(err)
			}

		}
		return nil, errors.Wrap(err, "Failed to execute aws cloudwatchlogs FilterLogEvents API.")
	}

	// Create response.
	res := &SearchLogEventResponse{
		LogEvents:        []*LogEvent{},
		NextForwardToken: aws.StringValue(o.NextToken),
	}
	for _, v := range o.Events {
		res.LogEvents = append(res.LogEvents, &LogEvent{
			Message:   aws.StringValue(v.Message),
			Timestamp: aws.Int64Value(v.Timestamp),
		})
	}

	return res, nil
}

// SearchLogs lists the all log groups and log streams.
func (s *CloudwatchService) SearchLogs(req *SearchLogRequest) (*SearchLogResponse, error) {
	// Create cloudwatch logs client.
	client := s.createAWSClient(&req.CommonRequest)
	// Search Log Groups.
	resLogGroups, err := s.searchLogGroups(client)
	if err != nil {
		return nil, err
	}
	// Search Log Streams.
	for _, v := range resLogGroups {
		v.LogStreams, err = s.searchLogStreams(client, v.Name)
		if err != nil {
			return nil, err
		}
	}
	// Create response.
	res := &SearchLogResponse{
		LogGroups: resLogGroups,
	}
	return res, nil
}

func (s *CloudwatchService) searchLogGroups(client *cloudwatchlogs.CloudWatchLogs) (res []*LogGroup, err error) {
	res = []*LogGroup{}
	err = client.DescribeLogGroupsPages(
		new(cloudwatchlogs.DescribeLogGroupsInput),
		// Function with the response data for each page.
		func(o *cloudwatchlogs.DescribeLogGroupsOutput, isLast bool) bool {
			for _, v := range o.LogGroups {
				g := &LogGroup{
					ID:   aws.StringValue(v.Arn),
					Name: aws.StringValue(v.LogGroupName),
				}
				res = append(res, g)
			}
			return isLast
		})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case credentials.ErrNoValidProvidersFoundInChain.Code():
				return nil, newAuthenticationError(err)
			case sts.ErrCodeExpiredTokenException:
				return nil, newAuthenticationError(err)
			}
		}
		return nil, errors.Wrap(err, "Faild to execute aws cloudwatchlogs DescribeLogGroupsInput API.")
	}
	return
}

func (s *CloudwatchService) searchLogStreams(client *cloudwatchlogs.CloudWatchLogs, logGroupName string) (res []*LogStream, err error) {
	res = []*LogStream{}
	err = client.DescribeLogStreamsPages(
		new(cloudwatchlogs.DescribeLogStreamsInput).SetLogGroupName(logGroupName),
		// Function with the response data for each page.
		func(o *cloudwatchlogs.DescribeLogStreamsOutput, isLast bool) bool {
			for _, v := range o.LogStreams {
				s := &LogStream{
					ID:                  aws.StringValue(v.Arn),
					Name:                aws.StringValue(v.LogStreamName),
					FirstEventTimestamp: aws.Int64Value(v.FirstEventTimestamp),
					LastEventTimestamp:  aws.Int64Value(v.LastEventTimestamp),
				}
				res = append(res, s)
			}
			return isLast
		})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case credentials.ErrNoValidProvidersFoundInChain.Code():
				return nil, newAuthenticationError(err)
			case sts.ErrCodeExpiredTokenException:
				return nil, newAuthenticationError(err)
			}
		}
		return nil, errors.Wrap(err, "Faild to execute aws cloudwarchlogs DescribeLogStreamsPages API.")
	}
	return
}

func (s *CloudwatchService) createAWSClient(req *CommonRequest) *cloudwatchlogs.CloudWatchLogs {
	// Create aws configuration.
	config := new(aws.Config).WithRegion(req.AWSRegion)
	if req.SessionToken != "" {
		cred := strings.SplitN(req.SessionToken, ":", 3)
		config.WithCredentials(credentials.NewStaticCredentials(cred[0], cred[1], cred[2]))
	}
	// Create cloudwatch logs client.
	return cloudwatchlogs.New(sess, config)
}
