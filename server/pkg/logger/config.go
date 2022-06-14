package logger

import "github.com/spf13/viper"

type config struct {
	Level string `mapstructure:"level"`
}

func defaultConfig() {
	viper.SetDefault("logger.level", "debug")
}

func initConfig() (*config, error) {
	var c config

	defaultConfig()

	if err := viper.UnmarshalKey("logger", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
