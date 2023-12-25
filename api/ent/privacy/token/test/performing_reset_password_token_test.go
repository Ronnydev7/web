package test

import (
	"api/ent/privacy/token"
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("performing_reset_password_token", func() {
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.Background()
	})

	Describe("NewContextWIthPerformingResetPasswordToken, PerformingResetPasswordTokenFromContext", func() {
		Context("context does not have the token", func() {
			It("returns nil", func() {
				Expect(token.PerformingResetPasswordTokenFromContext(ctx)).To(BeNil())
			})
		})

		Context("context has the token", func() {
			It("returns the token", func() {
				email := "test@example.com"
				newCtx := token.NewContextWithPerformingResetPasswordToken(ctx, email)
				actual := token.PerformingResetPasswordTokenFromContext(newCtx)
				Expect(actual).To(BeAssignableToTypeOf(&token.PerformingResetPasswordToken{}))
				Expect(actual.Email).To(Equal(email))
			})
		})
	})

	Describe("GetContextKey", func() {
		It("returns the correct context key", func() {
			email := "test@example.com"
			newCtx := token.NewContextWithPerformingResetPasswordToken(ctx, email)

			actual, ok := newCtx.Value(token.PerformingResetPasswordToken{}.GetContextKey()).(*token.PerformingResetPasswordToken)
			Expect(ok).To(BeTrue())
			Expect(actual.Email).To(Equal(email))
		})
	})
})
