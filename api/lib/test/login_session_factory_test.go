package test

import (
	"api/ent/entgenerated"
	"api/ent/privacy/entviewer"
	"api/lib"
	"api/lib/cookies"
	"api/lib/libmocks"
	"api/lib/testutils"
	"api/privacy/viewer"
	"context"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("login_session_factory", func() {
	var client *entgenerated.Client
	var mockNewEntClient testutils.MockWrapper[lib.NewEntClientFunc]
	var owner *entgenerated.User
	var ownerContext context.Context
	var ownedLoginSession *entgenerated.LoginSession

	var currentTime time.Time
	var userTokenFactory lib.UserTokenFactory
	var mockCookieReader libmocks.CookieReader

	BeforeEach(func() {
		client, mockNewEntClient = startValidNewEntClientMock()
		owner = createTestUser(client)
		ownerContext = viewer.NewContext(context.Background(), entviewer.NewUserViewerFromUser(owner))
		ownedLoginSession = client.LoginSession.Create().
			SetOwner(owner).
			SetLastLoginTime(time.Now()).
			SaveX(ownerContext)

		userTokenFactory = lib.NewUserTokenFactory()
		refreshToken, err := userTokenFactory.CreateRefreshToken(ownedLoginSession)
		if err != nil {
			panic(err)
		}

		mockCookieReader = *libmocks.NewCookieReader(TheT)
		currentTime = time.Now()
		lib.TimeNow = func() time.Time {
			return currentTime
		}
		mockCookieReader.
			On("ReadRefreshToken").
			Return(cookies.NewRefreshTokenCookie(refreshToken, currentTime.Add(time.Minute), http.SameSiteDefaultMode), err)
	})

	AfterEach(func() {
		client.Close()
		mockNewEntClient.Stop()
	})

	Context("with valid owner viewer in the context", func() {
		Describe("FromRefreshTokenCookie", func() {
			It("returns the correct LoginSession", func() {
				subject := lib.NewLoginSessionFactory(ownerContext, client)

				actual, err := subject.FromRefreshTokenCookie(&mockCookieReader, userTokenFactory)
				Expect(err).To(BeNil())
				Expect(actual).NotTo(BeNil())
				Expect(actual.ID).To(Equal(ownedLoginSession.ID))
			})
		})
	})

	Context("without valid owner viewer in the context", func() {
		var nonOwner *entgenerated.User
		var subject lib.LoginSessionFactory

		BeforeEach(func() {
			nonOwner = createTestUser(client)
			subject = lib.NewLoginSessionFactory(
				viewer.NewContext(context.Background(), entviewer.NewUserViewerFromUser(nonOwner)),
				client,
			)
		})

		Describe("FromRefreshTokenCookie", func() {
			It("returns nil", func() {
				actual, err := subject.FromRefreshTokenCookie(&mockCookieReader, userTokenFactory)
				Expect(actual).To(BeNil())
				Expect(entgenerated.IsNotFound(err)).To(BeTrue())
			})
		})

		Describe("FromRefreshTokenCookieBypassPrivacy", func() {
			It("returns the vlaid LoginSession", func() {
				actual, err := subject.FromRefreshTokenCookieBypassPrivacy(&mockCookieReader, userTokenFactory)
				Expect(err).To(BeNil())
				Expect(actual).NotTo(BeNil())
				Expect(actual.ID).To(Equal(ownedLoginSession.ID))
			})

			Context("when the RT cookie is not present", func() {
				BeforeEach(func() {
					mockCookieReader.On("ReadRefreshToken").Return(nil, http.ErrNoCookie)
				})

				It("returns not found", func() {
					actual, err := subject.FromRefreshTokenCookie(&mockCookieReader, userTokenFactory)
					Expect(actual).To(BeNil())
					Expect(entgenerated.IsNotFound(err)).To(BeTrue())
				})
			})

			Context("when the token is expired", func() {
				var mockTimeNow testutils.MockWrapper[lib.TimeNowFunc]

				BeforeEach(func() {
					mockTimeNow = testutils.StartMockWrapper[lib.TimeNowFunc](
						func() time.Time {
							return currentTime.Add(lib.REFRESH_TOKEN_DURATION_MONTH*time.Hour*24*30 + 1)
						},
						lib.TimeNow,
						func(aTimeNow lib.TimeNowFunc) {
							lib.TimeNow = aTimeNow
						},
					)
				})
				AfterEach(func() {
					mockTimeNow.Stop()
				})

				It("returns not found", func() {
					actual, err := subject.FromRefreshTokenCookie(&mockCookieReader, userTokenFactory)
					Expect(actual).To(BeNil())
					Expect(entgenerated.IsNotFound(err)).To(BeTrue())
				})
			})
		})
	})
})
