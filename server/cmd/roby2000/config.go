package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

const (
	configPath = ".roby200"
	configName = "config"
	configType = "json"
	fodlerPerm = 0o755
)

type config struct {
	Port    int    `mapstructure:"port"`
	GinMode string `mapstructure:"gin_mode"`
}

func defaultConfig() {
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.gin_mode", "debug")
}

func initConfig() (*config, error) {
	var c config

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath("$HOME/" + configPath)
	viper.WatchConfig()

	defaultConfig()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			var home string

			if home, err = os.UserHomeDir(); err != nil {
				return nil, err
			}

			filename := fmt.Sprintf("%s/%s/%s.%s", home, configPath, configName, configType)

			if err = os.MkdirAll(filepath.Dir(filename), fodlerPerm); err != nil {
				return nil, errors.Errorf("unable to create directory for file '%s': %s\n", filename, err.Error())
			}

			if err = viper.WriteConfigAs(filename); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	if err := viper.UnmarshalKey("server", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
