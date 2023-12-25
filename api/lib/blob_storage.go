package lib

import (
	"api/config"
	"api/intl"
	"api/intl/intlgenerated"
	"api/lib/aws"
	"fmt"
	"time"
)

type (
	BlobUploadSpec struct {
		Key string
		Md5 string
	}

	BlobStorage interface {
		GetSignedExternalMediaUploadUrl(spec *BlobUploadSpec) (string, intl.IntlError)
		GetSignedExternalMediaDownloadUrl(key string) (string, intl.IntlError)
		DeleteExternalMedia(key string) intl.IntlError
	}

	BlobStorageConfig struct {
		AwsConfig config.AwsConfig
	}

	S3BlobStorage struct {
		BlobStorage
		config *BlobStorageConfig
		logger Logger
	}

	NewBlobStorageFunc func(config *BlobStorageConfig) BlobStorage

	UnableToCreateUploadUrlError struct {
		intl.IntlError
	}

	UnableToCreateDownloadUrlError struct {
		intl.IntlError
	}

	UnableToDeleteError struct {
		intl.IntlError
		Bucket string
		Key    string
	}
)

const (
	// These should be moved to a config
	S3_REGION                = "ap-northeast-1"
	S3_EXTERNAL_MEDIA_BUCKET = "dokodoko-static-external"

	EXTERNAL_MEDIA_UPLOAD_URL_TTL   = 5 * time.Minute
	EXTERNAL_MEDIA_DOWNLOAD_URL_TTL = 5 * time.Minute

	BLOB_STORAGE_LOGGER_NAME = "blob_storage"
)

var NewBlobStorage NewBlobStorageFunc = func(blobStorageConfig *BlobStorageConfig) BlobStorage {
	return &S3BlobStorage{
		config: blobStorageConfig,
		logger: NewLogger(BLOB_STORAGE_LOGGER_NAME),
	}
}

func (s S3BlobStorage) createAwsClient() aws.S3Client {
	return aws.NewS3Client(s.config.AwsConfig)
}

func (s S3BlobStorage) GetSignedExternalMediaUploadUrl(spec *BlobUploadSpec) (string, intl.IntlError) {
	client := s.createAwsClient()
	url, err := client.CreateSignedPutObjectUrl(&aws.S3PutObjectRequest{
		Bucket: S3_EXTERNAL_MEDIA_BUCKET,
		Key:    spec.Key,
		Md5:    spec.Md5,
		Ttl:    EXTERNAL_MEDIA_UPLOAD_URL_TTL,
	})
	if err != nil {
		s.logger.LogError(err)
		return "", &UnableToCreateUploadUrlError{}
	}

	return url, nil
}

func (UnableToCreateUploadUrlError) Error() string {
	return "unable to create upload url"
}

func (UnableToCreateUploadUrlError) GetIntlKey() string {
	return intlgenerated.BLOB_STORAGE__UNABLE_TO_CREATE_UPLOAD_URL
}

func (s S3BlobStorage) GetSignedExternalMediaDownloadUrl(key string) (string, intl.IntlError) {
	client := s.createAwsClient()
	url, err := client.CreateSignedGetObjectUrl(&aws.S3GetObjectRequest{
		Bucket: S3_EXTERNAL_MEDIA_BUCKET,
		Key:    key,
		Ttl:    EXTERNAL_MEDIA_DOWNLOAD_URL_TTL,
	})
	if err != nil {
		s.logger.LogError(err)
		return "", &UnableToCreateDownloadUrlError{}
	}
	return url, nil
}

func (UnableToCreateDownloadUrlError) Error() string {
	return "unable to create download url"
}

func (UnableToCreateDownloadUrlError) GetIntlKey() string {
	return intlgenerated.BLOB_STORAGE__UNABLE_TO_CREATE_DOWNLOAD_URL
}

func (s S3BlobStorage) DeleteExternalMedia(key string) intl.IntlError {
	client := s.createAwsClient()
	err := client.DeleteObject(&aws.S3DeleteObjectRequest{
		Bucket: S3_EXTERNAL_MEDIA_BUCKET,
		Key:    key,
	})
	if err != nil {
		s.logger.LogError(err)
		return &UnableToDeleteError{
			Bucket: S3_EXTERNAL_MEDIA_BUCKET,
			Key:    key,
		}
	}
	return nil
}

func (e UnableToDeleteError) Error() string {
	return fmt.Sprintf("unable to delete file: Bucket('%s'), Key('%s')", e.Bucket, e.Key)
}

func (UnableToDeleteError) GetIntlKey() string {
	return intlgenerated.BLOB_STORAGE__UNABLE_TO_DELETE_OBJECT
}
