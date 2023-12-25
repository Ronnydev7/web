package config

import "github.com/spf13/viper"

type (
	MailerConfig interface {
		GetNoResponseEmailName() string
		GetNoResponseEmailAddress() string
	}

	viperMailerConfig struct {
		MailerConfig
	}

	GetMailerConfigFunc func() MailerConfig
)

const (
	config_MAILER_NO_RESPONSE_NAME          = "MAILER_NO_RESPONSE_NAME"
	config_MAILER_NO_RESPONSE_EMAIL_ADDRESS = "MAILER_NO_RESPONSE_EMAIL_ADDRESS"
)

var GetMailerConfig GetMailerConfigFunc = func() MailerConfig {
	return &viperMailerConfig{}
}

func (viperMailerConfig) GetNoResponseEmailName() string {
	return viper.GetString(config_MAILER_NO_RESPONSE_NAME)
}

func (viperMailerConfig) GetNoResponseEmailAddress() string {
	return viper.GetString(config_MAILER_NO_RESPONSE_EMAIL_ADDRESS)
}
