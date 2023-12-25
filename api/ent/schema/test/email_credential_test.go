package test

import (
	"api/ent/entgenerated"
	"api/ent/internal"
	"api/ent/privacy/entviewer"
	"api/ent/privacy/token"
	"api/privacy/viewer"
	"api/privacy/viewer/viewermocks"
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("email_credential", func() {
	var client *entgenerated.Client
	var owner *entgenerated.User

	BeforeEach(func() {
		client = internal.CreateEntClientForTest(TheT)
		owner = internal.CreateTestUser(client)
	})

	AfterEach(func() {
		client.Close()
	})

	Context("query", func() {
		var subject *entgenerated.EmailCredential

		BeforeEach(func() {
			ctx := viewer.NewContext(context.Background(), &testViewer{})
			subject = client.EmailCredential.
				Create().
				SetEmail("test@example.com").
				SetAlgorithm("bcrypt").
				SetPasswordHash([]byte("hash")).
				SetOwner(owner).
				SaveX(ctx)
		})

		expectSuccessfulResult := func(ctx context.Context) {
			actual, err := client.EmailCredential.Get(ctx, subject.ID)
			Expect(err).To(BeNil())
			Expect(actual).NotTo(BeNil())
			Expect(actual.ID).To(Equal(subject.ID))
		}

		Context("by admin", func() {
			It("returns the email credential", func() {
				v := viewermocks.Viewer{}
				v.On("IsAdmin").Return(true)
				expectSuccessfulResult(viewer.NewContext(context.Background(), &v))
			})
		})

		Context("by owner", func() {
			It("returns the email credential", func() {
				v := entviewer.NewUserViewerFromUser(owner)
				expectSuccessfulResult(viewer.NewContext(context.Background(), v))
			})
		})

		Context("by context with PerformingLoginToken", func() {
			It("returns the email credential", func() {
				ctx := token.NewContextWithPerformingLoginToken(context.Background(), "test@example.com")
				expectSuccessfulResult(ctx)
			})
		})

		expectFailedResult := func(ctx context.Context) {
			actual, err := client.EmailCredential.Get(ctx, subject.ID)
			Expect(entgenerated.IsNotFound(err)).To(BeTrue())
			Expect(actual).To(BeNil())
		}

		Context("by context with incorrect PerformingLoginToken", func() {
			It("returns nil", func() {
				ctx := token.NewContextWithPerformingLoginToken(context.Background(), "test2@example.com")
				expectFailedResult(ctx)
			})
		})

		Context("by non-owner", func() {
			It("returns nil", func() {
				user2 := internal.CreateTestUser(client)
				ctx := viewer.NewContext(context.Background(), entviewer.NewUserViewerFromUser(user2))
				expectFailedResult(ctx)
			})
		})

		Context("by context with correct PerformingResetPasswordToken", func() {
			It("returns the email credential", func() {
				ctx := token.NewContextWithPerformingResetPasswordToken(context.Background(), "test@example.com")
				expectSuccessfulResult(ctx)
			})
		})

		Context("by context with incorrect PerformingResetPasswordToken", func() {
			It("returns the email credential", func() {
				ctx := token.NewContextWithPerformingResetPasswordToken(context.Background(), "test2@example.com")
				expectFailedResult(ctx)
			})
		})
	})

	Context("mutation", func() {
		Context("create with EmailSignupToken", func() {
			It("should succeed", func() {
				v := entviewer.NewUserViewerFromUser(nil)
				ctx := viewer.NewContext(context.Background(), v)
				emailSignupToken := token.EmailSignupToken{
					Email: "test@example.com",
				}
				ctx = token.NewContextWithSignupToken(ctx, &emailSignupToken)
				_, err := client.EmailCredential.
					Create().
					SetEmail("test@example.com").
					SetAlgorithm("bcrypt").
					SetPasswordHash([]byte("hash")).
					SetOwner(owner).
					Save(ctx)
				Expect(err).To(BeNil())
			})
		})

		Context("create with invalid EmailSignupToken", func() {
			It("should fail", func() {
				v := entviewer.NewUserViewerFromUser(nil)
				ctx := viewer.NewContext(context.Background(), v)
				emailSignupToken := token.EmailSignupToken{
					Email: "test2@example.com",
				}
				ctx = token.NewContextWithSignupToken(ctx, &emailSignupToken)
				_, err := client.EmailCredential.
					Create().
					SetEmail("test@example.com").
					SetAlgorithm("bcrypt").
					SetPasswordHash([]byte("hash")).
					SetOwner(owner).
					Save(ctx)
				Expect(err).ToNot(BeNil())
			})
		})

		Context("create without token", func() {
			It("should fail", func() {
				v := entviewer.NewUserViewerFromUser(nil)
				ctx := viewer.NewContext(context.Background(), v)
				_, err := client.EmailCredential.
					Create().
					SetEmail("test@example.com").
					SetAlgorithm("bcrypt").
					SetPasswordHash([]byte("hash")).
					SetOwner(owner).
					Save(ctx)
				Expect(err).ToNot(BeNil())
			})
		})

		Context("mutate if owner", func() {
			It("should succeed", func() {
				emailCredential := internal.CreateTestEmailCredentialForUser(client, owner, "test@example.com")
				v := entviewer.NewUserViewerFromUser(owner)
				ctx := viewer.NewContext(context.Background(), v)
				_, err := client.EmailCredential.
					UpdateOne(emailCredential).
					SetEmail("test2@example.com").
					Save(ctx)
				Expect(err).To(BeNil())
			})
		})

		Context("mutate if not owner", func() {
			It("should fail", func() {
				emailCredential := internal.CreateTestEmailCredentialForUser(client, owner, "test@example.com")
				otherUser := internal.CreateTestUser(client)
				v := entviewer.NewUserViewerFromUser(otherUser)
				ctx := viewer.NewContext(context.Background(), v)
				_, err := client.EmailCredential.
					UpdateOne(emailCredential).
					SetEmail("test2@example.com").
					Save(ctx)
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
