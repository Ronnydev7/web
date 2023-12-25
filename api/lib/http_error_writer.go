package lib

import (
	"context"
	"net/http"
)

type (
	HttpErrorWriter interface {
		Write(msg string, statusCode int)
		WriteUnauthorized()
	}

	DefaultHttpErrorWriter struct {
		HttpErrorWriter
		responseWriter http.ResponseWriter
	}

	HttpErrorWriterFactory = func(http.ResponseWriter) HttpErrorWriter

	httpErrorWriterCtxKey struct{}

	HttpErrorWriterNotInContextError struct {
		error
	}
)

var NewHttpErrorWriter HttpErrorWriterFactory = func(responseWriter http.ResponseWriter) HttpErrorWriter {
	return &DefaultHttpErrorWriter{
		responseWriter: responseWriter,
	}
}

func (writer DefaultHttpErrorWriter) Write(msg string, statusCode int) {
	writer.responseWriter.Header().Set("X-Content-Type-Options", "nosniff")
	writer.responseWriter.WriteHeader(statusCode)
}

func (writer DefaultHttpErrorWriter) WriteUnauthorized() {
	writer.Write(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

func NewContextWithHttpErrorWriter(parent context.Context, writer HttpErrorWriter) context.Context {
	return context.WithValue(parent, httpErrorWriterCtxKey{}, writer)
}

func GetHttpErrorWriterFromContext(ctx context.Context) HttpErrorWriter {
	result, _ := ctx.Value(httpErrorWriterCtxKey{}).(HttpErrorWriter)
	return result
}

func RequireHttpErrorWriterFromContext(ctx context.Context) (HttpErrorWriter, error) {
	result, ok := ctx.Value(httpErrorWriterCtxKey{}).(HttpErrorWriter)
	if !ok {
		return nil, &HttpErrorWriterNotInContextError{}
	}
	return result, nil
}

func (HttpErrorWriterNotInContextError) Error() string {
	return "HttpErrorWriter not found in context. Remember to call utils.NewContextWithHttpErrorWriter"
}
