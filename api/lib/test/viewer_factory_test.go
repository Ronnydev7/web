package test

import (
	"api/config"
	"api/config/configmocks"
	"api/ent/entgenerated"
	"api/ent/privacy/entviewer"
	"api/intl"
	"api/lib"
	"api/lib/libmocks"
	"api/lib/testutils"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("ViewerFactory", func() {
	var request *http.Request
	var owner *entgenerated.User
	var entClient *entgenerated.Client
	var ownedToken string
	var subject lib.ViewerFactory

	BeforeEach(func() {
		config.GetHmacConfig = func() config.HmacConfig {
			result := configmocks.NewHmacConfig(TheT)
			result.On("GetAuthTokenSecret").Return("auth token hmac secret")
			return result
		}

		entClient = createEntClientForTest()
		lib.NewEntClient = func() (*entgenerated.Client, intl.IntlError) {
			return entClient, nil
		}

		userTokenFactory := lib.NewUserTokenFactory()

		owner = createTestUser(entClient)
		var err error
		ownedToken, err = userTokenFactory.CreateAuthToken(owner)
		if err != nil {
			panic(err)
		}

		subject = lib.NewViewerFactory()
	})

	Context("POST request", func() {
		BeforeEach(func() {
			request = httptest.NewRequest("POST", "/query", nil)
		})

		Describe("FromHttpAuthorizationHeader", func() {
			Context("with correct Authorization header and Bearer token", func() {
				BeforeEach(func() {
					headerValue := fmt.Sprintf("Bearer %s", ownedToken)
					request.Header.Add("Authorization", headerValue)
				})

				It("returns the correct user", func() {
					actual, err := subject.FromHttpAuthorizationHeader(request)
					Expect(err).To(BeNil())
					actualId, actualExists := actual.GetId()
					Expect(actualId).To(Equal(owner.ID))
					Expect(actualExists).To(BeTrue())
				})

				Describe("FromHttpRequestOrElseAnonymous", func() {
					It("returns the same instance", func() {
						actual := subject.FromHttpRequestOrElseAnonymous(entClient, request)
						id, hasId := actual.GetId()
						Expect(hasId).To(BeTrue())
						Expect(id).To(Equal(owner.ID))
					})
				})
			})

			Context("with Authorization header and incorrect Bearer token", func() {
				BeforeEach(func() {
					request.Header.Add("Authorization", "Bearer abcdefghijklmnopqrstuvwxyz")
				})

				It("returns nil and InvalidAuthTokenError", func() {
					actual, err := subject.FromHttpAuthorizationHeader(request)
					Expect(err).To(MatchError(&lib.InvalidAuthTokenError{}))
					Expect(actual).To(BeNil())
				})
			})

			Context("with incorrect Authorization header", func() {
				BeforeEach(func() {
					request.Header.Add("Authorization", "this is invalid authorization header")
				})

				It("returns nil and InvalidBearerAuthorizationHeaderError", func() {
					actual, err := subject.FromHttpAuthorizationHeader(request)
					Expect(err).To(MatchError(&lib.InvalidBearerAuthorizationHeaderError{}))
					Expect(actual).To(BeNil())
				})
			})

			Context("without Authorization header", func() {
				It("returns nil and InvalidBearerAuthorizationHeaderError", func() {
					actual, err := subject.FromHttpAuthorizationHeader(request)
					Expect(err).To(MatchError(&lib.InvalidBearerAuthorizationHeaderError{}))
					Expect(actual).To(BeNil())
				})
			})
		})

		Describe("FromHttpRequestOrElseAnonymous", func() {
			Context("with correct Authorization header and Bearer token", func() {
				BeforeEach(func() {
					headerValue := fmt.Sprintf("Bearer %s", ownedToken)
					request.Header.Add("Authorization", headerValue)
				})

				It("returns the viewer from the auth token", func() {
					actual := subject.FromHttpRequestOrElseAnonymous(entClient, request)
					id, hasId := actual.GetId()
					Expect(hasId).To(BeTrue())
					Expect(id).To(Equal(owner.ID))
				})
			})

			Context("without Authorization header", func() {
				var mockUserFactory *libmocks.UserFactory
				var mockNewUserFactory testutils.MockWrapper[lib.NewUserFactoryFunc]

				var mockGetHmacConfig testutils.MockWrapper[config.GetHmacConfigFunc]
				BeforeEach(func() {
					mockUserFactory = libmocks.NewUserFactory(TheT)
					mockNewUserFactory = testutils.StartMockWrapper[lib.NewUserFactoryFunc](
						func(_ *entgenerated.Client) lib.UserFactory {
							return mockUserFactory
						},
						lib.NewUserFactory,
						func(aNew lib.NewUserFactoryFunc) {
							lib.NewUserFactory = aNew
						},
					)
					mockGetHmacConfig = testutils.StartMockWrapper[config.GetHmacConfigFunc](
						func() config.HmacConfig {
							result := configmocks.NewHmacConfig(TheT)
							result.On("GetAuthTokenSecret").Return("auth token secret")
							return result
						},
						config.GetHmacConfig,
						func(aGet config.GetHmacConfigFunc) {
							config.GetHmacConfig = aGet
						},
					)
				})
				AfterEach(func() {
					mockNewUserFactory.Stop()
					mockGetHmacConfig.Stop()
				})

				Context("when user factory returns a valid user", func() {
					BeforeEach(func() {
						mockUserFactory.
							On("FromRefreshTokenCookieBypassPrivacy", mock.Anything, mock.Anything, mock.Anything).
							Return(owner, nil)
					})

					It("returns the user's viewer", func() {
						actual := subject.FromHttpRequestOrElseAnonymous(entClient, request)
						id, hasId := actual.GetId()
						Expect(hasId).To(BeTrue())
						Expect(id).To(Equal(owner.ID))
					})

					Context("when fails to create auth token", func() {
						var mockNewUserTokenFactory testutils.MockWrapper[lib.NewUserTokenFactoryFunc]

						BeforeEach(func() {
							mockNewUserTokenFactory = testutils.StartMockWrapper[lib.NewUserTokenFactoryFunc](
								func() lib.UserTokenFactory {
									result := libmocks.NewUserTokenFactory(TheT)
									result.On("CreateAuthToken", mock.Anything).Return("", &intl.InvalidJwtTokenError{})
									return result
								},
								lib.NewUserTokenFactory,
								func(aNew lib.NewUserTokenFactoryFunc) {
									lib.NewUserTokenFactory = aNew
								},
							)
						})

						AfterEach(func() {
							mockNewUserTokenFactory.Stop()
						})

						It("returns the anonymous viewer", func() {
							actual := subject.FromHttpRequestOrElseAnonymous(entClient, request)
							id, hasId := actual.GetId()
							Expect(hasId).To(BeFalse())
							Expect(id).To(Equal(0))
							Expect(actual).To(BeAssignableToTypeOf(entviewer.NewAnonymouseUserViewer()))
						})
					})
				})

				Context("when user factory returns nil", func() {
					BeforeEach(func() {
						mockUserFactory.
							On("FromRefreshTokenCookieBypassPrivacy", mock.Anything, mock.Anything, mock.Anything).
							Return(nil, errors.New("not found"))
					})

					It("returns the anonymous viewer", func() {
						actual := subject.FromHttpRequestOrElseAnonymous(entClient, request)
						id, hasId := actual.GetId()
						Expect(hasId).To(BeFalse())
						Expect(id).To(Equal(0))
						Expect(actual).To(BeAssignableToTypeOf(entviewer.NewAnonymouseUserViewer()))
					})
				})
			})
		})
	})
})
