package test

import (
	"api/config"
	"api/config/configmocks"
	"api/lib"
	"api/lib/aws"
	"api/lib/aws/awsmocks"
	"api/lib/libmocks"
	"errors"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("blob_storage", func() {
	var mockAwsConfig configmocks.AwsConfig
	var mockBlobStorageConfig *lib.BlobStorageConfig
	var mockS3Client awsmocks.S3Client
	var mockLogger libmocks.Logger
	var subject lib.BlobStorage

	var originalNewS3Client aws.NewS3ClientFunc
	var originalNewLogger lib.NewLoggerFunc

	BeforeEach(func() {
		mockAwsConfig = configmocks.AwsConfig{}
		mockAwsConfig.On("GetAccessKeyId").Return("mock-access-key-id")
		mockAwsConfig.On("GetSecretAccessKey").Return("mock-secret-access-key")
		mockBlobStorageConfig = &lib.BlobStorageConfig{
			AwsConfig: &mockAwsConfig,
		}

		mockS3Client = awsmocks.S3Client{}
		originalNewS3Client = aws.NewS3Client
		aws.NewS3Client = func(_ config.AwsConfig) aws.S3Client {
			return &mockS3Client
		}

		mockLogger = libmocks.Logger{}
		mockLogger.On("LogError", mock.Anything)
		originalNewLogger = lib.NewLogger
		lib.NewLogger = func(_name string) lib.Logger {
			return &mockLogger
		}
	})

	AfterEach(func() {
		aws.NewS3Client = originalNewS3Client
		lib.NewLogger = originalNewLogger
	})

	Describe("S3BlobStorage", func() {
		BeforeEach(func() {
			subject = lib.NewBlobStorage(mockBlobStorageConfig)
		})

		It("NewDefaultBlobStorage creates a new S3 blob storage instance", func() {
			Expect(subject).To(BeAssignableToTypeOf(&lib.S3BlobStorage{}))
		})

		Describe("GetSignedExternalMediaUploadUrl", func() {
			var expectedReq *aws.S3PutObjectRequest
			BeforeEach(func() {
				expectedReq = &aws.S3PutObjectRequest{
					Bucket: "dokodoko-static-external",
					Key:    "expected-key",
					Md5:    "expected-md5",
					Ttl:    5 * time.Minute,
				}
				mockS3Client.On("CreateSignedPutObjectUrl", mock.MatchedBy(func(req *aws.S3PutObjectRequest) bool {
					return req.Bucket == expectedReq.Bucket &&
						req.Key == expectedReq.Key &&
						req.Md5 == expectedReq.Md5 &&
						req.Ttl == expectedReq.Ttl
				})).Return("http://expected-upload-url", nil)
				mockS3Client.On("CreateSignedPutObjectUrl", mock.MatchedBy(func(req *aws.S3PutObjectRequest) bool {
					return req.Key == "error-key"
				})).Return("", errors.New("actual-error"))
			})

			Context("AWS client works as expected", func() {
				It("returns the upload url", func() {
					actual, err := subject.GetSignedExternalMediaUploadUrl(&lib.BlobUploadSpec{
						Key: "expected-key",
						Md5: "expected-md5",
					})
					Expect(err).To(BeNil())
					Expect(actual).To(Equal("http://expected-upload-url"))
				})
			})

			Context("AWS client returns error", func() {
				It("returns empty string and intl error", func() {
					actual, err := subject.GetSignedExternalMediaUploadUrl(&lib.BlobUploadSpec{
						Key: "error-key",
					})
					Expect(err).To(MatchError(&lib.UnableToCreateUploadUrlError{}))
					Expect(actual).To(Equal(""))
					Expect(mockLogger.AssertCalled(
						TheT,
						"LogError",
						errors.New("actual-error"),
					)).To(BeTrue())
				})
			})
		})

		Describe("GetSignedExternalMediaDownloadUrl", func() {
			var expectedReq *aws.S3GetObjectRequest
			BeforeEach(func() {
				expectedReq = &aws.S3GetObjectRequest{
					Bucket: "dokodoko-static-external",
					Key:    "expected-key",
					Ttl:    5 * time.Minute,
				}
				mockS3Client.On("CreateSignedGetObjectUrl", mock.MatchedBy(func(req *aws.S3GetObjectRequest) bool {
					return req.Bucket == expectedReq.Bucket &&
						req.Key == expectedReq.Key &&
						req.Ttl == expectedReq.Ttl
				})).Return("http://expected-download-url", nil)
				mockS3Client.On("CreateSignedGetObjectUrl", mock.MatchedBy(func(req *aws.S3GetObjectRequest) bool {
					return req.Key == "error-key"
				})).Return("", errors.New("actual-error"))
			})

			Context("AWS client works as expected", func() {
				It("returns the download url", func() {
					actual, err := subject.GetSignedExternalMediaDownloadUrl("expected-key")
					Expect(err).To(BeNil())
					Expect(actual).To(Equal("http://expected-download-url"))
				})
			})

			Context("AWS client returns error", func() {
				It("returns empty string and intl error", func() {
					actual, err := subject.GetSignedExternalMediaDownloadUrl("error-key")
					Expect(err).To(MatchError(&lib.UnableToCreateDownloadUrlError{}))
					Expect(actual).To(Equal(""))
					Expect(mockLogger.AssertCalled(
						TheT,
						"LogError",
						errors.New("actual-error"),
					)).To(BeTrue())
				})
			})
		})

		Describe("DeleteExternalMedia", func() {
			BeforeEach(func() {
				mockS3Client.On("DeleteObject", mock.MatchedBy(func(req *aws.S3DeleteObjectRequest) bool {
					return req.Key == "success-key"
				})).Return(nil)
				mockS3Client.On("DeleteObject", mock.MatchedBy(func(req *aws.S3DeleteObjectRequest) bool {
					return req.Key == "error-key"
				})).Return(errors.New("actual-error"))
			})

			Context("aws client delete object succeeded", func() {
				It("returns nil error", func() {
					err := subject.DeleteExternalMedia("success-key")
					Expect(err).To(BeNil())
				})
			})

			Context("aws client returns error", func() {
				It("returns UnableToDeleteError", func() {
					err := subject.DeleteExternalMedia("error-key")
					Expect(err).To(MatchError(&lib.UnableToDeleteError{
						Bucket: "dokodoko-static-external",
						Key:    "error-key",
					}))
					Expect(mockLogger.AssertCalled(
						TheT,
						"LogError",
						errors.New("actual-error"),
					)).To(BeTrue())
				})
			})
		})
	})
})
