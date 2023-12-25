package test

import (
	"api/ent/privacy/token"
	"context"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("performing_login_token", func() {
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.Background()
	})

	Describe("NewContextWithPerformingLoginToken, PerformingLoginTokenFromContext", func() {
		Context("context does not have the token", func() {
			It("returns nil", func() {
				Expect(token.PerformingLoginTokenFromContext(ctx)).To(BeNil())
			})
		})

		Context("context has token", func() {
			BeforeEach(func() {
				email := "test@example.com"
				ctx = token.NewContextWithPerformingLoginToken(ctx, email)
			})
			It("returns the token", func() {
				actual := token.PerformingLoginTokenFromContext(ctx)
				Expect(actual).To(BeAssignableToTypeOf(&token.PerformingLoginToken{}))
				Expect(actual.Email).To(Equal("test@example.com"))
			})

			Describe("GetContextKey", func() {
				It("returns the correct context key", func() {
					actual, ok := ctx.Value(token.PerformingLoginToken{}.GetContextKey()).(*token.PerformingLoginToken)
					Expect(ok).To(BeTrue())
					Expect(actual.Email).To(Equal("test@example.com"))
				})
			})
		})
	})

	Describe("GetPerformingLoginTokenValidatorFunc", func() {
		var subject func(token.PrivacyToken) error

		BeforeEach(func() {
			subject = token.GetPerformingLoginTokenValidatorFunc("test@example.com")
		})

		Context("token of incorrect type", func() {
			It("returns error", func() {
				Expect(subject(&token.PerformingResetPasswordToken{Email: "test@example.com"})).
					To(MatchError(errors.New("invalid privacy token type *token.PerformingResetPasswordToken")))
			})
		})

		Context("token of correct type", func() {
			Context("with correct data", func() {
				It("returns nil", func() {
					Expect(subject(&token.PerformingLoginToken{Email: "test@example.com"})).To(BeNil())
				})
			})
			Context("with incorrect data", func() {
				It("returns nil", func() {
					Expect(subject(&token.PerformingLoginToken{Email: "test2@example.com"})).To(
						MatchError(errors.New("privacy token has incorrect email")),
					)
				})
			})
		})
	})
})
