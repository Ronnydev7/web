package test

import (
	"api/lib"
	"context"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("cookie_writer", func() {
	var httpResponseWriter http.ResponseWriter

	BeforeEach(func() {
		httpResponseWriter = httptest.NewRecorder()
	})

	It("NewCookieWriter creates new instance", func() {
		instance := lib.NewCookieWriter(httpResponseWriter)
		Expect(instance).NotTo(BeNil())
	})

	Describe("with default CookieWriter", func() {
		var cookieWriter lib.CookieWriter

		BeforeEach(func() {
			cookieWriter = lib.NewCookieWriter(httpResponseWriter)
		})

		It("can set to the context", func() {
			ctx := lib.NewContextWithCookieWriter(context.Background(), cookieWriter)
			Expect(ctx).NotTo(BeNil())
		})

		Describe("when context set", func() {
			var ctx context.Context
			BeforeEach(func() {
				ctx = lib.NewContextWithCookieWriter(context.Background(), cookieWriter)
			})

			Describe("GetCookieWriterFromContext", func() {
				It("returns the cookie writer", func() {
					actual := lib.GetCookieWriterFromContext(ctx)
					Expect(actual).NotTo(BeNil())
				})
			})

			Describe("RequireCookieWriterFromContext", func() {
				It("returns the cookie writer", func() {
					actual, err := lib.RequireCookieWriterFromContext(ctx)
					Expect(err).To(BeNil())
					Expect(actual).NotTo(BeNil())
				})
			})
		})

		Describe("when context not set", func() {
			Describe("GetCookieWriterFromContext", func() {
				It("returns nil when not set in context", func() {
					Expect(lib.GetCookieWriterFromContext(context.Background())).To(BeNil())
				})
			})

			Describe("RequireCookieWriterFromContext", func() {
				It("returns nil with error", func() {
					actual, err := lib.RequireCookieWriterFromContext(context.Background())
					Expect(err).To(MatchError(&lib.CookieWriterNotInContextError{}))
					Expect(actual).To(BeNil())
				})
			})
		})

		Describe("Write", func() {
			It("write the cookie to the response", func() {
				cookieWriter.Write(&http.Cookie{
					Name:     "testcookie",
					Value:    "testvalue",
					HttpOnly: true,
					Path:     "/",
				})
				cookieWriter.Write(&http.Cookie{
					Name:  "testcookie2",
					Value: "testvalue2",
				})
				Expect(httpResponseWriter.Header()["Set-Cookie"]).To(
					Equal([]string{"testcookie=testvalue; Path=/; HttpOnly", "testcookie2=testvalue2"}),
				)
			})
		})
	})
})
