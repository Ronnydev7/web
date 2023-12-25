package config

import "github.com/spf13/viper"

type (
	SendgridConfig interface {
		GetApiKey() string
		GetEmailSignupConfirmationTemplateId() string
		GetResetPasswordTemplateId() string
	}

	viperSendgridConfig struct {
		SendgridConfig
	}

	GetSendgridConfigFunc func() SendgridConfig
)

const (
	config_SENDGRID_API_KEY                               = "SENDGRID_API_KEY"
	config_SENDGRID_EMAIL_SIGNUP_CONFIRMATION_TEMPLATE_ID = "SENDGRID_EMAIL_SIGNUP_CONFIRMATION_TEMPLATE_ID"
	config_SENDGRID_RESET_PASSWORD_TEMPLATE_ID            = "SENDGRID_RESET_PASSWORD_TEMPLATE_ID"
)

var GetSendgridConfig GetSendgridConfigFunc = func() SendgridConfig {
	return &viperSendgridConfig{}
}

func (viperSendgridConfig) GetApiKey() string {
	return viper.GetString(config_SENDGRID_API_KEY)
}

func (viperSendgridConfig) GetEmailSignupConfirmationTemplateId() string {
	return viper.GetString(config_SENDGRID_EMAIL_SIGNUP_CONFIRMATION_TEMPLATE_ID)
}

func (viperSendgridConfig) GetResetPasswordTemplateId() string {
	return viper.GetString(config_SENDGRID_RESET_PASSWORD_TEMPLATE_ID)
}
