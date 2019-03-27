package main

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"bytes"
	"os"
)

func upload(ctxt context.Context, p string) error {
	creds := credentials.NewStaticCredentials(screenConfig.Aws.ID, screenConfig.Aws.Secret, screenConfig.Aws.Token)
	_, err := creds.Get()
	if err != nil {
		return err
	}

	cfg := aws.NewConfig().WithRegion(screenConfig.Aws.Region).WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	file, err := os.Open(p)
	if err != nil {
		return err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	size := fileInfo.Size()
	buffer := make([]byte, size)
	fileBytes := bytes.NewReader(buffer)

	_, err = file.Read(buffer)
	if err != nil {
		return err
	}

	_, err = svc.PutObjectWithContext(ctxt, &s3.PutObjectInput{
		Bucket: aws.String(screenConfig.Aws.Bucket),
		Key:    aws.String(p),
		Body:   fileBytes,
	})

	return err
}
