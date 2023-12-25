package test

import (
	"api/ent/schema/validator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IsAlphanumericUnderscore", func() {
	testSubject := validator.IsAlphanumericUnderscore("testfield")

	It("allows alphanumeric uncerscore string", func() {
		Expect(testSubject("0legit_string")).To(BeNil())
	})

	It("disallows string with non-alphanumeric-underscore characters", func() {
		Expect(testSubject("not allowed")).To(MatchError("testfield must only consist of alphanumeric or underscore characters"))
	})
})
