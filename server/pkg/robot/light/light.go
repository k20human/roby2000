package light

import (
	"fmt"
	"github.com/k20human/roby2000/pkg/robot/system"
	"github.com/samber/lo"
	"image/color"
	"os"
	"strconv"
)

const (
	actionStop    = "stop"
	actionDisplay = "display"
	blinkingLeft  = "left"
	blinkingRight = "right"
)

type Controller interface {
	Front(c color.RGBA) error
	Back(c color.RGBA) error
	Blinking(c color.RGBA, light string) error
	Close() error
}

type light struct {
	config        *config
	nbLeds        int
	currentColors []color.RGBA
	sys           *system.Driver
}

func New() (*light, error) {
	var err error
	var c *config
	var sys *system.Driver

	if c, err = initConfig(); err != nil {
		return nil, err
	}

	if sys, err = system.New(); err != nil {
		return nil, err
	}

	nbLeds := len(c.LedsBack) + len(c.LedsFront) + len(c.LedsBlinking)

	return &light{
		config:        c,
		nbLeds:        nbLeds,
		currentColors: make([]color.RGBA, nbLeds),
		sys:           sys,
	}, nil
}

func (l *light) Front(c color.RGBA) error {
	return l.setColor(c, l.config.LedsFront)
}

func (l *light) Back(c color.RGBA) error {
	return l.setColor(c, l.config.LedsBack)
}

func (l *light) Blinking(c color.RGBA, light string) error {
	var led []int

	if light == blinkingLeft {
		led = l.config.LedsBlinking[:1]
	} else {
		led = l.config.LedsBlinking[1:]
	}

	return l.setColor(c, led)
}

func (l *light) setColor(c color.RGBA, leds []int) error {
	for i := range l.currentColors {
		if lo.Contains(leds, i) {
			l.currentColors[i] = c
		}
	}

	return l.call(actionDisplay)
}

func (l *light) Close() error {
	for i := range l.currentColors {
		l.currentColors[i] = color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 0,
		}
	}

	return l.call(actionStop)
}

func (l *light) call(action string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	args := []string{
		currentDir + "/" + l.config.PythonScriptPath,
		action,
		strconv.Itoa(l.nbLeds),
		strconv.FormatFloat(l.config.DefaultBrightness, 'f', -1, 64),
	}

	for _, c := range l.currentColors {
		args = append(args, colorToString(c))
	}

	return l.sys.Call("python3", args)
}

func colorToString(c color.RGBA) string {
	return fmt.Sprintf("%d,%d,%d", c.R, c.G, c.B)
}
