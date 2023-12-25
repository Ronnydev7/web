package test

import (
	"api/config"
	"api/config/configmocks"
	"api/ent/entgenerated"
	"api/graphql/gqlgenerated"
	"api/intl"
	"api/lib"
	"api/lib/cookies"
	"api/lib/libmocks"
	"api/lib/matchers"
	"api/lib/testutils"
	"api/privacy/viewer"
	"context"
	"net/http/httptest"
	"time"

	"entgo.io/ent/entql"
	"github.com/99designs/gqlgen/client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

var _ = Describe("email_credential.resolvers", func() {
	var mockNewEntClient testutils.MockWrapper[lib.NewEntClientFunc]
	var mockNewMailer testutils.MockWrapper[lib.NewMailerFunc]
	var entClient *entgenerated.Client
	var graphqlClient *client.Client
	var currentTime time.Time
	var httpResponse *httptest.ResponseRecorder
	var cookieWriter lib.CookieWriter
	var mockMailer libmocks.Mailer

	BeforeEach(func() {
		entClient = createEntClientForTest()
		mockNewEntClient = testutils.StartMockWrapper[lib.NewEntClientFunc](
			func() (*entgenerated.Client, intl.IntlError) {
				return entClient, nil
			},
			lib.NewEntClient,
			func(newEntClient lib.NewEntClientFunc) {
				lib.NewEntClient = newEntClient
			},
		)

		graphqlClient = createGraphqlClient(entClient)

		currentTime = time.Now()
		lib.TimeNow = func() time.Time {
			return currentTime
		}

		mockMailer = libmocks.Mailer{}
		mockNewMailer = testutils.StartMockWrapper[lib.NewMailerFunc](
			func(_ config.MailerConfig) lib.Mailer {
				return &mockMailer
			},
			lib.NewMailer,
			func(newMailer lib.NewMailerFunc) {
				lib.NewMailer = newMailer
			},
		)

		httpResponse = &httptest.ResponseRecorder{}
		cookieWriter = lib.NewCookieWriter(httpResponse)

		mockUrlConfig := configmocks.UrlConfig{}
		mockUrlConfig.On("GetProtocol").Return("https")
		mockUrlConfig.On("GetHostname").Return("test.hobbytrace.com")
		config.GetUrlConfig = func() config.UrlConfig {
			return &mockUrlConfig
		}

		mockHmacConfig := configmocks.HmacConfig{}
		mockHmacConfig.On("GetEmailSignupTokenSecret").Return("email_signup_token_secret")
		mockHmacConfig.On("GetRefreshTokenSecret").Return("refresh_token_secret")
		mockHmacConfig.On("GetAuthTokenSecret").Return("auth_token_secret")
		config.GetHmacConfig = func() config.HmacConfig {
			return &mockHmacConfig
		}
	})

	AfterEach(func() {
		mockNewEntClient.Stop()
		mockNewMailer.Stop()
		entClient.Close()
	})

	Describe("UserEmailSignup", func() {
		var mockTokenFactory testutils.MockWrapper[lib.NewEmailSignupJwtTokenFactoryFunc]
		BeforeEach(func() {
			mockMailer.On("SendConfirmSignupEmailEmail", "test@example.com", mock.AnythingOfType("string")).Return(nil)
			mockTokenFactory = testutils.StartMockWrapper[lib.NewEmailSignupJwtTokenFactoryFunc](
				func(_ config.HmacConfig) lib.EmailSignupJwtTokenFactory {
					result := libmocks.EmailSignupJwtTokenFactory{}
					result.On("Create", "test@example.com").Return("mock_token", nil)
					return &result
				},
				lib.NewEmailSignupJwtTokenFactory,
				func(factory lib.NewEmailSignupJwtTokenFactoryFunc) {
					lib.NewEmailSignupJwtTokenFactory = factory
				},
			)
		})

		AfterEach(func() {
			mockTokenFactory.Stop()
		})

		It("calls mailer to send mail and return true", func() {
			var response struct {
				UserEmailSignup bool `json:"userEmailSignup"`
			}
			graphqlClient.MustPost(`
				mutation {
					userEmailSignup(email: "test@example.com")
				}
			`, &response)
			Expect(response.UserEmailSignup).To(BeTrue())
			Expect(mockMailer.AssertCalled(TheT, "SendConfirmSignupEmailEmail", "test@example.com", mock.AnythingOfType("string"))).To(BeTrue())
		})
	})

	Describe("ConfirmUserEmailSignup", func() {
		var emailSignupJwtToken string
		var err error
		BeforeEach(func() {
			tokenBuilder := lib.NewEmailSignupJwtTokenFactory(config.GetHmacConfig())
			emailSignupJwtToken, err = tokenBuilder.Create("test@example.com")
			if err != nil {
				panic(err)
			}
		})

		type MutationResponse struct {
			ConfirmUserEmailSignup *gqlgenerated.UserToken `json:"confirmUserEmailSignup"`
		}

		doMutation := func(token string, handleName string, password string) (MutationResponse, error) {
			var response MutationResponse
			err := graphqlClient.Post(
				`
					mutation($token: String!, $handleName: String!, $password: String!) {
						confirmUserEmailSignup(
							emailSignupToken: $token,
							handleName: $handleName,
							rawPassword: $password,
						) {
							authToken
						}
					}
				`,
				&response,
				client.Var("token", token),
				client.Var("handleName", handleName),
				client.Var("password", password),
				requestWithContext(
					func(parent context.Context) context.Context {
						return lib.NewContextWithCookieWriter(parent, cookieWriter)
					},
				),
			)
			return response, err
		}

		It("confirms the email and creates a new user", func() {
			response, err := doMutation(emailSignupJwtToken, "test-handle", "this is a very complex password")
			Expect(err).To(BeNil())
			Expect(httpResponse).To(matchers.HaveSetCookieHeaderByName(cookies.REFRESH_TOKEN_COOKIE_NAME))

			userTokenFactory := lib.NewUserTokenFactory()

			ctx := viewer.NewContext(context.Background(), testViewer{})
			emailCredentialQuery := entClient.EmailCredential.Query()
			emailCredentialQuery.Filter().WhereEmail(entql.StringEQ("test@example.com"))
			user := emailCredentialQuery.QueryOwner().OnlyX(ctx)

			expectedAuthClaims := libmocks.AuthTokenClaims{}
			expectedAuthClaims.On("GetUserId").Return(user.ID, true)
			expectedAuthClaims.On("GetExpiresAt").Return(currentTime.Add(time.Minute*5), true)
			authTokenSignature, err := userTokenFactory.GetAuthTokenSignature(&expectedAuthClaims)
			Expect(err).To(BeNil())
			Expect(response.ConfirmUserEmailSignup.AuthToken).To(matchers.BeValidJwtToken(authTokenSignature))
		})

		It("rejects brainless password", func() {
			_, err := doMutation(emailSignupJwtToken, "test-handle", "abc")
			Expect(err).NotTo(BeNil())
		})

		It("same token cannot be used once registered", func() {
			_, err := doMutation(emailSignupJwtToken, "test-handle", "this is a very complex password")
			Expect(err).To(BeNil())
			Expect(httpResponse).To(matchers.HaveSetCookieHeaderByName(cookies.REFRESH_TOKEN_COOKIE_NAME))

			_, err = doMutation(emailSignupJwtToken, "test-handle2", "this is a very complex password")
			Expect(err).NotTo(BeNil())
		})

		It("rejects duplicate handle name", func() {
			existingUser := createTestUser(entClient)
			entClient.UserPublicProfile.Create().
				SetHandleName("existing-handle").
				SetOwner(existingUser).
				SaveX(viewer.NewContext(context.Background(), &testViewer{}))

			_, err := doMutation(emailSignupJwtToken, "existing-handle", "this is a very complex password")
			Expect(err.Error()).To(Equal("[{\"message\":\"public-profile.duplicate-handle-name\",\"path\":[\"confirmUserEmailSignup\"]}]"))
		})

		It("same token can be used twice if first transaction failed", func() {
			existingUser := createTestUser(entClient)
			entClient.UserPublicProfile.Create().
				SetHandleName("existing-handle").
				SetOwner(existingUser).
				SaveX(viewer.NewContext(context.Background(), &testViewer{}))

			_, err := doMutation(emailSignupJwtToken, "existing-handle", "this is a very complex password")
			Expect(err).NotTo(BeNil())

			_, err = doMutation(emailSignupJwtToken, "new-handle", "this is a very complex password")
			Expect(err).To(BeNil())
			Expect(httpResponse).To(matchers.HaveSetCookieHeaderByName(cookies.REFRESH_TOKEN_COOKIE_NAME))
		})
	})

	Describe("EmailLogin", func() {
		var email string
		var password string
		var ctx context.Context
		BeforeEach(func() {
			email = "testemaillogin@example.com"
			password = "testpassword"
			ctx = viewer.NewContext(context.Background(), testViewer{})

			user := entClient.User.Create().SaveX(ctx)
			passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				panic(err)
			}
			entClient.EmailCredential.Create().
				SetEmail(email).
				SetAlgorithm("bcrypt").
				SetPasswordHash(passwordHash).
				SetOwner(user).
				SaveX(ctx)
		})

		doLogin := func(emailInput string, passwordInput string) (*gqlgenerated.UserToken, error) {
			var response struct {
				EmailLogin gqlgenerated.UserToken `json:"emailLogin"`
			}
			err := graphqlClient.Post(
				`
					mutation($email: String!, $rawPassword: String!) {
						emailLogin(
							email: $email,
							rawPassword: $rawPassword,
						) {
							authToken
						}
					}
				`,
				&response,
				client.Var("email", emailInput),
				client.Var("rawPassword", passwordInput),
				requestWithContext(
					func(parent context.Context) context.Context {
						return lib.NewContextWithCookieWriter(parent, cookieWriter)
					},
				),
			)
			if err != nil {
				return nil, err
			}
			return &response.EmailLogin, nil
		}

		Context("when credential is correct", func() {
			It("returns user token", func() {
				result, err := doLogin(email, password)
				Expect(err).To(BeNil())
				Expect(httpResponse).To(matchers.HaveSetCookieHeaderByName(cookies.REFRESH_TOKEN_COOKIE_NAME))

				userTokenFactory := lib.NewUserTokenFactory()
				emailCredentialQuery := entClient.EmailCredential.Query()
				emailCredentialQuery.Filter().WhereEmail(entql.StringEQ(email))
				user := emailCredentialQuery.QueryOwner().OnlyX(ctx)
				expectedAuthClaims := libmocks.AuthTokenClaims{}
				expectedAuthClaims.On("GetUserId").Return(user.ID, true)
				expectedAuthClaims.On("GetExpiresAt").Return(currentTime.Add(time.Minute*5), true)
				authTokenSecret, err := userTokenFactory.GetAuthTokenSignature(&expectedAuthClaims)
				Expect(err).To(BeNil())
				Expect(result.AuthToken).To(matchers.BeValidJwtToken(authTokenSecret))
			})
		})

		Context("when credential is incorrect", func() {
			It("returns nil", func() {
				result, err := doLogin(email, password+"a")
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("[{\"message\":\"email-credential.invalid-email-credential\",\"path\":[\"emailLogin\"]}]"))
				Expect(result).To(BeNil())
			})
		})

		Context("when email is not found", func() {
			It("returns nil", func() {
				result, err := doLogin("not-found@some-email.com", "abcdefg")
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("[{\"message\":\"email-credential.invalid-email-credential\",\"path\":[\"emailLogin\"]}]"))
				Expect(result).To(BeNil())
			})
		})
	})

	Describe("ResetPassword", func() {
		BeforeEach(func() {
			mockMailer.On("SendResetPasswordEmail", "test@example.com", mock.AnythingOfType("string")).Return(nil)

			user := createTestUser(entClient)
			CreateTestEmailCredentialForUser(entClient, user, "test@example.com")
		})

		doPost := func(email string) (*bool, error) {
			var response struct {
				ResetPassword *bool `json:"resetPassword"`
			}

			err := graphqlClient.Post(
				`
					mutation($email: String!) {
						resetPassword(
							email: $email,
						)
					}
				`,
				&response,
				client.Var("email", email),
			)
			return response.ResetPassword, err
		}

		Context("when email exists", func() {
			It("sends reset password email to the registered address", func() {
				actual, err := doPost("test@example.com")
				Expect(err).To(BeNil())
				Expect(*actual).To(BeTrue())
				Expect(mockMailer.AssertCalled(TheT, "SendResetPasswordEmail", "test@example.com", mock.AnythingOfType("string"))).To(BeTrue())
			})
		})

		Context("when email doesn't exist", func() {
			It("return true but do not send email", func() {
				actual, err := doPost("testresetpassword@example.com")
				Expect(err).To(BeNil())
				Expect(*actual).To(BeTrue())
				Expect(mockMailer.AssertNotCalled(TheT, "SendResetPasswordEmail", "test@example.com", mock.AnythingOfType("string"))).To(BeTrue())
			})
		})
	})
})
