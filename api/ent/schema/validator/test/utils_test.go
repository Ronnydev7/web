package test

import (
	"api/ent/schema/validator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("utils", func() {
	Describe("ComposeValidators", func() {
		When("composing multiple validators", func() {
			testSubject := validator.ComposeValidators[string](
				validator.MaxUtf8RuneCount("test field", 5),
				validator.IsAlphanumericUnderscore("testfield"),
			)
			When("input is valid", func() {
				It("does not return an error", func() {
					Expect(testSubject("abcde")).To(BeNil())
				})
			})

			When("input is invalid", func() {
				It("returns the first validation error", func() {
					Expect(testSubject("abcdef")).To(MatchError("test field must not exceed 5 characters"))
				})
			})
		})
	})
})
