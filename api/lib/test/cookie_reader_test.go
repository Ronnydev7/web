package test

import (
	"api/lib"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("cookie_reader", func() {
	var httpRequest *http.Request

	BeforeEach(func() {
		httpRequest = httptest.NewRequest(http.MethodGet, "/query", nil)
	})

	It("NewCookieReader creates new instance", func() {
		instance := lib.NewCookieReader(httpRequest)
		Expect(instance).To(BeAssignableToTypeOf(lib.DefaultCookieReader{}))
	})

	Describe("with default CookieReader", func() {
		var cookieReader lib.CookieReader
		BeforeEach(func() {
			cookieReader = lib.NewCookieReader(httpRequest)
		})

		Describe("when context set", func() {
			var ctx context.Context
			BeforeEach(func() {
				ctx = lib.NewContextWithCookieReader(context.Background(), cookieReader)
			})

			Describe("GetCookieWriterFromContext", func() {
				It("returns the cookie reader", func() {
					actual := lib.GetCookieReaderFromContext(ctx)
					Expect(actual).To(BeAssignableToTypeOf(lib.DefaultCookieReader{}))
				})
			})

			Describe("RequireCookieReaderFromContext", func() {
				It("returns the cookie reader", func() {
					actual, err := lib.RequireCookieReaderFromContext(ctx)
					Expect(err).To(BeNil())
					Expect(actual).To(BeAssignableToTypeOf(lib.DefaultCookieReader{}))
				})
			})
		})
		Describe("when context not set", func() {
			Describe("GetCookieWriterFromContext", func() {
				It("returns the cookie reader", func() {
					actual := lib.GetCookieReaderFromContext(context.Background())
					Expect(actual).To(BeNil())
				})
			})

			Describe("RequireCookieReaderFromContext", func() {
				It("returns the cookie reader", func() {
					actual, err := lib.RequireCookieReaderFromContext(context.Background())
					Expect(err).To(MatchError(&lib.CookieReaderNotInContextError{}))
					Expect(actual).To(BeNil())
				})
			})

			Describe("Read", func() {
				Context("when cookie is found", func() {
					BeforeEach(func() {
						httpRequest.AddCookie(&http.Cookie{
							Name:  "testcookie",
							Value: "testcookievalue",
						})
					})

					It("read the cookie from request", func() {
						result, err := cookieReader.Read("testcookie")
						Expect(err).To(BeNil())
						Expect(result).To(Equal(&http.Cookie{
							Name:  "testcookie",
							Value: "testcookievalue",
						}))
					})
				})

				Context("when cookie is not found", func() {
					It("returns nil", func() {
						result, err := cookieReader.Read("testcookie")
						Expect(err).To(MatchError(errors.New("http: named cookie not present")))
						Expect(result).To(BeNil())
					})
				})
			})
		})
	})
})
