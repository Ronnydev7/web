package test

import (
	"api/config"
	"api/config/configmocks"
	"api/ent/entgenerated"
	"api/ent/privacy/entviewer"
	"api/graphql/gqlgenerated"
	"api/intl"
	"api/lib"
	"api/lib/cookies"
	"api/lib/libmocks"
	"api/lib/matchers"
	"api/privacy/viewer"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("user.resolvers", func() {
	var currentTime time.Time
	var originalTimeNow func() time.Time

	var entClient *entgenerated.Client
	var graphqlClient *client.Client

	var user *entgenerated.User
	var loginSession *entgenerated.LoginSession

	var originalGetAwsConfig config.GetAwsConfigFunc
	var originalNewEntClient lib.NewEntClientFunc
	var originalNewLogger lib.NewLoggerFunc

	var mockCookieReader libmocks.CookieReader
	var mockHttpErrorWriter libmocks.HttpErrorWriter
	var mockLogger libmocks.Logger
	superCtx := viewer.NewContext(context.Background(), &testViewer{})

	BeforeEach(func() {
		currentTime = time.Now()
		originalTimeNow = lib.TimeNow
		lib.TimeNow = func() time.Time {
			return currentTime
		}

		originalGetAwsConfig = config.GetAwsConfig
		testAwsConfig := configmocks.AwsConfig{}
		testAwsConfig.On("GetRegion").Return("apac-1")
		testAwsConfig.On("GetAccessKeyId").Return("mock-access-key-id")
		testAwsConfig.On("GetSecretAccessKey").Return("mock-secret-access-key")
		config.GetAwsConfig = func() config.AwsConfig {
			return &testAwsConfig
		}

		entClient = createEntClientForTest()
		graphqlClient = createGraphqlClient(entClient)

		ctx := viewer.NewContext(context.Background(), testViewer{})
		user = entClient.User.Create().SaveX(ctx)
		loginSession = entClient.LoginSession.Create().SetOwner(user).SetLastLoginTime(lib.TimeNow()).SaveX(ctx)

		mockCookieReader = libmocks.CookieReader{}
		mockHttpErrorWriter = libmocks.HttpErrorWriter{}
		mockHttpErrorWriter.On("WriteUnauthorized")

		originalNewEntClient = lib.NewEntClient
		lib.NewEntClient = func() (*entgenerated.Client, intl.IntlError) {
			return entClient, nil
		}

		mockLogger = libmocks.Logger{}
		mockLogger.On("LogError", mock.Anything)
		originalNewLogger = lib.NewLogger
		lib.NewLogger = func(_name string) lib.Logger {
			return &mockLogger
		}
	})

	AfterEach(func() {
		config.GetAwsConfig = originalGetAwsConfig
		lib.TimeNow = originalTimeNow
		lib.NewEntClient = originalNewEntClient
		lib.NewLogger = originalNewLogger
		entClient.Close()
	})

	Context("queries depending on refresh token", func() {
		type RefreshUserTokenQuery struct {
			RefreshUserToken *gqlgenerated.UserToken `json:"refreshUserToken"`
		}

		DoPostRefreshUserToken := func(option client.Option) (*RefreshUserTokenQuery, error) {
			var response RefreshUserTokenQuery
			err := graphqlClient.Post(
				`
				{
					refreshUserToken {
						authToken
						refreshToken
					}
				}
			`,
				&response,
				option,
			)
			if err != nil {
				return nil, err
			}

			return &response, nil
		}

		type LogoutQuery struct {
			Logout bool `json:"logout"`
		}

		DoPostLogout := func(option client.Option) (*LogoutQuery, error) {
			var response LogoutQuery
			err := graphqlClient.Post(
				`
				mutation {
					logout
				}
			`,
				&response,
				option,
			)
			if err != nil {
				return nil, err
			}

			return &response, nil
		}
		assertDependencyError := func(response interface{}, err error, path string) {
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(fmt.Sprintf("[{\"message\":\"common-strings.unknown-server-error\",\"path\":[\"%s\"]}]", path)))
			Expect(response).To(BeNil())
		}

		Context("missing cookie reader from context", func() {
			Describe("refreshUserToken query", func() {
				It("returns unknown error", func() {
					response, err := DoPostRefreshUserToken(
						requestWithContext(
							func(parent context.Context) context.Context {
								return parent
							},
						),
					)
					assertDependencyError(response, err, "refreshUserToken")
				})
			})

			Describe("logout query", func() {
				It("returns unknown error", func() {
					response, err := DoPostLogout(
						requestWithContext(
							func(parent context.Context) context.Context {
								return parent
							},
						),
					)
					assertDependencyError(response, err, "logout")
				})
			})
		})

		Context("missing http error writer from context", func() {
			Describe("refreshUserToken query", func() {
				It("returns unknown error", func() {
					response, err := DoPostRefreshUserToken(
						requestWithContext(
							func(parent context.Context) context.Context {
								return lib.NewContextWithCookieReader(parent, &mockCookieReader)
							},
						),
					)
					assertDependencyError(response, err, "refreshUserToken")
				})
			})

			Describe("logout", func() {
				It("returns unknown error", func() {
					response, err := DoPostLogout(
						requestWithContext(
							func(parent context.Context) context.Context {
								return lib.NewContextWithCookieReader(parent, &mockCookieReader)
							},
						),
					)
					assertDependencyError(response, err, "logout")
				})
			})
		})

		queryOptionsWithoutViewer := requestWithContext(
			func(parent context.Context) context.Context {
				result := lib.NewContextWithCookieReader(parent, &mockCookieReader)
				result = lib.NewContextWithHttpErrorWriter(result, &mockHttpErrorWriter)
				return result
			},
		)
		queryOptionsWithViewer := requestWithContext(
			func(parent context.Context) context.Context {
				result := lib.NewContextWithCookieReader(parent, &mockCookieReader)
				result = lib.NewContextWithHttpErrorWriter(result, &mockHttpErrorWriter)
				result = viewer.NewContext(result, entviewer.NewUserViewerFromUser(user))
				return result
			},
		)

		Context("with refresh token cookie", func() {
			Context("cookie is valid", func() {
				var userTokenFactory lib.UserTokenFactory
				var refreshToken string
				BeforeEach(func() {
					userTokenFactory = lib.NewUserTokenFactory()
					var err error
					refreshToken, err = userTokenFactory.CreateRefreshToken(loginSession)
					if err != nil {
						panic(err)
					}

					mockCookieReader.On("ReadRefreshToken").Return(
						cookies.NewRefreshTokenCookie(
							refreshToken,
							currentTime.Add(time.Hour),
							http.SameSiteStrictMode,
						),
						nil,
					)
				})

				Context("login session ID is valid", func() {
					Describe("refreshUserToken query", func() {
						It("returns a valid auth token", func() {
							response, err := DoPostRefreshUserToken(queryOptionsWithoutViewer)
							Expect(err).To(BeNil())
							Expect(response.RefreshUserToken.AuthToken).To(BeAssignableToTypeOf(""))
							expectedAuthClaims := libmocks.AuthTokenClaims{}
							expectedAuthClaims.On("GetUserId").Return(user.ID, true)
							expectedAuthClaims.On("GetExpiresAt").Return(currentTime.Add(time.Minute*5), true)
							expectedSignature, err := userTokenFactory.GetAuthTokenSignature(&expectedAuthClaims)
							Expect(err).To(BeNil())
							Expect(response.RefreshUserToken.AuthToken).To(matchers.BeValidJwtToken(expectedSignature))

							Expect(response.RefreshUserToken.RefreshToken).To(BeAssignableToTypeOf(""))
						})
					})

					Context("with owner as viewer", func() {
						Describe("logout query", func() {
							It("returns true and deletes the session", func() {
								response, err := DoPostLogout(queryOptionsWithViewer)
								Expect(err).To(BeNil())
								Expect(response.Logout).To(BeTrue())
								_, err = entClient.LoginSession.Get(superCtx, loginSession.ID)
								Expect(entgenerated.IsNotFound(err)).To(BeTrue())
							})
						})
					})
				})

				Context("login session ID is invalid", func() {
					BeforeEach(func() {
						ctx := viewer.NewContext(context.Background(), &testViewer{})
						entClient.LoginSession.DeleteOne(loginSession).ExecX(ctx)
					})

					Describe("refreshAuthToken query", func() {
						It("returns unauthorized error", func() {
							response, err := DoPostRefreshUserToken(queryOptionsWithoutViewer)
							Expect(err.Error()).To(Equal("[{\"message\":\"common-strings.unauthorized\",\"path\":[\"refreshUserToken\"]}]"))
							Expect(response).To(BeNil())
						})
					})

					Describe("logout query", func() {
						It("returns unauthorized error", func() {
							response, err := DoPostLogout(queryOptionsWithViewer)
							Expect(err.Error()).To(Equal("[{\"message\":\"common-strings.unauthorized\",\"path\":[\"logout\"]}]"))
							Expect(response).To(BeNil())
						})
					})
				})

				Context("owner is not found", func() {
					BeforeEach(func() {
						ctx := viewer.NewContext(context.Background(), &testViewer{})
						entClient.User.DeleteOne(user).ExecX(ctx)
					})

					Describe("refreshUserToken query", func() {
						It("returns unauthorized error", func() {
							response, err := DoPostRefreshUserToken(queryOptionsWithoutViewer)
							Expect(err.Error()).To(Equal("[{\"message\":\"common-strings.unauthorized\",\"path\":[\"refreshUserToken\"]}]"))
							Expect(response).To(BeNil())
						})
					})

					Describe("logout query", func() {
						It("returns unauthorized error", func() {
							response, err := DoPostLogout(queryOptionsWithViewer)
							Expect(err.Error()).To(Equal("[{\"message\":\"common-strings.unauthorized\",\"path\":[\"logout\"]}]"))
							Expect(response).To(BeNil())
						})
					})
				})

				Context("error creating auth token", func() {
					var originalNewUserTokenFactory lib.NewUserTokenFactoryFunc
					var mockedUserTokenFactory *libmocks.UserTokenFactory
					BeforeEach(func() {
						originalNewUserTokenFactory = lib.NewUserTokenFactory
						lib.NewUserTokenFactory = func() lib.UserTokenFactory {
							result := libmocks.UserTokenFactory{}
							result.On("ParseRefreshToken", refreshToken).Return(userTokenFactory.ParseRefreshToken(refreshToken))
							result.On("CreateAuthToken", mock.Anything).Return("", &intl.InvalidJwtTokenError{})
							mockedUserTokenFactory = &result
							return &result
						}
					})

					AfterEach(func() {
						lib.NewUserTokenFactory = originalNewUserTokenFactory
					})

					Describe("refreshUserToken query", func() {
						It("returns unauthorized error", func() {
							response, err := DoPostRefreshUserToken(
								requestWithContext(
									func(parent context.Context) context.Context {
										result := lib.NewContextWithCookieReader(parent, &mockCookieReader)
										result = lib.NewContextWithHttpErrorWriter(result, &mockHttpErrorWriter)
										return result
									},
								),
							)
							Expect(err.Error()).To(Equal("[{\"message\":\"common-strings.unauthorized\",\"path\":[\"refreshUserToken\"]}]"))
							Expect(response).To(BeNil())
							mockedUserTokenFactory.AssertCalled(TheT, "CreateAuthToken", mock.Anything)
							arg, ok := mockedUserTokenFactory.Calls[1].Arguments.Get(0).(*entgenerated.User)
							Expect(ok).To(BeTrue())
							Expect(arg.ID).To(Equal(user.ID))
						})
					})
				})
			})

			Context("cookie is invalid", func() {
				BeforeEach(func() {
					mockCookieReader.
						On("ReadRefreshToken").
						Return(
							cookies.NewRefreshTokenCookie(
								"invalid cookie value",
								currentTime.Add(time.Hour),
								http.SameSiteStrictMode,
							),
							nil,
						)
				})

				Describe("refreshUserToken query", func() {
					It("throws unauthorized error", func() {
						response, err := DoPostRefreshUserToken(queryOptionsWithoutViewer)
						Expect(err.Error()).To(Equal("[{\"message\":\"common-strings.unauthorized\",\"path\":[\"refreshUserToken\"]}]"))
						Expect(response).To(BeNil())
					})
				})

				Describe("logout query", func() {
					It("throws unauthorized error", func() {
						response, err := DoPostLogout(queryOptionsWithViewer)
						Expect(err.Error()).To(Equal("[{\"message\":\"common-strings.unauthorized\",\"path\":[\"logout\"]}]"))
						Expect(response).To(BeNil())
					})
				})
			})
		})

		Context("without refresh token cookie", func() {
			BeforeEach(func() {
				mockCookieReader.
					On("ReadRefreshToken").
					Return(nil, http.ErrNoCookie)
			})

			Describe("refreshUserToken query", func() {
				It("returns nil response", func() {
					response, err := DoPostRefreshUserToken(queryOptionsWithoutViewer)
					Expect(err).To(BeNil())
					Expect(response.RefreshUserToken).To(BeNil())
				})
			})
		})
	})

	Describe("viewer", func() {
		var queryContext context.Context
		type User struct {
			Id string `json:"id"`
		}
		type ViewerQueryResult struct {
			Viewer *User `json:"viewer"`
		}

		doViewerQuery := func(ctx context.Context) (*ViewerQueryResult, error) {
			var response *ViewerQueryResult
			err := graphqlClient.Post(
				`
				{
					viewer {
						id
					}
				}
			`,
				&response,
				requestWithContext(func(parent context.Context) context.Context {
					return queryContext
				}),
			)
			return response, err
		}

		Context("with viewer with correct user ID in the context", func() {
			BeforeEach(func() {
				queryContext = viewer.NewContext(context.Background(), entviewer.NewUserViewerFromUser(user))
			})

			It("returns the correct viewer", func() {
				result, err := doViewerQuery(queryContext)
				Expect(err).To(BeNil())
				Expect(result.Viewer.Id).To(Equal(fmt.Sprint(user.ID)))
			})
		})

		Context("with viewer with invalid user ID in the context", func() {
			BeforeEach(func() {
				queryContext = viewer.NewContext(context.Background(), entviewer.NewUserViewerFromId(1, true))
			})

			It("returns nil viewer", func() {
				result, err := doViewerQuery(queryContext)
				Expect(err).To(BeNil())
				Expect(result.Viewer).To(BeNil())
			})
		})

		Context("without viewer in the context", func() {
			BeforeEach(func() {
				queryContext = context.Background()
			})

			It("returns nil viewer", func() {
				result, err := doViewerQuery(queryContext)
				Expect(err).To(BeNil())
				Expect(result.Viewer).To(BeNil())
			})
		})
	})

	Describe("createProfilePhotoUploadUrl mutation", func() {
		var owner *entgenerated.User
		var ownerProfile *entgenerated.UserPublicProfile
		var queryViewer viewer.Viewer

		type MutationResponse struct {
			CreateProfilePhotoUploadUrl *string `json:"createProfilePhotoUploadUrl"`
		}

		doMutation := func() (*MutationResponse, error) {
			response := MutationResponse{}
			err := graphqlClient.Post(
				`
				mutation CreateProfilePhotoUploadUrlMutation($userId: ID!, $md5: String!) {
					createProfilePhotoUploadUrl(userId: $userId, md5: $md5)
				}
				`,
				&response,
				client.Var("userId", owner.ID),
				client.Var("md5", "test-md5"),
				requestWithContext(func(parent context.Context) context.Context {
					return viewer.NewContext(parent, queryViewer)
				}),
			)
			return &response, err
		}

		BeforeEach(func() {
			owner = user
			ownerProfile = entClient.UserPublicProfile.
				Create().
				SetOwner(owner).
				SetHandleName("test-handle").
				SaveX(superCtx)
		})

		Context("mutation by owner", func() {
			BeforeEach(func() {
				queryViewer = entviewer.NewUserViewerFromUser(owner)
			})

			Context("without existing profile photo", func() {
				It("creates a new blob key and generate the url", func() {
					Expect(ownerProfile.PhotoBlobKey).To(Equal(""))
					actual, err := doMutation()
					Expect(err).To(BeNil())
					Expect(*actual.CreateProfilePhotoUploadUrl).NotTo(BeEmpty())
				})
			})

			Context("with existing profile photo", func() {
				var existingKey string

				getUpdatedPhotoBlobKey := func() string {
					ownerProfile = entClient.UserPublicProfile.GetX(superCtx, ownerProfile.ID)
					return ownerProfile.PhotoBlobKey
				}

				BeforeEach(func() {
					doMutation()
					existingKey = getUpdatedPhotoBlobKey()
				})

				It("does not create a new blob key", func() {
					actual, err := doMutation()
					Expect(err).To(BeNil())
					Expect(*actual.CreateProfilePhotoUploadUrl).NotTo(BeEmpty())
					newKey := getUpdatedPhotoBlobKey()
					Expect(newKey).To(Equal(existingKey))
				})
			})
		})

		Context("mutation by non-owner", func() {
			BeforeEach(func() {
				newUser := createTestUser(entClient)
				queryViewer = entviewer.NewUserViewerFromUser(newUser)
			})

			It("is not allowed", func() {
				actual, err := doMutation()
				Expect(mockLogger.AssertCalled(
					TheT,
					"LogError",
					mock.MatchedBy(func(err error) bool {
						return entgenerated.IsNotFound(err)
					}),
				)).To(BeTrue())
				Expect(err.Error()).To(Equal("[{\"message\":\"blob-storage.unable-to-create-upload-url\",\"path\":[\"createProfilePhotoUploadUrl\"]}]"))
				Expect(actual.CreateProfilePhotoUploadUrl).To(BeNil())
			})
		})
	})
})
