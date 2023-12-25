package test

import (
	"api/ent/entgenerated"
	"api/ent/entgenerated/privacy"
	"api/ent/internal"
	"api/ent/privacy/rule"
	"api/privacy/viewer"
	"api/privacy/viewer/viewermocks"
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AllowIfSuperuser", func() {
	var client *entgenerated.Client
	var subject privacy.QueryMutationRule

	BeforeEach(func() {
		client = internal.CreateEntClientForTest(TheT)
		subject = rule.AllowIfSuperuser()
	})

	AfterEach(func() {
		client.Close()
	})

	doTest := func(v viewer.Viewer) error {
		ctx := viewer.NewContext(context.Background(), v)
		return subject.EvalQuery(ctx, client.User.Query())
	}

	Context("viewer is superuser", func() {
		It("should allow", func() {
			v := viewermocks.Viewer{}
			v.On("IsSuperuser").Return(true)
			actual := doTest(&v)
			Expect(actual).To(MatchError(privacy.Allow))
		})
	})

	Context("viewer is not superuser", func() {
		It("should skip", func() {
			v := viewermocks.Viewer{}
			v.On("IsSuperuser").Return(false)
			actual := doTest(&v)
			Expect(actual).To(MatchError(privacy.Skip))
		})
	})
})
