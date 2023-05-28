package config

import (
	"reflect"

	"github.com/BurntSushi/toml"
)

// GlobalConfig ...
type GlobalConfig struct {
	DatabaseConfig
	LoggerConfig
	SyncerConfig
}

var globalConfig GlobalConfig

// Value ...
type Value reflect.Value

// GetValue ...
func GetValue(key string) Value {
	val := reflect.Indirect(reflect.ValueOf(globalConfig)).FieldByName(key)
	return Value(val)
}

// Bool ...
func (c Value) Bool() bool {
	if reflect.Value(c).Kind() == reflect.Invalid {
		return false
	}

	return reflect.Value(c).Bool()
}

// String ...
func (c Value) String() string {
	if reflect.Value(c).Kind() == reflect.Invalid {
		return ""
	}

	return reflect.Value(c).String()
}

// Strings ...
func (c Value) Strings() []string {
	if reflect.Value(c).Kind() == reflect.Invalid {
		return nil
	}

	l := reflect.Value(c).Len()
	sl := make([]string, 0, l)
	for i := 0; i < l; i++ {
		sl = append(sl, reflect.Value(c).Index(i).String())
	}

	return sl
}

// Uint64 ...
func (c Value) Uint64() uint64 {
	if reflect.Value(c).Kind() == reflect.Invalid {
		return 0
	}

	return uint64(reflect.Value(c).Uint())
}

// ParseConfigFile ...
func ParseConfigFile(path string) error {
	_, err := toml.DecodeFile(path, &globalConfig)
	if err != nil {
		return err
	}

	return nil
}
