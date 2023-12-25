package test

import (
	"api/ent/entgenerated"
	"api/ent/entgenerated/privacy"
	"api/ent/internal"
	"api/ent/privacy/entviewer"
	"api/ent/privacy/rule"
	"api/privacy/viewer"
	"context"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AllowMutationIfOwner rule", func() {
	index := 1000

	var client *entgenerated.Client
	var subject privacy.MutationRule
	var owner *entgenerated.User
	var mutation *entgenerated.EmailCredentialMutation

	createViewerContext := func(user *entgenerated.User) context.Context {
		return viewer.NewContext(context.Background(), entviewer.NewUserViewerFromUser(user))
	}

	getIndex := func() int {
		index += 1
		return index
	}

	BeforeEach(func() {
		client = createTestEntClient()
		owner = internal.CreateTestUser(client)
		subject = rule.AllowMutationAfterApplyingOwnerFilter()
	})

	AfterEach(func() {
		client.Close()
	})

	Context("mutation", func() {
		var emailCredential *entgenerated.EmailCredential
		BeforeEach(func() {
			emailCredential = internal.CreateTestEmailCredentialForUser(
				client,
				owner,
				fmt.Sprintf("%d@example.com", getIndex()),
			)
			mutation = client.EmailCredential.UpdateOne(emailCredential).SetOwner(owner).Mutation()
		})

		It("allows owner mutation", func() {
			ctx := createViewerContext(owner)
			actual := subject.EvalMutation(ctx, mutation)
			Expect(actual).To(MatchError(privacy.Allowf("applied owner filter")))
		})

		It("allows non-owner mutation", func() {
			nonOwner := internal.CreateTestUser(client)
			ctx := createViewerContext(nonOwner)
			actual := subject.EvalMutation(ctx, mutation)
			Expect(actual).To(MatchError(privacy.Allow))
		})
	})

	// Context("creation", func() {
	// 	Context("unowned", func() {
	// 		var creation *entgenerated.EmailCredentialCreate

	// 		BeforeEach(func() {
	// 			creation = client.EmailCredential.Create()
	// 			mutation = creation.Mutation()
	// 		})

	// 		It("is allowed", func() {
	// 			ctx := createViewerContext(owner)
	// 			actual := subject.EvalMutation(ctx, mutation)
	// 			Expect(actual).To(MatchError(privacy.Allow))
	// 		})

	// 		It("cannot create", func() {
	// 			ctx := createViewerContext(owner)
	// 			creation.ExecX(ctx)
	// 		})
	// 	})

	// 	Context("owned", func() {
	// 		BeforeEach(func() {
	// 			mutation = client.EmailCredential.Create().SetOwner(owner).Mutation()
	// 		})

	// 		It("allow owner creation", func() {
	// 			ctx := createViewerContext(owner)
	// 			actual := subject.EvalMutation(ctx, mutation)
	// 			Expect(actual).To(MatchError(privacy.Allow))
	// 		})

	// 		It("allows non-owner creation", func() {
	// 			nonOwner := internal.CreateTestUser(client)
	// 			ctx := createViewerContext(nonOwner)
	// 			actual := subject.EvalMutation(ctx, mutation)
	// 			Expect(actual).To(MatchError(privacy.Allow))
	// 		})
	// 	})
	// })
})
