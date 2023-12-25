package matchers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/onsi/gomega"
)

type haveSetCookieHeaderByNameMatcher struct {
	gomega.OmegaMatcher
	cookieName string
}

func HaveSetCookieHeaderByName(cookieName string) gomega.OmegaMatcher {
	return &haveSetCookieHeaderByNameMatcher{
		cookieName: cookieName,
	}
}

func (matcher haveSetCookieHeaderByNameMatcher) Match(actual interface{}) (bool, error) {
	actualResponse, ok := actual.(http.ResponseWriter)
	if !ok {
		return false, fmt.Errorf("expecting http.ResponseWriter. Got %T", actual)
	}

	setCookies := actualResponse.Header()["Set-Cookie"]
	for _, setCookie := range setCookies {
		if strings.HasPrefix(setCookie, matcher.cookieName) {
			return true, nil
		}
	}
	return false, fmt.Errorf("cookie with name '%s' not found", matcher.cookieName)
}
