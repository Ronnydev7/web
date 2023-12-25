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

var _ = Describe("AllowIfOwned", func() {
	var (
		client *entgenerated.Client
		query  entgenerated.Query
		ctx    context.Context

		subject privacy.QueryMutationRule
	)

	BeforeEach(func() {
		client = internal.CreateEntClientForTest(TheT)
		query = client.SuperuserProfile.Query()
		ctx = context.Background()
		subject = rule.AllowIfOwnedByViewer()
	})

	AfterEach(func() {
		client.Close()
	})

	Context("viewer not in context", func() {
		It("skips", func() {
			Expect(subject.EvalQuery(ctx, query)).To(MatchError(privacy.Skipf("missing viewer in context")))
		})
	})

	Context("viewer is anonymous", func() {
		It("skips", func() {
			ctx = viewer.NewContext(ctx, entviewer.NewUserViewerFromUser(nil))
			Expect(subject.EvalQuery(ctx, query)).To(MatchError(privacy.Skipf("anonymous viewer")))
		})
	})

	Context("viewer is a user", func() {
		var (
			user *entgenerated.User
			ctx  context.Context
		)

		BeforeEach(func() {
			user, _ = internal.CreateTestSuperuser(client, "user")
			ctx = viewer.NewContext(context.Background(), entviewer.NewUserViewerFromUser(user))
		})

		Context("query of modal without owner relationship", func() {
			It("denies", func() {
				query = client.User.Query()
				Expect(subject.EvalQuery(ctx, query)).To(MatchError(privacy.Denyf("unexpected filter type *entgenerated.UserFilter")))
			})
		})

		Context("query of modal with owner relationship", func() {
			It("skips with filter applied", func() {
				query := client.SuperuserProfile.Query()
				Expect(subject.EvalQuery(ctx, query)).To(MatchError(privacy.Allowf("applied owner filter")))
			})
		})
	})
})
