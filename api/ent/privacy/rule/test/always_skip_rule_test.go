package test

import (
	"api/ent/entgenerated/privacy"
	"api/ent/privacy/rule"
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AlwaysSkipRule", func() {
	It("always skips", func() {
		subject := rule.AlwaysSkipRule()
		client := createTestEntClient()
		actual := subject.EvalQuery(context.Background(), client.User.Query())
		Expect(actual).To(MatchError(privacy.Skip))
	})
})
