package distance

import "github.com/spf13/viper"

type config struct {
	Trigger uint8 `mapstructure:"trigger"`
	Echo    uint8 `mapstructure:"echo"`
	Timeout int   `mapstructure:"timeout"`
}

func defaultConfig() {
	viper.SetDefault("distance.trigger", 24)
	viper.SetDefault("distance.echo", 23)
	viper.SetDefault("distance.timeout", 3600)
}

func initConfig() (*config, error) {
	var c config

	defaultConfig()

	if err := viper.UnmarshalKey("distance", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
