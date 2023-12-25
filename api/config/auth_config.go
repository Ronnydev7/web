package config

import "net/http"

type (
	AuthConfig interface {
		GetRefreshTokenCookieSameSiteMode() http.SameSite
	}

	defaultAuthConfig struct {
		AuthConfig
	}

	AuthConfigGetFunc = func() AuthConfig
)

var GetAuthConfig AuthConfigGetFunc = func() AuthConfig {
	return defaultAuthConfig{}
}

func (c defaultAuthConfig) GetRefreshTokenCookieSameSiteMode() http.SameSite {
	if isLocalEnv() {
		return http.SameSiteLaxMode
	}
	return http.SameSiteStrictMode
}
