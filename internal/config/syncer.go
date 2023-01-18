package config

import "fmt"

// SyncerConfig ...
type SyncerConfig struct {
	GamesSyncerPeriod     uint64 `toml:"games_syncer_period"`
	GamesSyncerEnabled    bool   `toml:"games_syncer_enabled"`
	RegistratorAPIAddress string `toml:"registrator_api_address"`
	RegistratorAPIPort    uint16 `toml:"registrator_api_port"`
}

// GetRegistratorAPIAddress ...
func GetRegistratorAPIAddress() string {
	return fmt.Sprintf("%s:%d", globalConfig.RegistratorAPIAddress, globalConfig.RegistratorAPIPort)
}
