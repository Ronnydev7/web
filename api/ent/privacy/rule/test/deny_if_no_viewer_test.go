package test

import (
	"api/ent/entgenerated/privacy"
	"api/ent/internal"
	"api/ent/privacy/entviewer"
	"api/ent/privacy/rule"
	"api/privacy/viewer"
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DenyIfNoViewer", func() {
	eval := func(ctx context.Context) error {
		client := internal.CreateEntClientForTest(TheT)
		defer client.Close()
		rule := rule.DenyIfNoViewer()
		return rule.EvalQuery(ctx, client.User.Query())
	}

	Context("has viewer", func() {
		It("skips", func() {
			ctx := viewer.NewContext(context.Background(), entviewer.NewUserViewerFromUser(nil))
			actual := eval(ctx)
			Expect(actual).To(MatchError(privacy.Skip))
		})
	})

	Context("no viewer", func() {
		It("denies", func() {
			actual := eval(context.Background())
			Expect(actual).To(MatchError(privacy.Deny))
		})
	})
})
