package main

// import (
// 	"fmt"
//
// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	cwl "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
// )
//
// func main() {
// 	s := session.Must(session.NewSession())
// 	fmt.Println("hoge")
// 	c := cwl.New(s, new(aws.Config).WithRegion("ap-northeast-1"))
// 	c.Config.WithRegion("region")
//
// 	cwl.New(c)
// 	fmt.Println("hoge")
// 	p, err := c.DescribeLogGroups(new(cwl.DescribeLogGroupsInput))
// 	fmt.Println(p)
// 	fmt.Println(err)
// }
