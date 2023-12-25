package test

import (
	"api/ent/entgenerated"
	"api/ent/entgenerated/privacy"
	"api/ent/internal"
	"api/ent/privacy/rule"
	"api/ent/privacy/token"
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AllowIfContextHasValidPrivacyToken", func() {
	var client *entgenerated.Client
	var query entgenerated.Query

	BeforeEach(func() {
		client = internal.CreateEntClientForTest(TheT)
		query = client.User.Query()
	})

	AfterEach(func() {
		client.Close()
	})

	eval := func(ctx context.Context, applyFilter func(t token.PrivacyToken, filter privacy.Filter)) error {
		rule := rule.AllowAfterApplyingPrivacyTokenFilter(&token.PerformingLoginToken{}, applyFilter)
		return rule.EvalQuery(ctx, query)
	}

	Context("token of correct type", func() {
		It("allows after applying filter", func() {
			ctx := token.NewContextWithPerformingLoginToken(context.Background(), "test@example.com")
			Expect(eval(
				ctx,
				func(token.PrivacyToken, privacy.Filter) {},
			)).To(MatchError(privacy.Allowf("applied privacy token filter")))
		})
	})

	Context("token of incorrect type", func() {
		It("skips", func() {
			ctx := token.NewContextWithPerformingResetPasswordToken(context.Background(), "test@example.com")
			Expect(eval(
				ctx,
				func(token.PrivacyToken, privacy.Filter) {},
			)).To(MatchError(privacy.Skipf("no token found from context with type *token.PerformingLoginToken")))
		})
	})
})
