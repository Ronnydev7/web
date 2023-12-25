package lib

import (
	"encoding/json"
	"io"
	"os"
)

type (
	JsonUtils interface {
		UnmarshalFile(path string, result interface{}) error
	}

	defaultJsonUtils struct {
		JsonUtils
	}

	NewJsonUtilsFunc func() JsonUtils
)

var NewJsonUtils NewJsonUtilsFunc = func() JsonUtils {
	return &defaultJsonUtils{}
}

func (defaultJsonUtils) UnmarshalFile(path string, result interface{}) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	json.Unmarshal(byteValue, result)
	return nil
}
