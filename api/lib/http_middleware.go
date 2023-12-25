package lib

import (
	"api/privacy/viewer"
	"net/http"
)

func ApplyViewerMiddleware(resolveViewer func(r *http.Request) viewer.Viewer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v := resolveViewer(r)
			newCtx := viewer.NewContext(r.Context(), v)
			r = r.WithContext(newCtx)
			next.ServeHTTP(w, r)
		})
	}
}

func ApplyCookieReaderWriterMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				cookieWriter := NewCookieWriter(w)
				cookieReader := NewCookieReader(r)
				ctx := NewContextWithCookieWriter(r.Context(), cookieWriter)
				ctx = NewContextWithCookieReader(ctx, cookieReader)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
			},
		)
	}
}

func ApplyHttpErrorWriterMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				writer := NewHttpErrorWriter(w)
				ctx := NewContextWithHttpErrorWriter(r.Context(), writer)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
			},
		)
	}
}
