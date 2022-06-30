package robot

import (
	"github.com/k20human/roby2000/pkg/robot/distance"
	"github.com/k20human/roby2000/pkg/robot/light"
	"github.com/k20human/roby2000/pkg/robot/movement"
	"github.com/stianeikeland/go-rpio/v4"
)

const (
	On  = "on"
	Off = "off"
)

type Robot interface {
	MoveForward() error
	MoveBackward()
	MoveLeft()
	MoveRight()
	LightsFront(action, c string) error
	LightsBack(action, c string) error
	LightsBlinking(action string) error
	Close() error
}

type robot struct {
	mover    movement.Mover
	distance distance.Measure
	light    light.Controller
}

func New() (*robot, error) {
	var err error
	var r robot

	/*var err error
	  var c *config

	  if c, err = initConfig(); err != nil {
	  	return nil, err
	  }*/

	if err = rpio.Open(); err != nil {
		return nil, err
	}

	if err = rpio.SpiBegin(rpio.Spi0); err != nil {
		return nil, err
	}

	if r.mover, err = movement.New(); err != nil {
		return nil, err
	}

	if r.distance, err = distance.New(); err != nil {
		return nil, err
	}

	if r.light, err = light.New(); err != nil {
		return nil, err
	}

	return &r, nil
}

func (r *robot) Close() error {
	// TODO do all closes and concat all errors
	if err := r.mover.Close(); err != nil {
		return err
	}

	if err := r.distance.Close(); err != nil {
		return err
	}

	if err := r.light.Close(); err != nil {
		return err
	}

	rpio.SpiEnd(rpio.Spi0)

	return rpio.Close()
}
