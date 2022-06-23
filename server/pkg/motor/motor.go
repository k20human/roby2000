// Package motor
/* Inspired from https://github.com/shanghuiyang/rpi-devices
L298N is a motor driver used to control the direction and speed of DC motors.
Spec:
           _________________________________________
          |                                         |
          |                                         |
    OUT1 -|                 L298N                   |- OUT3
    OUT2 -|                                         |- OUT4
          |                                         |
          |_________________________________________|
              |   |   |     |   |   |   |   |   |
             12v GND  5V   EN1 IN1 IN2 IN3 IN4 EN2
Pins:
 - OUT1: dc motor A+
 - OUT2: dc motor A-
 - OUT3: dc motor B+
 - OUT4: dc motor B-
 - 12v: +battery
 - GND: -battery (and any gnd pin of raspberry pi if motors and raspberry pi use different battery sources)
 - IN1: any data pin
 - IN2: any data pin
 - IN3: any data pin
 - IN4: any data pin
 - EN1: must be one of GPIO 12, 13, 18 or 19 (pwn pins)
 - EN2: must be one of GPIO 12, 13, 18 or 19 (pwn pins)
*/
package motor

import (
	"github.com/stianeikeland/go-rpio/v4"
)

type Engine interface {
	Forward()
	Backward()
	Left()
	Right()
	Stop()
	Speed(s uint32)
}

type engine struct {
	inLeft1  rpio.Pin
	inLeft2  rpio.Pin
	inRight1 rpio.Pin
	inRight2 rpio.Pin
	pwmLeft  rpio.Pin
	pwmRight rpio.Pin
}

func New(inLeft1, inLeft2, pwmLeft, inRight1, inRight2, pwmRight uint8) *engine {
	e := &engine{
		inLeft1:  rpio.Pin(inLeft1),
		inLeft2:  rpio.Pin(inLeft2),
		inRight1: rpio.Pin(inRight1),
		inRight2: rpio.Pin(inRight2),
		pwmLeft:  rpio.Pin(pwmLeft),
		pwmRight: rpio.Pin(pwmRight),
	}

	e.inLeft1.Output()
	e.inLeft2.Output()
	e.inRight1.Output()
	e.inRight2.Output()
	e.inLeft1.Low()
	e.inLeft2.Low()
	e.inRight1.Low()
	e.inRight2.Low()
	e.pwmLeft.Pwm()
	e.pwmRight.Pwm()
	e.pwmLeft.Freq(1000)
	e.pwmRight.Freq(1000)

	return e
}

func (e *engine) Forward() {
	e.inLeft1.High()
	e.inLeft2.Low()
	e.inRight1.High()
	e.inRight2.Low()
	e.pwmLeft.High()
	e.pwmRight.High()
}

func (e *engine) Backward() {
	e.inLeft1.Low()
	e.inLeft2.High()
	e.inRight1.Low()
	e.inRight2.High()
	e.pwmLeft.High()
	e.pwmRight.High()
}

func (e *engine) Left() {
	e.inLeft1.Low()
	e.inLeft2.High()
	e.inRight1.High()
	e.inRight2.Low()
	e.pwmLeft.High()
	e.pwmRight.High()
}

func (e *engine) Right() {
	e.inLeft1.High()
	e.inLeft2.Low()
	e.inRight1.Low()
	e.inRight2.High()
	e.pwmLeft.High()
	e.pwmRight.High()
}

func (e *engine) Stop() {
	e.inLeft1.Low()
	e.inLeft2.Low()
	e.inRight1.Low()
	e.inRight2.Low()
	//e.pwmLeft.Low()
	//e.pwmRight.Low()
}

func (e *engine) Speed(s uint32) {
	e.pwmLeft.DutyCycle(s, 100)
	e.pwmRight.DutyCycle(s, 100)
}
