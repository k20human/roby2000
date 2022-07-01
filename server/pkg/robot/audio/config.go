package audio

import (
	"github.com/spf13/viper"
	"os"
)

type config struct {
	Directory string `mapstructure:"directory"`
}

func defaultConfig() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	viper.SetDefault("audio.directory", path+"/")

	return nil
}

func initConfig() (*config, error) {
	var c config

	if err := defaultConfig(); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("audio", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
