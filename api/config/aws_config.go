package config

import "github.com/spf13/viper"

type (
	AwsConfig interface {
		GetRegion() string
		GetAccessKeyId() string
		GetSecretAccessKey() string
	}

	defaultAwsConfig struct {
		AwsConfig
	}

	GetAwsConfigFunc func() AwsConfig
)

const (
	config_AWS_REGION            = "AWS_REGION"
	config_AWS_ACCESS_KEY_ID     = "AWS_ACCESS_KEY_ID"
	config_AWS_SECRET_ACCESS_KEY = "AWS_SECRET_ACCESS_KEY"
)

var GetAwsConfig GetAwsConfigFunc = func() AwsConfig {
	return &defaultAwsConfig{}
}

func (defaultAwsConfig) GetRegion() string {
	return viper.GetString(config_AWS_REGION)
}

func (defaultAwsConfig) GetAccessKeyId() string {
	return viper.GetString(config_AWS_ACCESS_KEY_ID)
}

func (defaultAwsConfig) GetSecretAccessKey() string {
	return viper.GetString(config_AWS_SECRET_ACCESS_KEY)
}
