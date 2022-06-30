package movement

import "github.com/spf13/viper"

type configMotor struct {
	Pin1 uint8 `mapstructure:"pin1"`
	Pin2 uint8 `mapstructure:"pin2"`
	Pwm  uint8 `mapstructure:"pwm"`
}

type config struct {
	MotorLeft    configMotor `mapstructure:"motor_left"`
	MotorRight   configMotor `mapstructure:"motor_right"`
	DefaultSpeed uint32      `mapstructure:"speed"`
}

func defaultConfig() {
	viper.SetDefault("movement.motor_left.pin1", 14)
	viper.SetDefault("movement.motor_left.pin2", 26)
	viper.SetDefault("movement.motor_left.pwm", 19)

	viper.SetDefault("movement.motor_right.pin1", 22)
	viper.SetDefault("movement.motor_right.pin2", 27)
	viper.SetDefault("movement.motor_right.pwm", 13)

	viper.SetDefault("movement.speed", 50)
}

func initConfig() (*config, error) {
	var c config

	defaultConfig()

	if err := viper.UnmarshalKey("movement", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
