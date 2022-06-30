package distance

import (
	"github.com/k20human/roby2000/pkg/rpio/ultrasonic"
)

type Measure interface {
	Dist() (float64, error)
	Close() error
}

type captor struct {
	driver ultrasonic.DistanceMeter
}

func New() (*captor, error) {
	var err error
	var c *config

	if c, err = initConfig(); err != nil {
		return nil, err
	}

	return &captor{
		driver: ultrasonic.New(c.Trigger, c.Echo, c.Timeout),
	}, nil
}

func (c *captor) Dist() (float64, error) {
	return c.driver.Dist()
}

func (c *captor) Close() error {
	c.driver.Close()

	return nil
}
