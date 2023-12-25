package config

import "github.com/spf13/viper"

type (
	GoogleRecaptchaConfig interface {
		GetProjectID() string
		GetSiteKey() string
	}

	envGoogleRecaptchaConfig struct {
		GoogleRecaptchaConfig
	}

	GetGoogleRecaptchaConfigFunc func() GoogleRecaptchaConfig
)

const (
	config_GOOGLE_APP_PROJECT_ID = "GOOGLE_APP_PROJECT_ID"

	config_GOOGLE_RECAPTCHA_SITE_KEY = "GOOGLE_RECAPTCHA_SITE_KEY"
)

var (
	defaultGoogleRecaptchaConfig = envGoogleRecaptchaConfig{}

	GetGoogleRecaptchaConfig GetGoogleRecaptchaConfigFunc = func() GoogleRecaptchaConfig {
		return &defaultGoogleRecaptchaConfig
	}
)

func (envGoogleRecaptchaConfig) GetProjectID() string {
	return viper.GetString(config_GOOGLE_APP_PROJECT_ID)
}

func (envGoogleRecaptchaConfig) GetSiteKey() string {
	return viper.GetString(config_GOOGLE_RECAPTCHA_SITE_KEY)
}
