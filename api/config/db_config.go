package config

import "github.com/spf13/viper"

type (
	DbConfig interface {
		GetDbUrl() string

		// This is the dialect used by the DB Connection
		GetDialect() string

		// This is the dialect used by entgo when querying the DB
		GetEntSQLDialect() string
	}

	defaultDbConfig struct {
		DbConfig
	}

	GetDbConfigFunc func() DbConfig
)

const (
	config_DB_URL = "DB_URL"

	config_DB_DIALECT = "DB_DIALECT"

	config_ENT_SQL_DIALECT = "ENT_SQL_DIALECT"
)

var (
	dbConfig                    = defaultDbConfig{}
	GetDbConfig GetDbConfigFunc = func() DbConfig {
		return &dbConfig
	}
)

func (defaultDbConfig) GetDbUrl() string {
	return viper.GetString(config_DB_URL)
}

func (defaultDbConfig) GetDialect() string {
	return viper.GetString(config_DB_DIALECT)
}

func (defaultDbConfig) GetEntSQLDialect() string {
	return viper.GetString(config_ENT_SQL_DIALECT)
}
