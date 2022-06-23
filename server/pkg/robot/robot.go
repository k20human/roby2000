package robot

import (
	"github.com/k20human/roby2000/pkg/movement"
	"github.com/stianeikeland/go-rpio/v4"
)

type Robot interface {
	Forward()
	Backward()
	Left()
	Right()
	Close() error
}

type robot struct {
	mover movement.Mover
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

	if r.mover, err = movement.New(); err != nil {
		return nil, err
	}

	return &r, nil
}

func (r *robot) Close() error {
	r.mover.Close()

	return rpio.Close()
}
