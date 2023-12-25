package test

import (
	"api/ent/schema/validator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MaxUtf8RuneCount", func() {
	testSubject := validator.MaxUtf8RuneCount("test field", 5)

	It("Allows an ASCII string not longer than expected rune count", func() {
		Expect(testSubject("abcde")).To(BeNil())
	})

	It("Allows international string not longer than expected rune count", func() {
		Expect(testSubject("一二三四五")).To(BeNil())
	})

	It("Does not allow ASCII string longer than expected count", func() {
		Expect(testSubject("abcdef")).To(MatchError(("test field must not exceed 5 characters")))
	})

	It("Does not allow international string longer than expected count", func() {
		Expect(testSubject("一二三四五六")).To(MatchError("test field must not exceed 5 characters"))
	})
})
