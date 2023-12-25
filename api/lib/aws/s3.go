package aws

import (
	"api/config"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type (
	S3PutObjectRequest struct {
		Bucket string
		Key    string
		Md5    string
		Ttl    time.Duration
	}

	S3GetObjectRequest struct {
		Bucket string
		Key    string
		Ttl    time.Duration
	}

	S3DeleteObjectRequest struct {
		Bucket string
		Key    string
	}

	S3Client interface {
		CreateSignedPutObjectUrl(*S3PutObjectRequest) (string, error)
		CreateSignedGetObjectUrl(*S3GetObjectRequest) (string, error)
		DeleteObject(*S3DeleteObjectRequest) error
	}

	DefaultS3Client struct {
		S3Client
		config config.AwsConfig
	}

	NewS3ClientFunc = func(config.AwsConfig) S3Client
)

var NewS3Client NewS3ClientFunc = func(awsConfig config.AwsConfig) S3Client {
	return &DefaultS3Client{
		config: awsConfig,
	}
}

func (c DefaultS3Client) createAwsSession() (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Region:      aws.String(c.config.GetRegion()),
		Credentials: credentials.NewStaticCredentials(c.config.GetAccessKeyId(), c.config.GetSecretAccessKey(), ""),
	})
}

func (c DefaultS3Client) newS3() (*s3.S3, error) {
	session, err := c.createAwsSession()
	if err != nil {
		return nil, err
	}

	return s3.New(session), nil
}

func (c DefaultS3Client) CreateSignedPutObjectUrl(request *S3PutObjectRequest) (string, error) {
	svc, err := c.newS3()
	if err != nil {
		return "", err
	}

	resp, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket:     aws.String(request.Bucket),
		Key:        aws.String(request.Key),
		ContentMD5: aws.String(request.Md5),
	})

	url, err := resp.Presign(request.Ttl)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (c DefaultS3Client) CreateSignedGetObjectUrl(request *S3GetObjectRequest) (string, error) {
	svc, err := c.newS3()
	if err != nil {
		return "", err
	}

	resp, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(request.Bucket),
		Key:    aws.String(request.Key),
	})

	url, err := resp.Presign(request.Ttl)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (c DefaultS3Client) DeleteObject(request *S3DeleteObjectRequest) error {
	svc, err := c.newS3()
	if err != nil {
		return err
	}

	_, s3Err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(request.Bucket),
		Key:    aws.String(request.Key),
	})

	return s3Err
}
