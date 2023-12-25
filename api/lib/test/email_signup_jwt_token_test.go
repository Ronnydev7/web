package test

import (
	"api/config/configmocks"
	"api/intl"
	"api/lib"
	"api/lib/libmocks"
	"api/lib/matchers"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("email_signup_jwt_token", func() {
	var c *configmocks.HmacConfig

	BeforeEach(func() {
		c = &configmocks.HmacConfig{}
		c.On("GetEmailSignupTokenSecret").Return("email signup token hmac secret")
	})

	Describe("NewEmailSignupJwtTokenFactory", func() {
		It("creates a new instance", func() {
			actual := lib.NewEmailSignupJwtTokenFactory(c)
			Expect(actual).To(BeAssignableToTypeOf(&lib.EmailSignupJwtTokenFactoryWithConfig{}))
		})
	})

	Context("with default factory", func() {
		var factory lib.EmailSignupJwtTokenFactory
		var currentTime time.Time

		var originalTimeNow lib.TimeNowFunc

		BeforeEach(func() {
			factory = lib.NewEmailSignupJwtTokenFactory(c)
			currentTime = time.Now()
			originalTimeNow = lib.TimeNow
			lib.TimeNow = func() time.Time {
				return currentTime
			}
		})

		AfterEach(func() {
			lib.TimeNow = originalTimeNow
		})

		Describe("Create", func() {
			It("creates a token with the expected claims", func() {
				email := "test@example.com"
				tokenString, err := factory.Create(email)
				Expect(err).To(BeNil())

				expectedClaims := libmocks.EmailSignupTokenClaims{}
				expectedClaims.On("GetEmail").Return(email, true)
				expectedClaims.On("GetExpiresAt").Return(currentTime.Add(time.Hour), true)

				expectedSignature, err := factory.GetSignatureSecret(&expectedClaims)
				Expect(err).To(BeNil())
				Expect(tokenString).To(matchers.BeValidJwtToken(expectedSignature))
			})
		})

		Describe("Parse", func() {
			var tokenString string
			const email = "test@example.com"

			BeforeEach(func() {
				var err error
				tokenString, err = factory.Create(email)
				if err != nil {
					panic(err)
				}
			})
			It("parse a valid token to claims", func() {
				claims, err := factory.Parse(tokenString)
				Expect(err).To(BeNil())

				actualEmail, exists := claims.GetEmail()
				Expect(exists).To(BeTrue())
				Expect(actualEmail).To(Equal(email))

				actualExpiresAt, exists := claims.GetExpiresAt()
				Expect(exists).To(BeTrue())
				Expect(actualExpiresAt).To(matchers.EqualTimeInSeconds(currentTime.Add(time.Hour)))
			})

			Context("expired token", func() {
				var originalJwtTimeFunc func() time.Time
				BeforeEach(func() {
					originalJwtTimeFunc = jwt.TimeFunc
				})
				AfterEach(func() {
					jwt.TimeFunc = originalJwtTimeFunc
				})

				It("is rejected", func() {
					jwt.TimeFunc = func() time.Time {
						return currentTime.Add(time.Hour + time.Second)
					}
					claims, err := factory.Parse(tokenString)
					Expect(err).To(MatchError(&intl.ExpiredJwtTokenError{}))
					Expect(claims).To(BeNil())
				})
			})

			Context("tampered token", func() {
				It("is rejected", func() {
					email1 := "test@example.com"
					email2 := "test2@example.com"

					email1Token, err := factory.Create(email1)
					Expect(err).To(BeNil())

					email2Token, err := factory.Create(email2)
					Expect(err).To(BeNil())

					email2Payload := strings.Split(email2Token, ".")[1]
					tamperedEmail1TokenData := strings.Split(email1Token, ".")
					tamperedEmail1TokenData[1] = email2Payload
					tamperedEmail1Token := strings.Join(tamperedEmail1TokenData, ".")

					parsedToken, err := factory.Parse(tamperedEmail1Token)
					Expect(err).To(MatchError(&intl.InvalidJwtTokenSignatureError{}))
					Expect(parsedToken).To(BeNil())
				})
			})
		})
	})
})
