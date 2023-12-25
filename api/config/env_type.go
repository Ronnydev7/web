package config

import "github.com/spf13/viper"

type envType = int

const (
	LOCAL envType = iota
	PROD
)

const (
	config_ENV_TYPE = "ENV_TYPE"
)

func getEnvType() envType {
	envTypeString := viper.GetString(config_ENV_TYPE)

	switch envTypeString {
	case "PROD":
		return PROD
	}
	return LOCAL
}

func isLocalEnv() bool {
	return getEnvType() == LOCAL
}
