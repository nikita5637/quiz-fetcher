package config

import "github.com/spf13/viper"

func initSynchronizerConfigureParams() {
	_ = viper.BindEnv("synchronizer.sync.timeout")
	_ = viper.BindEnv("synchronizer.syncer.games.disabled_leagues")
	_ = viper.BindEnv("synchronizer.syncer.games.enabled")
	_ = viper.BindEnv("synchronizer.syncer.games.name")
	_ = viper.BindEnv("synchronizer.syncer.games.period")
	_ = viper.BindEnv("synchronizer.syncer.results.disabled_leagues")
	_ = viper.BindEnv("synchronizer.syncer.results.enabled")
	_ = viper.BindEnv("synchronizer.syncer.results.name")
	_ = viper.BindEnv("synchronizer.syncer.results.period")
	_ = viper.BindEnv("synchronizer.syncer.results.team")
}
