package lib

import (
	"api/lib/cookies"
	"context"
	"net/http"
)

type (
	CookieReader interface {
		Read(name string) (*http.Cookie, error)
		ReadRefreshToken() (*http.Cookie, error)
	}

	DefaultCookieReader struct {
		CookieReader
		request *http.Request
	}

	NewCookieReaderFunc func(*http.Request) CookieReader

	cookieReaderCtxKey struct{}

	CookieReaderNotInContextError struct {
		error
	}
)

func NewContextWithCookieReader(parent context.Context, cookieReader CookieReader) context.Context {
	return context.WithValue(parent, cookieReaderCtxKey{}, cookieReader)
}

func GetCookieReaderFromContext(ctx context.Context) CookieReader {
	result, _ := ctx.Value(cookieReaderCtxKey{}).(CookieReader)
	return result
}

func RequireCookieReaderFromContext(ctx context.Context) (CookieReader, error) {
	result, ok := ctx.Value(cookieReaderCtxKey{}).(CookieReader)
	if !ok {
		return nil, &CookieReaderNotInContextError{}
	}
	return result, nil
}

var NewCookieReader = func(request *http.Request) CookieReader {
	return DefaultCookieReader{
		request: request,
	}
}

func (reader DefaultCookieReader) Read(name string) (*http.Cookie, error) {
	return reader.request.Cookie(name)
}

func (reader DefaultCookieReader) ReadRefreshToken() (*http.Cookie, error) {
	return reader.Read(cookies.REFRESH_TOKEN_COOKIE_NAME)
}

func (CookieReaderNotInContextError) Error() string {
	return "CookieReader is not set in the context. Remember to use utils.NewContextWithCookieReader"
}
