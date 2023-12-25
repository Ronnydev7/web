package test

import (
	"api/config"
	"api/config/configmocks"
	"api/ent/entgenerated"
	"api/intl"
	"api/lib"
	"api/lib/libmocks"
	"api/lib/matchers"
	"api/privacy/viewer"
	"context"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"
)

var _ = Describe("user_token", func() {
	var user *entgenerated.User
	var loginSession *entgenerated.LoginSession
	var emailCredential *entgenerated.EmailCredential
	var subject lib.UserTokenFactory
	var currentTime time.Time
	var originalTimeFunc func() time.Time
	var ctx context.Context
	var client *entgenerated.Client

	BeforeEach(func() {
		config.GetHmacConfig = func() config.HmacConfig {
			result := configmocks.HmacConfig{}
			result.On("GetAuthTokenSecret").Return("auth_token_secret")
			result.On("GetRefreshTokenSecret").Return("refresh_token_secret")
			return &result
		}

		subject = lib.NewUserTokenFactory()
		currentTime = time.Now()
		lib.TimeNow = func() time.Time {
			return currentTime
		}

		client = createEntClientForTest()
		ctx = viewer.NewContext(context.Background(), createTestAdminViewer())
		// Create User
		user = client.User.Create().SaveX(ctx)

		// Create Email Credential
		passwordHash, err := bcrypt.GenerateFromPassword([]byte("testpassword"), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		emailCredential = client.EmailCredential.
			Create().
			SetOwner(user).
			SetAlgorithm("bcrypt").
			SetEmail("test@example.com").
			SetPasswordHash(passwordHash).
			SaveX(ctx)

		// Create Login Session
		loginSession = client.LoginSession.Create().
			SetOwner(user).
			SetLastLoginTime(lib.TimeNow()).
			SaveX(ctx)

		originalTimeFunc = jwt.TimeFunc
		jwt.TimeFunc = func() time.Time {
			return currentTime
		}
	})

	AfterEach(func() {
		jwt.TimeFunc = originalTimeFunc
		client.Close()
	})

	Describe("CreateAuthToken, GetAuthTokenSignatureSecret, ParseAuthToken", func() {
		sharedTestStep := func() string {
			token, err := subject.CreateAuthToken(user)
			Expect(err).To(BeNil())

			expectedClaims := libmocks.AuthTokenClaims{}
			expectedClaims.On("GetUserId").Return(user.ID, true)
			expectedClaims.On("GetExpiresAt").Return(
				currentTime.Add(time.Minute*5),
				true,
			)
			expectedSecret, err := subject.GetAuthTokenSignature(&expectedClaims)
			Expect(err).To(BeNil())
			Expect(token).To(matchers.BeValidJwtToken(expectedSecret))
			return token
		}

		It("creates valid jwt token", func() {
			jwtToken := sharedTestStep()
			parsedToken, err := subject.ParseAuthToken(jwtToken)
			Expect(err).To(BeNil())
			userId, exists := parsedToken.GetUserId()
			Expect(exists).To(BeTrue())
			Expect(userId).To(Equal(user.ID))
			expiresAt, exists := parsedToken.GetExpiresAt()
			Expect(exists).To(BeTrue())
			Expect(expiresAt).To(matchers.EqualTimeInSeconds(currentTime.Add(time.Minute * 5)))
		})

		Context("expired token", func() {
			It("unable to parse", func() {
				jwtToken := sharedTestStep()
				jwt.TimeFunc = func() time.Time {
					return currentTime.Add(time.Second * 301)
				}
				parsedToken, err := subject.ParseAuthToken(jwtToken)
				Expect(err).NotTo(BeNil())
				Expect(parsedToken).To(BeNil())
			})
		})

		Context("compromised token", func() {
			It("unable to parse", func() {
				jwtToken := sharedTestStep()
				jwtParts := strings.Split(jwtToken, ".")

				anotherUser := client.User.Create().SaveX(ctx)
				anotherJwtToken, err := subject.CreateAuthToken(anotherUser)
				Expect(err).To(BeNil())
				anotherJwtParts := strings.Split(anotherJwtToken, ".")

				jwtParts[1] = anotherJwtParts[1]
				compromisedToken := strings.Join(jwtParts, ".")
				parsedToken, err := subject.ParseAuthToken(compromisedToken)
				Expect(parsedToken).To(BeNil())
				Expect(err).NotTo(BeNil())
			})
		})
	})

	Describe("CreateRefreshToken, GetRefreshTokenSignatureSecret, ParseRefreshToken", func() {
		sharedTestStep := func() string {
			expectedClaims := libmocks.RefreshTokenClaims{}
			expectedClaims.On("GetLoginSessionId").Return(
				loginSession.ID,
				true,
			)
			expectedClaims.On("GetIssuedAt").Return(
				currentTime,
				true,
			)
			expectedClaims.On("GetExpiresAt").Return(
				currentTime.Add(time.Hour*24*30*6),
				true,
			)
			expectedSecret, err := subject.GetRefreshTokenSignature(&expectedClaims)
			Expect(err).To(BeNil())

			token, err := subject.CreateRefreshToken(loginSession)
			Expect(err).To(BeNil())
			Expect(token).To(matchers.BeValidJwtToken(expectedSecret))
			return token
		}

		It("creates valid jwt token", func() {
			jwtToken := sharedTestStep()
			parsedToken, err := subject.ParseRefreshToken(jwtToken)
			Expect(err).To(BeNil())
			sessionId, exists := parsedToken.GetLoginSessionId()
			Expect(exists).To(BeTrue())
			Expect(sessionId).To(Equal(loginSession.ID))

			issuedAt, exists := parsedToken.GetIssuedAt()
			Expect(exists).To(BeTrue())
			Expect(issuedAt).To(matchers.EqualTimeInSeconds(currentTime))

			expiresAt, exists := parsedToken.GetExpiresAt()
			Expect(exists).To(BeTrue())
			Expect(expiresAt).To(matchers.EqualTimeInSeconds(currentTime.Add(
				time.Hour * 24 * 30 * 6,
			)))
		})

		Context("invalid token", func() {
			It("unable to parse", func() {
				jwtToken := sharedTestStep()
				jwt.TimeFunc = func() time.Time {
					return currentTime.Add(time.Hour*24*30*lib.REFRESH_TOKEN_DURATION_MONTH + time.Second)
				}
				parsedToken, err := subject.ParseRefreshToken(jwtToken)
				Expect(err).NotTo(BeNil(), "expected error not returned")
				Expect(parsedToken).To(BeNil())
			})
		})
	})

	Describe("CreateResetPasswordToken, GetResetPasswordTokenSIgnature, ParseResetPasswordToken", func() {
		sharedTestStep := func() string {
			expectedClaims := libmocks.ResetPasswordTokenClaims{}
			expectedClaims.On("GetEmailCredentialId").Return(
				emailCredential.ID,
				true,
			)
			expectedClaims.On("GetExpiresAt").Return(
				currentTime.Add(time.Minute*30),
				true,
			)

			expectedSecret, err := subject.GetResetPasswordTokenSignature(&expectedClaims, emailCredential)
			Expect(err).To(BeNil())

			token, err := subject.CreateResetPasswordToken(emailCredential)
			Expect(err).To(BeNil())
			Expect(token).To(matchers.BeValidJwtToken(expectedSecret))
			return token
		}

		It("create valid jwt token", func() {
			jwtToken := sharedTestStep()
			parsedToken, err := subject.ParseResetPasswordToken(ctx, jwtToken, client)
			Expect(err).To(BeNil())
			emailCredentialId, exists := parsedToken.GetEmailCredentialId()
			Expect(exists).To(BeTrue())
			Expect(emailCredentialId).To(Equal(emailCredential.ID))

			expiresAt, exists := parsedToken.GetExpiresAt()
			Expect(exists).To(BeTrue())
			Expect(expiresAt).To(matchers.EqualTimeInSeconds(currentTime.Add(
				time.Minute * 30,
			)))
		})

		Context("expired token", func() {
			It("fails parsing", func() {
				jwtToken := sharedTestStep()
				jwt.TimeFunc = func() time.Time {
					return currentTime.Add(time.Minute*30 + time.Second)
				}

				parsedToken, err := subject.ParseResetPasswordToken(ctx, jwtToken, client)
				Expect(err).To(MatchError(&intl.ExpiredJwtTokenError{}))
				Expect(parsedToken).To(BeNil())
			})
		})

		Context("password changed", func() {
			It("fails parsing", func() {
				jwtToken := sharedTestStep()
				client.EmailCredential.UpdateOne(emailCredential).SetPasswordHash([]byte("newpassword")).SaveX(ctx)
				parsedToken, err := subject.ParseResetPasswordToken(ctx, jwtToken, client)
				Expect(err).To(MatchError(&intl.InvalidJwtTokenSignatureError{}))
				Expect(parsedToken).To(BeNil())
			})
		})

		Context("email credential not found", func() {
			It("fails parsing", func() {
				jwtToken := sharedTestStep()
				client.EmailCredential.DeleteOne(emailCredential).Exec(ctx)

				parsedToken, err := subject.ParseResetPasswordToken(ctx, jwtToken, client)
				Expect(err).To(MatchError(&intl.InvalidJwtTokenError{}))
				Expect(parsedToken).To(BeNil())
			})
		})
	})
})
