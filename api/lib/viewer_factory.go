package lib

import (
	"api/ent/entgenerated"
	"api/ent/privacy/entviewer"
	"api/intl"
	"api/intl/intlgenerated"
	"api/privacy/viewer"
	"fmt"
	"net/http"
	"strings"
)

type (
	ViewerFactory interface {
		FromHttpAuthorizationHeader(*http.Request) (viewer.Viewer, intl.IntlError)
		// Try to load the viewer from an HttpRequest, otherwise return
		FromHttpRequestOrElseAnonymous(*entgenerated.Client, *http.Request) viewer.Viewer
	}

	defaultViewerFactory struct {
		ViewerFactory
	}

	NewViewerFactoryFunc func() ViewerFactory

	InvalidBearerAuthorizationHeaderError struct {
		intl.IntlError
	}

	InvalidAuthTokenError struct {
		intl.IntlError
	}
)

const (
	AUTHORIZATION_HTTP_HEADER_KEY = "Authorization"
	BEARER_FORMAT                 = "Bearer %s"
)

var NewViewerFactory NewViewerFactoryFunc = func() ViewerFactory {
	return &defaultViewerFactory{}
}

func (f defaultViewerFactory) FromHttpAuthorizationHeader(r *http.Request) (viewer.Viewer, intl.IntlError) {
	authorizationHeader := r.Header.Get(AUTHORIZATION_HTTP_HEADER_KEY)
	var token string
	_, err := fmt.Fscanf(strings.NewReader(authorizationHeader), BEARER_FORMAT, &token)
	if err != nil {
		return nil, &InvalidBearerAuthorizationHeaderError{}
	}

	userTokenFactory := NewUserTokenFactory()
	parsedToken, err := userTokenFactory.ParseAuthToken(token)
	if err != nil {
		return nil, &InvalidAuthTokenError{}
	}

	userId, exists := parsedToken.GetUserId()

	return entviewer.NewUserViewerFromId(userId, exists), nil
}

func (f defaultViewerFactory) FromHttpRequestOrElseAnonymous(client *entgenerated.Client, r *http.Request) viewer.Viewer {
	v, err := f.FromHttpAuthorizationHeader(r)
	if err == nil {
		// Loaded viewer from AuthorizationHeader, return it
		return v
	}

	// if cannot load viewr from the authorization header, try to refresh the auth using the refresh token cookie
	cookieReader := NewCookieReader(r)
	userFactory := NewUserFactory(client)
	userTokenFactory := NewUserTokenFactory()
	user, userLoadErr := userFactory.FromRefreshTokenCookieBypassPrivacy(r.Context(), cookieReader, userTokenFactory)
	if userLoadErr == nil {
		// Given we have a valid refresh token, we can use it and generate a new auth token to the client
		authToken, err := userTokenFactory.CreateAuthToken(user)
		if err == nil {
			r.WithContext(NewContextWithAuthToken(r.Context(), authToken))
			// To be safe, only return login viewer when auth token is successfully created. Otherwise, even with a
			// valid refresh token, we return anonymous viewer
			return entviewer.NewUserViewerFromUser(user)
		}
	}

	// Attempt to refresh the Authorization and obtain the viewer
	return entviewer.NewAnonymouseUserViewer()
}

// ------------- Error Methods ---------------

func (InvalidBearerAuthorizationHeaderError) Error() string {
	return "Invalid bearer authorization header"
}

func (InvalidBearerAuthorizationHeaderError) GetIntlKey() string {
	return intlgenerated.VIEWER__INVALID_LOGIN_CREDENTIAL
}

func (InvalidAuthTokenError) Error() string {
	return "Invalid auth token"
}

func (InvalidAuthTokenError) GetIntlKey() string {
	return intlgenerated.VIEWER__INVALID_LOGIN_CREDENTIAL
}
