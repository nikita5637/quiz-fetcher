package config

import "github.com/spf13/viper"

func initSynchronizerConfigureParams() {
	_ = viper.BindEnv("synchronizer.syncer.results.team")
}
