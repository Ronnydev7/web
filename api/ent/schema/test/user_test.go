package test

import (
	"api/ent/entgenerated"
	"api/ent/entgenerated/privacy"
	"api/ent/entgenerated/user"
	"api/ent/internal"
	"api/ent/privacy/entviewer"
	"api/ent/privacy/token"
	"api/privacy/viewer"
	"api/privacy/viewer/viewermocks"
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("user", func() {
	var emptyCtx context.Context
	var userCtx context.Context
	var client *entgenerated.Client
	var existingUser *entgenerated.User

	BeforeEach(func() {
		emptyCtx = context.Background()
		client = internal.CreateEntClientForTest(TheT)
		existingUser = internal.CreateTestUser(client)
		userCtx = viewer.NewContext(context.Background(), entviewer.NewUserViewerFromUser(existingUser))
	})
	AfterEach(func() {
		client.Close()
	})

	Describe("query", func() {
		It("is always allowed", func() {
			actual, err := client.User.Get(context.Background(), existingUser.ID)
			Expect(err).To(BeNil())
			Expect(actual).To(BeAssignableToTypeOf(&entgenerated.User{}))
		})
	})

	Describe("mutation", func() {
		var superuserCtx context.Context
		BeforeEach(func() {
			v := viewermocks.Viewer{}
			v.On("IsAdmin").Return(true)
			v.On("IsSuperuser").Return(true)
			superuserCtx = viewer.NewContext(emptyCtx, &v)
		})

		Describe("creation", func() {
			expectCanCreate := func(ctx context.Context) {
				actual, err := client.User.Create().Save(ctx)
				Expect(err).To(BeNil())
				Expect(actual).To(BeAssignableToTypeOf(&entgenerated.User{}))
			}
			expectCreationDenied := func(ctx context.Context) {
				actual, err := client.User.Create().Save(ctx)
				Expect(err).To(MatchError(privacy.Deny))
				Expect(actual).To(BeNil())
			}

			Context("context has EmailSignupToken", func() {
				It("can create", func() {
					ctx := token.NewContextWithSignupToken(emptyCtx, &token.EmailSignupToken{Email: "test@example.com"})
					expectCanCreate(ctx)
				})
			})

			Context("context has no viewer", func() {
				It("cannot create", func() {
					expectCreationDenied(emptyCtx)
				})
			})

			Context("context has admin viewer", func() {
				It("can create", func() {
					expectCanCreate(superuserCtx)
				})
			})

			Context("context has normal viewer without token", func() {
				It("cannot create", func() {
					v := entviewer.NewUserViewerFromUser(existingUser)
					expectCreationDenied(viewer.NewContext(emptyCtx, v))
				})
			})
		})

		expectSuccess := func(actual *entgenerated.User, err error) {
			Expect(actual).To(BeAssignableToTypeOf(&entgenerated.User{}))
		}
		expectOneFailure := func(actual *entgenerated.User, err error) {
			Expect(err).To(MatchError(privacy.Deny))
			Expect(actual).To(BeNil())
		}
		expectFailure := func(actual int, err error) {
			Expect(err).To(MatchError(privacy.Deny))
			Expect(actual).To(Equal(0))
		}

		Describe("UpdateOne", func() {
			Context("by admin", func() {
				It("is allowed", func() {
					actual, err := client.User.UpdateOne(existingUser).Save(superuserCtx)
					expectSuccess(actual, err)
				})
			})
			Context("by non-admin", func() {
				It("is always denied", func() {
					actual, err := client.User.UpdateOne(existingUser).Save(userCtx)
					expectOneFailure(actual, err)
				})
			})
		})

		Describe("Update", func() {
			Context("by non-admin", func() {
				It("is always denied", func() {
					actual, err := client.User.Update().Save(userCtx)
					expectFailure(actual, err)
				})
			})
			Context("by admin", func() {
				It("is allowed", func() {
					actual, err := client.User.Update().Save(superuserCtx)
					Expect(actual).To(Equal(0))
					Expect(err).To(BeNil())
				})
			})
		})

		Describe("DeleteOne", func() {
			Context("by non-admin", func() {
				It("is always denied", func() {
					err := client.User.DeleteOne(existingUser).Exec(userCtx)
					Expect(err).To(MatchError(privacy.Deny))
				})
			})
			Context("by admin", func() {
				It("is allowed", func() {
					err := client.User.DeleteOne(existingUser).Exec(superuserCtx)
					Expect(err).To(BeNil())
				})
			})
		})

		Describe("Delete", func() {
			Context("by non-admin", func() {
				It("is always denied", func() {
					actual, err := client.User.Delete().Where(user.ID(existingUser.ID)).Exec(userCtx)
					expectFailure(actual, err)
				})
			})
			Context("by admin", func() {
				It("is allowed", func() {
					actual, err := client.User.Delete().Where(user.ID(existingUser.ID)).Exec(superuserCtx)
					Expect(actual).To(Equal(1))
					Expect(err).To(BeNil())
				})
			})
		})
	})
})
