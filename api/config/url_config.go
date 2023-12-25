package config

import "github.com/spf13/viper"

type (
	UrlConfig interface {
		GetProtocol() string
		GetHostname() string
	}

	defaultUrlConfig struct {
		UrlConfig
	}

	GetUrlConfigFunc func() UrlConfig
)

const (
	config_URL_PROTOCOL = "ORIGIN_PROTOCOL"
	config_URL_HOSTNAME = "ORIGIN_HOSTNAME"
)

var GetUrlConfig GetUrlConfigFunc = func() UrlConfig {
	return &defaultUrlConfig{}
}

func (defaultUrlConfig) GetProtocol() string {
	return viper.GetString(config_URL_PROTOCOL)
}

func (defaultUrlConfig) GetHostname() string {
	return viper.GetString(config_URL_HOSTNAME)
}
