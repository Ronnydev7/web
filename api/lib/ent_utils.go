package lib

import (
	"api/ent/entgenerated"
	"api/ent/entgenerated/privacy"
	"api/intl/intlgenerated"
	"errors"
)

type (
	EntUtils interface {
		RollbackTx(tx *entgenerated.Tx, permissionError error) error
		HandlePermissionError(error) error
	}

	defaultEntUtils struct {
		EntUtils
	}

	NewEntUtilsFunc func() EntUtils
)

var NewEntUtils NewEntUtilsFunc = func() EntUtils {
	return &defaultEntUtils{}
}

const ENT_UTILS_LOGGER_NAME = "lib.EntUtils"

func newLogger() Logger {
	return NewLogger(ENT_UTILS_LOGGER_NAME)
}

func (utils defaultEntUtils) RollbackTx(tx *entgenerated.Tx, permissionError error) error {
	finalError := utils.HandlePermissionError(permissionError)
	if rerr := tx.Rollback(); rerr != nil {
		newLogger().LogError(rerr)
		finalError = errors.New(intlgenerated.COMMON_STRINGS__UNKNOWN_SERVER_ERROR)
	}
	return finalError
}

func (defaultEntUtils) HandlePermissionError(err error) error {
	if errors.Is(err, privacy.Deny) {
		return errors.New(intlgenerated.COMMON_STRINGS__UNAUTHORIZED)
	}
	return err
}
