package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"bytes"
	"log"
	"os"
)

func upload(p string) {
	creds := credentials.NewStaticCredentials(screenConfig.Aws.Id, screenConfig.Aws.Secret, screenConfig.Aws.Token)
	_, err := creds.Get()
	if err != nil {
		log.Fatal(err)
	}

	cfg := aws.NewConfig().WithRegion(screenConfig.Aws.Region).WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	file, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	fileBytes := bytes.NewReader(buffer)

	file.Read(buffer)

	params := &s3.PutObjectInput{
		Bucket: aws.String(screenConfig.Aws.Bucket),
		Key:    aws.String(p),
		Body:   fileBytes,
	}

	_, err = svc.PutObject(params)
	if err != nil {
		log.Fatal(err)
	}
}
