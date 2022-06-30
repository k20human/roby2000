// Package ultrasonic
/* Inspired from https://github.com/shanghuiyang/rpi-devices/
HC-SR04 is an ultrasonic distance meter used to measure the distance to objects.
Spec:
  - power supply:	+5V DC
  - range:			2 - 450cm
  - resolution:		0.3cm
	 ___________________________
    |                           |
    |          HC-SR04          |
    |                           |
    |___________________________|
         |     |     |     |
        vcc  trig   echo  gnd
Connect to Raspberry Pi:
  - vcc:	any 5v pin
  - gnd:	any gnd pin
  - trig:	any data pin
  - echo:	any data pin
*/
package ultrasonic

import (
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	voiceSpeed = 34000.0 // voice speed in cm/s
)

type DistanceMeter interface {
	Dist() (float64, error)
	Close()
}

type captor struct {
	trig    rpio.Pin
	echo    rpio.Pin
	timeout int
}

func New(trig uint8, echo uint8, timeout int) *captor {
	c := &captor{
		trig:    rpio.Pin(trig),
		echo:    rpio.Pin(echo),
		timeout: timeout,
	}

	c.trig.Output()
	c.trig.Low()
	c.echo.Input()

	return c
}

func (c *captor) Dist() (float64, error) {
	c.trig.Low()
	time.Sleep(1 * time.Microsecond)
	c.trig.High()
	time.Sleep(1 * time.Microsecond)

	for n := 0; n < c.timeout && c.echo.Read() != rpio.High; n++ {
		time.Sleep(10 * time.Microsecond)
	}
	start := time.Now()

	for n := 0; n < c.timeout && c.echo.Read() != rpio.Low; n++ {
		time.Sleep(10 * time.Microsecond)
	}
	return time.Since(start).Seconds() * voiceSpeed / 2.0, nil
}

func (c *captor) Close() {
	c.trig.Low()
}
