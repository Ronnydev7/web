package test

import (
	"api/ent/entgenerated"
	"api/ent/entgenerated/privacy"
	"api/ent/internal"
	"api/ent/privacy/entviewer"
	"api/ent/privacy/rule"
	"api/privacy/viewer"
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DenyIfChangeOwnership", func() {
	var client *entgenerated.Client
	var user *entgenerated.User
	var subject privacy.MutationRule

	BeforeEach(func() {
		client = internal.CreateEntClientForTest(TheT)
		user = internal.CreateTestUser(client)
		subject = rule.DenyIfChangeOwnership()
	})

	AfterEach(func() {
		client.Close()
	})

	eval := func(ctx context.Context, mutation entgenerated.Mutation) error {
		return subject.EvalMutation(ctx, mutation)
	}

	Context("unexpected mutation", func() {
		It("skips", func() {
			mutation := client.User.UpdateOne(user).Mutation()
			actual := eval(context.Background(), mutation)
			Expect(actual).To(MatchError(privacy.Skip))
		})
	})

	Context("expected mutation", func() {
		var ctx context.Context
		var mutation *entgenerated.EmailCredentialMutation
		BeforeEach(func() {
			ctx = viewer.NewContext(context.Background(), entviewer.NewUserViewerFromUser(user))
			emailCredential := internal.CreateTestEmailCredentialForUser(client, user, "test@example.com")
			mutation = client.EmailCredential.UpdateOne(emailCredential).Mutation()
		})

		Context("not setting new owner", func() {
			It("skips", func() {
				actual := subject.EvalMutation(ctx, mutation)
				Expect(actual).To(MatchError(privacy.Skip))
			})
		})

		Context("changing owner", func() {
			It("denies", func() {
				user2 := internal.CreateTestUser(client)
				mutation.SetOwnerID(user2.ID)
				actual := eval(ctx, mutation)
				Expect(actual).To(MatchError(privacy.Deny))
			})
		})

		Context("setting the same owner id", func() {
			It("skips", func() {
				mutation.SetOwnerID(user.ID)
				actual := eval(ctx, mutation)
				Expect(actual).To(MatchError(privacy.Skip))
			})
		})
	})
})
