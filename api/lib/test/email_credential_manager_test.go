package test

import (
	"api/config"
	"api/config/configmocks"
	"api/ent/entgenerated"
	"api/ent/entgenerated/loginsession"
	"api/ent/entgenerated/user"
	"api/lib"
	"api/lib/libmocks"
	"api/lib/matchers"
	"api/privacy/viewer"
	"context"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"
)

var _ = Describe("email_credential", func() {
	var subject lib.EmailCredentialManager
	var currentTime time.Time
	var userTokenFactory lib.UserTokenFactory

	BeforeEach(func() {
		hmacConfig := configmocks.HmacConfig{}
		hmacConfig.On("GetAuthTokenSecret").Return("auth token hmac secret")
		hmacConfig.On("GetRefreshTokenSecret").Return("refresh token hmac secret")
		config.GetHmacConfig = func() config.HmacConfig {
			return &hmacConfig
		}

		subject = lib.NewEmailCredentialManager(config.GetAuthConfig())

		currentTime = time.Now()
		lib.TimeNow = func() time.Time {
			return currentTime
		}

		userTokenFactory = lib.NewUserTokenFactory()
	})

	Describe("NewEmailCredentialManager", func() {
		It("creates an EmailCredentialManagerWithConfig instance", func() {
			Expect(subject).To(BeAssignableToTypeOf(&lib.EmailCredentialManagerWithConfig{}))
		})
	})

	Describe("Login", func() {
		var entClient *entgenerated.Client
		var registeredUser *entgenerated.User
		var loginCtx context.Context
		var httpWriter httptest.ResponseRecorder

		BeforeEach(func() {
			entClient = createEntClientForTest()
			registeredUser = createTestUser(entClient)
			pwHash, err := bcrypt.GenerateFromPassword([]byte("test-password"), bcrypt.DefaultCost)
			if err != nil {
				panic(err)
			}

			v := testViewer{}
			ctx := viewer.NewContext(context.Background(), v)
			entClient.EmailCredential.Create().
				SetOwner(registeredUser).
				SetAlgorithm("bcrypt").
				SetEmail("test@example.com").
				SetPasswordHash(pwHash).
				SaveX(ctx)

			httpWriter = httptest.ResponseRecorder{}
			cookieWriter := lib.NewCookieWriter(&httpWriter)
			loginCtx = lib.NewContextWithCookieWriter(context.Background(), cookieWriter)
		})

		getExpectedAuthSignature := func() []byte {
			expectedAuthClaims := libmocks.AuthTokenClaims{}
			expectedAuthClaims.On("GetUserId").Return(registeredUser.ID, true)
			expectedAuthClaims.On("GetExpiresAt").Return(currentTime.Add(time.Minute*5), true)
			authTokenSignature, err := userTokenFactory.GetAuthTokenSignature(&expectedAuthClaims)
			if err != nil {
				panic(err)
			}
			return authTokenSignature
		}

		getExpectedRefreshSignature := func() []byte {
			ctx := viewer.NewContext(context.Background(), &testViewer{})
			lastLoginSession := entClient.LoginSession.Query().Where(
				loginsession.HasOwnerWith(user.ID(registeredUser.ID)),
			).OnlyX(ctx)
			expectedClaims := libmocks.RefreshTokenClaims{}
			expectedClaims.On("GetLoginSessionId").Return(lastLoginSession.ID, true)
			expectedClaims.On("GetIssuedAt").Return(currentTime, true)
			expectedClaims.On("GetExpiresAt").Return(currentTime.Add(lib.REFRESH_TOKEN_DURATION_MONTH*time.Hour*24*30), true)
			refreshSignature, err := userTokenFactory.GetRefreshTokenSignature(&expectedClaims)
			if err != nil {
				panic(err)
			}
			return refreshSignature
		}

		Context("with correct email and password", func() {
			It("returns the correct tokens", func() {
				tokens, err := subject.Login(loginCtx, entClient, "test@example.com", "test-password")
				Expect(err).To(BeNil())
				Expect(tokens.AuthToken).To(matchers.BeValidJwtToken(getExpectedAuthSignature()))
				Expect(tokens.RefreshToken).To(matchers.BeValidJwtToken(getExpectedRefreshSignature()))
			})
		})

		Context("with correct email but wrong password", func() {
			It("returns nil and EmailCredentialNotFoundError", func() {
				tokens, err := subject.Login(loginCtx, entClient, "test@example.com", "wrong-password")
				Expect(err).To(MatchError(&lib.EmailCredentialNotFoundError{}))
				Expect(tokens).To(BeNil())
			})
		})

		Context("with incorrect email", func() {
			It("return nil and EMailCredentialNotFoundError", func() {
				tokens, err := subject.Login(loginCtx, entClient, "test2@example.com", "test-password")
				Expect(err).To(MatchError(&lib.EmailCredentialNotFoundError{}))
				Expect(tokens).To(BeNil())
			})
		})
	})
})
