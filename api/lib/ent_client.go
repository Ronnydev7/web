package lib

import (
	"api/config"
	"api/ent/entgenerated"
	_ "api/ent/entgenerated/runtime"
	"api/intl"
	"api/intl/intlgenerated"
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type (
	NewEntClientFunc func() (*entgenerated.Client, intl.IntlError)

	CloseEntClientFunc func(client *entgenerated.Client) intl.IntlError

	CreateEntClientError struct {
		intl.IntlError
		Inner error
	}

	CloseEntClientError struct {
		intl.IntlError
		inner error
	}
)

const ENT_CLIENT_LOGGER_NAME = "lib/ent_client"

var NewEntClient NewEntClientFunc = func() (*entgenerated.Client, intl.IntlError) {
	dbConfig := config.GetDbConfig()
	db, err := sql.Open(dbConfig.GetDialect(), dbConfig.GetDbUrl())
	if err != nil {
		wrappedError := CreateEntClientError{
			Inner: err,
		}
		NewLogger(ENT_CLIENT_LOGGER_NAME).LogError(&wrappedError)
		return nil, &wrappedError
	}

	driver := entsql.OpenDB(dbConfig.GetEntSQLDialect(), db)
	result := entgenerated.NewClient(entgenerated.Driver(driver))
	return result, nil
}

var CloseEntClient CloseEntClientFunc = func(client *entgenerated.Client) intl.IntlError {
	err := client.Close()
	if err != nil {
		wrappedError := CloseEntClientError{
			inner: err,
		}
		NewLogger(ENT_CLIENT_LOGGER_NAME).LogError(&wrappedError)
		return &wrappedError
	}
	return nil
}

func (err CreateEntClientError) Error() string {
	return err.Inner.Error()
}

func (CreateEntClientError) GetIntlKey() string {
	return intlgenerated.COMMON_STRINGS__UNKNOWN_SERVER_ERROR
}

func (err CloseEntClientError) Error() string {
	return err.inner.Error()
}

func (CloseEntClientError) GetIntlKey() string {
	return intlgenerated.COMMON_STRINGS__UNKNOWN_SERVER_ERROR
}
