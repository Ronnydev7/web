package cookies

import (
	"net/http"
	"time"
)

const REFRESH_TOKEN_COOKIE_NAME = "rt"

func NewRefreshTokenCookie(token string, expiredAt time.Time, sameSite http.SameSite) *http.Cookie {
	return &http.Cookie{
		Name:     REFRESH_TOKEN_COOKIE_NAME,
		Value:    token,
		HttpOnly: true,
		SameSite: sameSite,
		Path:     "/",
		Expires:  expiredAt,
	}
}
