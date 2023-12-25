package test

import (
	"api/ent/entgenerated"
	"api/ent/entgenerated/privacy"
	"api/ent/internal"
	"api/ent/privacy/entviewer"
	"api/ent/privacy/token"
	"api/privacy/viewer"
	"api/privacy/viewer/viewermocks"
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("login_session", func() {
	var client *entgenerated.Client
	BeforeEach(func() {
		client = createTestEntClient()
	})
	AfterEach(func() {
		client.Close()
	})

	Describe("mutation", func() {
		var owner *entgenerated.User
		var nonOwner *entgenerated.User
		BeforeEach(func() {
			owner = internal.CreateTestUser(client)
			nonOwner = internal.CreateTestUser(client)
		})

		Describe("creation", func() {
			expectCreated := func(ctx context.Context) {
				result := client.LoginSession.Create().
					SetOwner(owner).
					SetLastLoginTime(time.Now()).
					SaveX(ctx)
				Expect(result).ToNot(BeNil())
			}

			Context("with owner as the viewer", func() {
				It("can be created", func() {
					v := entviewer.NewUserViewerFromUser(owner)
					ctx := viewer.NewContext(context.Background(), v)
					expectCreated(ctx)
				})
			})

			Context("with admin as the viewer", func() {
				It("can be created", func() {
					v := viewermocks.Viewer{}
					v.On("IsAdmin").Return(true)
					ctx := viewer.NewContext(context.Background(), &v)
					expectCreated(ctx)
				})
			})

			Context("with anonymous user as the viewer", func() {
				It("cannot be created", func() {
					v := entviewer.NewUserViewerFromUser(nil)
					ctx := viewer.NewContext(context.Background(), v)
					result, err := client.LoginSession.Create().
						SetOwner(owner).
						SetLastLoginTime(time.Now()).
						Save(ctx)
					Expect(err).To(MatchError(privacy.Deny))
					Expect(result).To(BeNil())
				})
			})
		})

		Describe("deletion", func() {
			var existingSession *entgenerated.LoginSession
			var ownerCtx context.Context

			BeforeEach(func() {
				v := entviewer.NewUserViewerFromUser(owner)
				ownerCtx = viewer.NewContext(context.Background(), v)
				existingSession = client.LoginSession.Create().
					SetOwner(owner).
					SetLastLoginTime(time.Now()).
					SaveX(ownerCtx)
			})
			expectDeleted := func(ctx context.Context) {
				client.LoginSession.DeleteOne(existingSession).ExecX(ctx)
				session, err := client.LoginSession.Get(ctx, existingSession.ID)
				Expect(entgenerated.IsNotFound(err)).To(BeTrue())
				Expect(session).To(BeNil())
			}

			expectNotDeleted := func(ctx context.Context) {
				err := client.LoginSession.DeleteOne(existingSession).Exec(ctx)
				Expect(entgenerated.IsNotFound(err)).To(BeTrue())

				sessionAfterDelete, err := client.LoginSession.Get(ownerCtx, existingSession.ID)
				Expect(err).To(BeNil())
				Expect(sessionAfterDelete).NotTo(BeNil())
			}

			Context("with owner as the viewer", func() {
				It("can be deleted", func() {
					expectDeleted(ownerCtx)
				})
			})

			Context("with admin as the viewer", func() {
				It("can be deleted", func() {
					v := viewermocks.Viewer{}
					v.On("IsAdmin").Return(true)
					ctx := viewer.NewContext(context.Background(), &v)
					expectDeleted(ctx)
				})
			})

			Context("with non-owner as the viewer", func() {
				It("cannot be deleted", func() {
					v := entviewer.NewUserViewerFromUser(nonOwner)
					ctx := viewer.NewContext(context.Background(), v)
					expectNotDeleted(ctx)
				})
			})

			Context("with anonymous user as the viewer", func() {
				It("cannot be deleted", func() {
					v := entviewer.NewUserViewerFromUser(nil)
					ctx := viewer.NewContext(context.Background(), v)
					err := client.LoginSession.DeleteOne(existingSession).Exec(ctx)
					Expect(err).To(MatchError(privacy.Deny))

					sessionAfterDelete, err := client.LoginSession.Get(ownerCtx, existingSession.ID)
					Expect(err).To(BeNil())
					Expect(sessionAfterDelete).NotTo(BeNil())
				})
			})
		})
	})

	Describe("query", func() {
		var user *entgenerated.User
		var existingSession *entgenerated.LoginSession
		BeforeEach(func() {
			user = internal.CreateTestUser(client)
			v := entviewer.NewUserViewerFromUser(user)
			ctx := viewer.NewContext(context.Background(), v)
			internal.CreateTestEmailCredentialForUser(client, user, "test@example.com")
			existingSession = client.LoginSession.Create().
				SetOwner(user).
				SetLastLoginTime(time.Now()).
				SaveX(ctx)
		})

		expectFound := func(ctx context.Context) {
			queryResult, err := client.LoginSession.Get(ctx, existingSession.ID)
			Expect(err).To(BeNil())
			Expect(queryResult).NotTo(BeNil())
			Expect(queryResult.ID).To(Equal(existingSession.ID))
		}

		expectNotFound := func(ctx context.Context) {
			queryResult, err := client.LoginSession.Get(ctx, existingSession.ID)
			Expect(entgenerated.IsNotFound(err)).To(BeTrue())
			Expect(queryResult).To(BeNil())
		}

		expectError := func(ctx context.Context, expectedError error) {
			queryResult, err := client.LoginSession.Get(ctx, existingSession.ID)
			Expect(err).To(MatchError(expectedError))
			Expect(queryResult).To(BeNil())
		}

		Context("by owner", func() {
			It("is allowed", func() {
				ctx := viewer.NewContext(context.Background(), entviewer.NewUserViewerFromUser(user))
				expectFound(ctx)
			})
		})

		Context("by context with performing login token", func() {
			Context("token is correct", func() {
				It("is allowed", func() {
					ctx := token.NewContextWithPerformingLoginToken(context.Background(), "test@example.com")
					expectFound(ctx)
				})
			})

			Context("token is incorrect", func() {
				It("is not found", func() {
					ctx := token.NewContextWithPerformingLoginToken(context.Background(), "test2@example.com")
					expectNotFound(ctx)
				})
			})
		})

		Context("by context with performing auth refresh token", func() {
			Context("with valid token", func() {
				It("is allowed", func() {
					ctx := token.NewContextWithPerformingAuthRefreshToken(context.Background(), existingSession.ID)
					expectFound(ctx)
				})
			})

			Context("with invalid token", func() {
				It("is not found", func() {
					ctx := token.NewContextWithPerformingAuthRefreshToken(context.Background(), existingSession.ID+1)
					expectNotFound(ctx)
				})
			})
		})

		Context("by non-owner viewer", func() {
			It("is not found", func() {
				user2 := internal.CreateTestUser(client)
				ctx := viewer.NewContext(context.Background(), entviewer.NewUserViewerFromUser(user2))
				expectNotFound(ctx)
			})
		})

		Context("by anonymous viewer", func() {
			It("is not found", func() {
				ctx := viewer.NewContext(context.Background(), entviewer.NewUserViewerFromUser(nil))
				expectError(ctx, privacy.Deny)
			})
		})

		Context("by admin viewer", func() {
			It("is allowed", func() {
				v := viewermocks.Viewer{}
				v.On("IsAdmin").Return(true)
				expectFound(viewer.NewContext(context.Background(), &v))
			})
		})
	})
})
