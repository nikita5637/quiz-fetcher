package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func initFetcherConfigureParams() {
	_ = viper.BindEnv("fetcher.registrator_api.address")
}

// GetRegistratorAPIAddress ...
func GetRegistratorAPIAddress() string {
	return fmt.Sprintf("%s:%d",
		viper.GetString("fetcher.registrator_api.address"),
		viper.GetUint32("fetcher.registrator_api.port"),
	)
}
