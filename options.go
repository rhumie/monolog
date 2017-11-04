package main

import kingpin "gopkg.in/alecthomas/kingpin.v2"

// Options hogehoge
type options struct {
	port    *int
	service *string
}

// Port Hogehoge
func (o *options) Port() int {
	return *o.port
}

func (o *options) Service() string {
	return *o.service
}

var opts = parse()

func parse() *options {
	o := &options{
		port:    kingpin.Flag("port", "Server port").Short('p').Default("8000").Int(),
		service: kingpin.Flag("service", "Log service").Short('s').Default(awsCloudWatchLogs).Enum(awsCloudWatchLogs, awsS3, elasticsearch),
	}
	kingpin.Parse()
	return o
}
