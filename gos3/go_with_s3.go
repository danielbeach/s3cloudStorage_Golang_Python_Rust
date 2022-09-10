package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	defer timer("main")()
	os.Setenv("AWS_ACCESS_KEY_ID", "")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "")
	var sess = get_aws_session()
	svc := s3.New(sess)
	rsp := list_bucket(svc, "202201")
	fmt.Println(rsp)
}

func get_aws_session() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		fmt.Println(err)
	}
	return sess
}

func list_bucket(svc *s3.S3, search_string string) *s3.ListObjectsV2Output {
	bucket := "confessions-of-a-data-guy"
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket),
		Prefix: aws.String(search_string)})
	if err != nil {
		fmt.Println(err)
	}
	return resp
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
