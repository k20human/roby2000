package audio

import (
	"github.com/k20human/roby2000/pkg/robot/system"
)

type Player interface {
	Play(filename string) error
	Close() error
}

type audio struct {
	directory string
	sys       *system.Driver
}

func New() (*audio, error) {
	var err error
	var c *config
	var sys *system.Driver

	if c, err = initConfig(); err != nil {
		return nil, err
	}

	if sys, err = system.New(); err != nil {
		return nil, err
	}

	return &audio{
		directory: c.Directory,
		sys:       sys,
	}, nil
}

func (a *audio) Play(filename string) error {
	return a.sys.Call("mplayer", []string{a.directory + filename})
}

func (a *audio) Close() error {
	return nil
}
