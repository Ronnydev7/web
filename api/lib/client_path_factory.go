package lib

import (
	"api/config"
	"fmt"
	"net/url"
)

type (
	ClientPathFactory interface {
		CreateConfirmEmailSignupPath(emailSignupToken string) *url.URL
		CreateResetPasswordUrl(resetPasswordToken string) *url.URL
	}

	ClientPathFactoryWithConfig struct {
		ClientPathFactory
		urlConfig config.UrlConfig
	}
)

var NewClientPathFactory = func(urlConfig config.UrlConfig) ClientPathFactory {
	return ClientPathFactoryWithConfig{
		urlConfig: urlConfig,
	}
}

func (factory ClientPathFactoryWithConfig) createBaseUrl(path string) *url.URL {
	baseUrlString := fmt.Sprintf(
		"%s://%s%s",
		factory.urlConfig.GetProtocol(),
		factory.urlConfig.GetHostname(),
		path,
	)
	result, err := url.Parse(baseUrlString)
	if err != nil {
		panic(err)
	}
	return result
}

func (factory ClientPathFactoryWithConfig) CreateConfirmEmailSignupPath(emailSignupToken string) *url.URL {
	result := factory.createBaseUrl("/confirm-email-signup")
	query := result.Query()
	query.Add("email_signup_token", emailSignupToken)
	result.RawQuery = query.Encode()
	return result
}

func (factory ClientPathFactoryWithConfig) CreateResetPasswordUrl(resetPasswordToken string) *url.URL {
	result := factory.createBaseUrl("/reset-password")
	query := result.Query()
	query.Add("reset_password_token", resetPasswordToken)
	result.RawQuery = query.Encode()
	return result
}
