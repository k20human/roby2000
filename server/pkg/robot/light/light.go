package light

import (
	"bufio"
	"fmt"
	"github.com/k20human/roby2000/pkg/logger"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"image/color"
	"io"
	"os/exec"
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
	logger        *zap.Logger
}

func New() (*light, error) {
	var err error
	var c *config
	var l light

	if c, err = initConfig(); err != nil {
		return nil, err
	}

	l.config = c
	l.nbLeds = len(c.LedsBack) + len(c.LedsFront) + len(c.LedsBlinking)
	l.currentColors = make([]color.RGBA, l.nbLeds)

	if l.logger, err = logger.New(); err != nil {
		return nil, err
	}

	return &l, nil
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
	args := []string{
		l.config.PythonScriptPath,
		action,
		strconv.Itoa(l.nbLeds),
		strconv.FormatFloat(l.config.DefaultBrightness, 'f', -1, 64),
	}

	for _, c := range l.currentColors {
		args = append(args, colorToString(c))
	}

	cmd := exec.Command("python3", args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	go l.copyOutput(stdout, false)
	go l.copyOutput(stderr, true)

	err = cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}

func (l *light) copyOutput(r io.Reader, isError bool) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if isError {
			l.logger.Error(scanner.Text())
		} else {
			l.logger.Info(scanner.Text())
		}
	}
}

func colorToString(c color.RGBA) string {
	return fmt.Sprintf("%d,%d,%d", c.R, c.G, c.B)
}
