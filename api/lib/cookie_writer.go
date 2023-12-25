package lib

import (
	"context"
	"net/http"
)

type (
	CookieWriter interface {
		Write(*http.Cookie)
	}

	DefaultCookieWriter struct {
		CookieWriter
		writer http.ResponseWriter
	}

	cookieWriterCtxKey struct{}

	CookieWriterNotInContextError struct {
		error
	}
)

func NewContextWithCookieWriter(parent context.Context, cookieWriter CookieWriter) context.Context {
	return context.WithValue(parent, cookieWriterCtxKey{}, cookieWriter)
}

func GetCookieWriterFromContext(ctx context.Context) CookieWriter {
	result, _ := ctx.Value(cookieWriterCtxKey{}).(CookieWriter)
	return result
}

func RequireCookieWriterFromContext(ctx context.Context) (CookieWriter, error) {
	result, ok := ctx.Value(cookieWriterCtxKey{}).(CookieWriter)
	if !ok {
		return nil, &CookieWriterNotInContextError{}
	}
	return result, nil
}

var NewCookieWriter = func(writer http.ResponseWriter) CookieWriter {
	return &DefaultCookieWriter{
		writer: writer,
	}
}

func (writer DefaultCookieWriter) Write(cookie *http.Cookie) {
	http.SetCookie(writer.writer, cookie)
}
