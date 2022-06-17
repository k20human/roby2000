package robot

import "github.com/spf13/viper"

type config struct {
	Name string `mapstructure:"name"`
}

func defaultConfig() {
	viper.SetDefault("robot.name", "Roby2000")
}

func initConfig() (*config, error) {
	var c config

	defaultConfig()

	if err := viper.UnmarshalKey("robot", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
