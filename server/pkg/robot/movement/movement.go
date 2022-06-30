package movement

import (
	"github.com/k20human/roby2000/pkg/rpio/motor"
	"time"
)

const (
	waitBeforeStop = time.Millisecond * 200
)

type Mover interface {
	Forward()
	Backward()
	Left()
	Right()
	Close() error
}

type move struct {
	wheels motor.Engine
	speed  uint32
}

func New() (*move, error) {
	var err error
	var c *config

	if c, err = initConfig(); err != nil {
		return nil, err
	}

	return &move{
		speed: c.DefaultSpeed,
		wheels: motor.New(
			c.MotorLeft.Pin1,
			c.MotorLeft.Pin2,
			c.MotorLeft.Pwm,
			c.MotorRight.Pin1,
			c.MotorRight.Pin2,
			c.MotorRight.Pwm),
	}, nil
}

func (m *move) Forward() {
	go func() {
		m.wheels.Speed(m.speed)
		m.wheels.Forward()
		time.Sleep(waitBeforeStop)
		m.wheels.Stop()
	}()
}

func (m *move) Backward() {
	go func() {
		m.wheels.Speed(m.speed)
		m.wheels.Backward()
		time.Sleep(waitBeforeStop)
		m.wheels.Stop()
	}()
}

func (m *move) Left() {
	go func() {
		m.wheels.Speed(m.speed)
		m.wheels.Left()
		time.Sleep(waitBeforeStop)
		m.wheels.Stop()
	}()
}

func (m *move) Right() {
	go func() {
		m.wheels.Speed(m.speed)
		m.wheels.Right()
		time.Sleep(waitBeforeStop)
		m.wheels.Stop()
	}()
}

func (m *move) Close() error {
	m.wheels.Stop()

	return nil
}
