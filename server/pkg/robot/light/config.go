package light

import "github.com/spf13/viper"

type config struct {
	DefaultBrightness float64 `mapstructure:"brightness"`
	PythonScriptPath  string  `mapstructure:"python"`
	LedsBack          []int   `mapstructure:"leds_back"`
	LedsFront         []int   `mapstructure:"leds_front"`
	LedsBlinking      []int   `mapstructure:"leds_blinking"`
}

func defaultConfig() {
	viper.SetDefault("light.brightness", 1)
	viper.SetDefault("light.python", "leds.py")

	viper.SetDefault("light.leds_back", []int{0, 1})
	viper.SetDefault("light.leds_front", []int{2, 3})
	viper.SetDefault("light.leds_blinking", []int{4, 5})
}

func initConfig() (*config, error) {
	var c config

	defaultConfig()

	if err := viper.UnmarshalKey("light", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
