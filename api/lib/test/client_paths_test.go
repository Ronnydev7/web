package test

import (
	"api/config/configmocks"
	"api/lib"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("client_path", func() {
	var subject lib.ClientPathFactory
	BeforeEach(func() {
		c := &configmocks.UrlConfig{}
		c.On("GetProtocol").Return("https")
		c.On("GetHostname").Return("test.hobbytrace.com")

		subject = lib.NewClientPathFactory(c)
	})

	Describe("CreateConfirmEmailSignupPath", func() {
		It("creates a valid url", func() {
			actual := subject.CreateConfirmEmailSignupPath("token")
			Expect(actual.String()).To(Equal("https://test.hobbytrace.com/confirm-email-signup?email_signup_token=token"))
		})
	})

	Describe("CreateResetpasswordUrl", func() {
		It("creates a valid url", func() {
			Expect(subject.CreateResetPasswordUrl("token").String()).To(
				Equal("https://test.hobbytrace.com/reset-password?reset_password_token=token"),
			)
		})
	})
})
