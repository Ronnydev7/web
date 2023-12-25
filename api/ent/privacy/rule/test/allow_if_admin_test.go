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

var _ = Describe("AllowIfAdmin rule", func() {
	var client *entgenerated.Client
	var subject privacy.QueryMutationRule

	BeforeEach(func() {
		client = internal.CreateEntClientForTest(TheT)
		subject = rule.AllowIfAdmin()
	})

	AfterEach(func() {
		client.Close()
	})

	Context("viewer is superuser", func() {
		var superuser *entgenerated.User
		var superuserCtx context.Context
		BeforeEach(func() {
			superuser, _ = internal.CreateTestSuperuser(client, "superuser")
			v, err := entviewer.NewSuperuserViewerFromUser(superuser)
			if err != nil {
				panic(err)
			}
			superuserCtx = viewer.NewContext(context.Background(), v)
		})

		It("is allowed", func() {
			actual := subject.EvalQuery(superuserCtx, client.User.Query())
			Expect(actual).To(MatchError(privacy.Allow))
		})
	})

	Context("viewer is normaluser", func() {
		var normaluser *entgenerated.User
		var normaluserCtx context.Context
		BeforeEach(func() {
			normaluser = internal.CreateTestUser(client)
			v := entviewer.NewUserViewerFromUser(normaluser)
			normaluserCtx = viewer.NewContext(context.Background(), v)
		})

		It("is skipped", func() {
			actual := subject.EvalQuery(normaluserCtx, client.User.Query())
			Expect(actual).To(MatchError(privacy.Skip))
		})
	})
})
