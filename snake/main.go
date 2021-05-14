// Snake
//
// Seeeduino Xiao Expansion Board
// Buttons connected to D0 and D7

package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ssd1306"
)

type point struct {
	x, y int16
}

var display ssd1306.Device
var leftPin = machine.D0
var rightPin = machine.D7

var length = 15
var snake = make([]point, length)

func setup() {
	pinMode(machine.LED, machine.PinOutput)
	pinMode(leftPin, machine.PinInputPullup)
	pinMode(rightPin, machine.PinInputPullup)
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})
	display = ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: 0x3C,
		Width:   128,
		Height:  64,
	})
	display.ClearDisplay()
	for i := 0; i < length; i = i + 1 {
		snake[i].x = 128/2 - int16(length/2-1) + int16(i)
		snake[i].y = 64 / 2
	}
}

func loop() {

	on := color.RGBA{255, 255, 255, 255}
	off := color.RGBA{0, 0, 0, 255}

	leftButton := !leftPin.Get()
	rightButton := !rightPin.Get()

	if leftButton {
		display.SetPixel(snake[length-1].x, snake[length-1].y, off)
		for i := 0; i < length; i++ {
			snake[i].x = snake[i].x - 1
			if snake[i].x < 0 {
				snake[i].x = 127
			}
		}
	}

	if rightButton {
		display.SetPixel(snake[0].x, snake[0].y, off)
		for i := 0; i < length; i++ {
			snake[i].x = snake[i].x + 1
			if snake[i].x > 127 {
				snake[i].x = 0
			}
		}
	}

	for i := 0; i < length; i++ {
		display.SetPixel(snake[i].x, snake[i].y, on)
	}

	display.Display()

	time.Sleep(10 * time.Millisecond)
}

func main() {
	setup()
	for {
		loop()
	}
}

func pinMode(pin machine.Pin, mode machine.PinMode) {
	pin.Configure(machine.PinConfig{Mode: mode})
}
