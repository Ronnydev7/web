package config

import "github.com/spf13/viper"

type (
	HmacConfig interface {
		GetRefreshTokenSecret() string
		GetAuthTokenSecret() string
		GetEmailSignupTokenSecret() string
	}

	defaultHmacConfig struct {
		HmacConfig
	}

	GetHmacConfigFunc func() HmacConfig
)

const (
	config_REFRESH_TOKEN_SECRET      = "REFRESH_TOKEN_SECRET"
	config_AUTH_TOKEN_SECRET         = "AUTO_TOKEN_SECRET"
	cofnig_EMAIL_SIGNUP_TOKEN_SECRET = "EMAIL_SIGNUP_TOKEN_SECRET"
)

var GetHmacConfig GetHmacConfigFunc = func() HmacConfig {
	return &defaultHmacConfig{}
}

func (defaultHmacConfig) GetRefreshTokenSecret() string {
	return viper.GetString(config_REFRESH_TOKEN_SECRET)
}

func (defaultHmacConfig) GetAuthTokenSecret() string {
	return viper.GetString(config_AUTH_TOKEN_SECRET)
}

func (defaultHmacConfig) GetEmailSignupTokenSecret() string {
	return viper.GetString(cofnig_EMAIL_SIGNUP_TOKEN_SECRET)
}
