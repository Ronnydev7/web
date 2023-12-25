package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var packageInitDone = false

const env_PRODUCT_ID = "PRODUCT_ID"

func init() {
	initViper()
}

func initViper() {
	viper.SetConfigType("env")
	productID := os.Getenv("PRODUCT_ID")
	if productID == "" {
		panic(fmt.Errorf("Required environment variable \"%s\"", env_PRODUCT_ID))
	}
	viper.SetEnvPrefix(productID)
	viper.AutomaticEnv()
}
