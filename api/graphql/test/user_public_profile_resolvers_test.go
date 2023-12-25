package test

import (
	"api/config"
	"api/config/configmocks"
	"api/ent/entgenerated"
	"api/ent/privacy/entviewer"
	"api/lib"
	"api/lib/libmocks"
	"api/lib/testutils"
	"api/privacy/viewer"
	"context"

	"github.com/99designs/gqlgen/client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("user_public_profile.resolvers", func() {
	var entClient *entgenerated.Client
	var mockBlobStorage *libmocks.BlobStorage
	var graphqlClient *client.Client
	var owner *entgenerated.User
	var ownerCtx context.Context

	var mockGetHmacConfig testutils.MockWrapper[config.GetHmacConfigFunc]
	var mockGetAwsConfig testutils.MockWrapper[config.GetAwsConfigFunc]
	var mockNewBlobStorage testutils.MockWrapper[lib.NewBlobStorageFunc]

	BeforeEach(func() {
		mockGetHmacConfig = testutils.StartMockWrapper[config.GetHmacConfigFunc](
			func() config.HmacConfig {
				mockConfig := configmocks.HmacConfig{}
				mockConfig.On("GetRefreshTokenSecret").Return("refresh_token_secret")
				mockConfig.On("GetAuthTokenSecret").Return("auth_token_secret")
				return &mockConfig
			},
			config.GetHmacConfig,
			func(aGet config.GetHmacConfigFunc) {
				config.GetHmacConfig = aGet
			},
		)

		mockGetAwsConfig = testutils.StartMockWrapper[config.GetAwsConfigFunc](
			func() config.AwsConfig {
				mockAwsConfig := configmocks.AwsConfig{}
				mockAwsConfig.On("GetAccessKeyId").Return("mock-access-key-id")
				mockAwsConfig.On("GetSecretAccessKey").Return("mock-secret-access-key")
				return &mockAwsConfig
			},
			config.GetAwsConfig,
			func(aGet config.GetAwsConfigFunc) {
				config.GetAwsConfig = aGet
			},
		)

		mockBlobStorage = &libmocks.BlobStorage{}
		mockBlobStorage.On("GetSignedExternalMediaDownloadUrl", "correct-key").Return("http://expected-download-url", nil)
		mockBlobStorage.On("GetSignedExternalMediaDownloadUrl", "error-key").Return("", &lib.UnableToCreateDownloadUrlError{})
		mockNewBlobStorage = testutils.StartMockWrapper[lib.NewBlobStorageFunc](
			func(_ *lib.BlobStorageConfig) lib.BlobStorage {
				return mockBlobStorage
			},
			lib.NewBlobStorage,
			func(aNew lib.NewBlobStorageFunc) {
				lib.NewBlobStorage = aNew
			},
		)

		entClient = createEntClientForTest()
		graphqlClient = createGraphqlClient(entClient)

		owner = createTestUser(entClient)
		ownerCtx = viewer.NewContext(context.Background(), entviewer.NewUserViewerFromUser(owner))

		// Old

	})

	AfterEach(func() {
		entClient.Close()
		mockGetHmacConfig.Stop()
		mockNewBlobStorage.Stop()
		mockGetAwsConfig.Stop()
	})

	Describe("PhotoDownloadUrl", func() {
		type UserProfilePhotoQuery struct {
			User *struct {
				PublicProfile *struct {
					PhotoDownloadUrl *string `json:"photoDownloadUrl"`
				} `json:"publicProfile"`
			} `json:"user"`
		}

		doPhotoDownloadUrlQuery := func() (*string, error) {
			var response UserProfilePhotoQuery
			err := graphqlClient.Post(
				`
				query UserPhotoDownloadUrlQuery($id: ID!) {
					user(id: $id) {
						publicProfile {
							photoDownloadUrl
						}
					}
				}
				`,
				&response,
				client.Var("id", owner.ID),
				requestWithContext(
					func(parent context.Context) context.Context {
						return viewer.NewContext(parent, entviewer.NewAnonymouseUserViewer())
					},
				),
			)
			return response.User.PublicProfile.PhotoDownloadUrl, err
		}

		Context("no photo blob", func() {
			BeforeEach(func() {
				entClient.UserPublicProfile.Create().SetOwner(owner).SetHandleName("handle-name").SaveX(ownerCtx)
			})

			It("returns nil", func() {
				actual, err := doPhotoDownloadUrlQuery()
				Expect(err).To(BeNil())
				Expect(actual).To(BeNil())
			})
		})

		Context("has photo blob and can generate URL", func() {
			BeforeEach(func() {
				entClient.UserPublicProfile.Create().SetOwner(owner).SetHandleName("handle-name").SetPhotoBlobKey("correct-key").SaveX(ownerCtx)
			})

			It("returns the expected download URL", func() {
				actual, err := doPhotoDownloadUrlQuery()
				Expect(err).To(BeNil())
				expected := "http://expected-download-url"
				Expect(actual).To(Equal(&expected))
			})
		})

		Context("blob storage returns error", func() {
			BeforeEach(func() {
				entClient.UserPublicProfile.Create().SetOwner(owner).SetHandleName("handle-name").SetPhotoBlobKey("error-key").SaveX(ownerCtx)
			})

			It("returns nil and error", func() {
				actual, err := doPhotoDownloadUrlQuery()
				Expect(err.Error()).To(Equal("[{\"message\":\"blob-storage.unable-to-create-download-url\",\"path\":[\"user\",\"publicProfile\",\"photoDownloadUrl\"]}]"))
				Expect(actual).To(BeNil())
			})
		})
	})
})
