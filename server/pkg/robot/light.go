package robot

import (
	"github.com/pkg/errors"
	"image/color"
	"strconv"
	"strings"
	"time"
)

var ErrWrongColor = errors.New("Wrong color parameter")

func (r *robot) LightsFront(action, c string) error {
	if action == Off {
		return r.light.Front(blackColor())
	} else {
		cl, err := stringToColor(c)
		if err != nil {
			return err
		}

		return r.light.Front(cl)
	}
}

func (r *robot) LightsBack(action, c string) error {
	if action == Off {
		return r.light.Back(blackColor())
	} else {
		cl, err := stringToColor(c)
		if err != nil {
			return err
		}

		return r.light.Back(cl)
	}
}

func (r *robot) LightsBlinking(direction string) error {
	cl, err := stringToColor("230,138,0")
	if err != nil {
		return err
	}

	go func() {
		for i := 1; i <= 3; i++ {
			r.light.Blinking(cl, direction)
			time.Sleep(200 * time.Millisecond)
			r.light.Blinking(blackColor(), direction)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	return nil
}

func blackColor() color.RGBA {
	return color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 0,
	}
}

func stringToColor(c string) (color.RGBA, error) {
	var r, g, b int
	var err error

	colors := strings.Split(c, ",")

	if len(colors) != 3 {
		return color.RGBA{}, ErrWrongColor
	}

	if r, err = strconv.Atoi(colors[0]); err != nil {
		return color.RGBA{}, err
	}

	if g, err = strconv.Atoi(colors[1]); err != nil {
		return color.RGBA{}, err
	}

	if b, err = strconv.Atoi(colors[2]); err != nil {
		return color.RGBA{}, err
	}

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 0,
	}, nil
}
