package test

import (
	"api/ent/entgenerated"
	"api/ent/privacy/entviewer"
	"api/lib"
	"api/lib/libmocks"
	"api/lib/testutils"
	"api/privacy/viewer"
	"context"
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("user_factory", func() {
	var client *entgenerated.Client

	BeforeEach(func() {
		client = createEntClientForTest()
	})

	Describe("FromRefreshTokenCookieBypassPrivacy", func() {
		var mockLoginSessionFactory libmocks.LoginSessionFactory
		var mockNewLoginSessionFactory testutils.MockWrapper[lib.NewLoginSessionFactoryFunc]

		BeforeEach(func() {
			mockLoginSessionFactory = *libmocks.NewLoginSessionFactory(TheT)
			mockNewLoginSessionFactory = testutils.StartMockWrapper[lib.NewLoginSessionFactoryFunc](
				func(_ context.Context, _ *entgenerated.Client) lib.LoginSessionFactory {
					return &mockLoginSessionFactory
				},
				lib.NewLoginSessionFactory,
				func(aNew lib.NewLoginSessionFactoryFunc) {
					lib.NewLoginSessionFactory = aNew
				},
			)
		})

		AfterEach(func() {
			mockNewLoginSessionFactory.Stop()
		})

		Context("when login session is not found", func() {
			BeforeEach(func() {
				mockLoginSessionFactory.On(
					"FromRefreshTokenCookieBypassPrivacy",
					mock.Anything,
					mock.Anything,
				).Return(nil, http.ErrNoCookie)
			})

			It("returns nil", func() {
				subject := lib.NewUserFactory(client)
				httpRequest := httptest.NewRequest(http.MethodGet, "/query", nil)
				cookieReader := lib.NewCookieReader(httpRequest)
				userTokenFactory := lib.NewUserTokenFactory()
				actual, err := subject.FromRefreshTokenCookieBypassPrivacy(context.Background(), cookieReader, userTokenFactory)
				Expect(err).To(MatchError(http.ErrNoCookie))
				Expect(actual).To(BeNil())
			})
		})

		Context("when login session is found", func() {
			var testUser *entgenerated.User

			BeforeEach(func() {
				testUser = createTestUser(client)

				testUserViewer := viewer.NewContext(context.Background(), entviewer.NewUserViewerFromUser(testUser))
				loginSession := client.LoginSession.Create().
					SetOwner(testUser).
					SetLastLoginTime(time.Now()).
					SaveX(testUserViewer)

				mockLoginSessionFactory.On(
					"FromRefreshTokenCookieBypassPrivacy",
					mock.Anything,
					mock.Anything,
				).Return(loginSession, nil)
			})

			It("returns the user", func() {
				subject := lib.NewUserFactory(client)
				httpRequest := httptest.NewRequest(http.MethodGet, "/query", nil)
				cookieReader := lib.NewCookieReader(httpRequest)
				userTokenFactory := lib.NewUserTokenFactory()
				actual, err := subject.FromRefreshTokenCookieBypassPrivacy(context.Background(), cookieReader, userTokenFactory)
				Expect(err).To(BeNil())
				Expect(actual.ID).To(Equal(testUser.ID))
			})
		})
	})
})
